package lexers_test

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/g"
)

const lexerBenchSource = `// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP client implementation. See RFC 7230 through 7235.
//
// This is the low-level Transport implementation of RoundTripper.
// The high-level interface is in client.go.

package http

import (
	"bufio"
	"compress/gzip"
	"container/list"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http/httptrace"
	"net/textproto"
	"net/url"
	"os"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang_org/x/net/http/httpguts"
	"golang_org/x/net/http/httpproxy"
)

// DefaultTransport is the default implementation of Transport and is
// used by DefaultClient. It establishes network connections as needed
// and caches them for reuse by subsequent calls. It uses HTTP proxies
// as directed by the $HTTP_PROXY and $NO_PROXY (or $http_proxy and
// $no_proxy) environment variables.
var DefaultTransport RoundTripper = &Transport{
	Proxy: ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

// DefaultMaxIdleConnsPerHost is the default value of Transport's
// MaxIdleConnsPerHost.
const DefaultMaxIdleConnsPerHost = 2

// connsPerHostClosedCh is a closed channel used by MaxConnsPerHost
// for the property that receives from a closed channel return the
// zero value.
var connsPerHostClosedCh = make(chan struct{})

func init() {
	close(connsPerHostClosedCh)
}

// Transport is an implementation of RoundTripper that supports HTTP,
// HTTPS, and HTTP proxies (for either HTTP or HTTPS with CONNECT).
//
// By default, Transport caches connections for future re-use.
// This may leave many open connections when accessing many hosts.
// This behavior can be managed using Transport's CloseIdleConnections method
// and the MaxIdleConnsPerHost and DisableKeepAlives fields.
//
// Transports should be reused instead of created as needed.
// Transports are safe for concurrent use by multiple goroutines.
//
// A Transport is a low-level primitive for making HTTP and HTTPS requests.
// For high-level functionality, such as cookies and redirects, see Client.
//
// Transport uses HTTP/1.1 for HTTP URLs and either HTTP/1.1 or HTTP/2
// for HTTPS URLs, depending on whether the server supports HTTP/2,
// and how the Transport is configured. The DefaultTransport supports HTTP/2.
// To explicitly enable HTTP/2 on a transport, use golang.org/x/net/http2
// and call ConfigureTransport. See the package docs for more about HTTP/2.
//
// The Transport will send CONNECT requests to a proxy for its own use
// when processing HTTPS requests, but Transport should generally not
// be used to send a CONNECT request. That is, the Request passed to
// the RoundTrip method should not have a Method of "CONNECT", as Go's
// HTTP/1.x implementation does not support full-duplex request bodies
// being written while the response body is streamed. Go's HTTP/2
// implementation does support full duplex, but many CONNECT proxies speak
// HTTP/1.x.
//
// Responses with status codes in the 1xx range are either handled
// automatically (100 expect-continue) or ignored. The one
// exception is HTTP status code 101 (Switching Protocols), which is
// considered a terminal status and returned by RoundTrip. To see the
// ignored 1xx responses, use the httptrace trace package's
// ClientTrace.Got1xxResponse.
type Transport struct {
	idleMu     sync.Mutex
	wantIdle   bool                                // user has requested to close all idle conns
	idleConn   map[connectMethodKey][]*persistConn // most recently used at end
	idleConnCh map[connectMethodKey]chan *persistConn
	idleLRU    connLRU

	reqMu       sync.Mutex
	reqCanceler map[*Request]func(error)

	altMu    sync.Mutex   // guards changing altProto only
	altProto atomic.Value // of nil or map[string]RoundTripper, key is URI scheme

	connCountMu          sync.Mutex
	connPerHostCount     map[connectMethodKey]int
	connPerHostAvailable map[connectMethodKey]chan struct{}

	// Proxy specifies a function to return a proxy for a given
	// Request. If the function returns a non-nil error, the
	// request is aborted with the provided error.
	//
	// The proxy type is determined by the URL scheme. "http",
	// "https", and "socks5" are supported. If the scheme is empty,
	// "http" is assumed.
	//
	// If Proxy is nil or returns a nil *URL, no proxy is used.
	Proxy func(*Request) (*url.URL, error)

	// DialContext specifies the dial function for creating unencrypted TCP connections.
	// If DialContext is nil (and the deprecated Dial below is also nil),
	// then the transport dials using package net.
	//
	// DialContext runs concurrently with calls to RoundTrip.
	// A RoundTrip call that initiates a dial may end up using
	// an connection dialed previously when the earlier connection
	// becomes idle before the later DialContext completes.
	DialContext func(ctx context.Context, network, addr string) (net.Conn, error)

	// Dial specifies the dial function for creating unencrypted TCP connections.
	//
	// Dial runs concurrently with calls to RoundTrip.
	// A RoundTrip call that initiates a dial may end up using
	// an connection dialed previously when the earlier connection
	// becomes idle before the later Dial completes.
	//
	// Deprecated: Use DialContext instead, which allows the transport
	// to cancel dials as soon as they are no longer needed.
	// If both are set, DialContext takes priority.
	Dial func(network, addr string) (net.Conn, error)

	// DialTLS specifies an optional dial function for creating
	// TLS connections for non-proxied HTTPS requests.
	//
	// If DialTLS is nil, Dial and TLSClientConfig are used.
	//
	// If DialTLS is set, the Dial hook is not used for HTTPS
	// requests and the TLSClientConfig and TLSHandshakeTimeout
	// are ignored. The returned net.Conn is assumed to already be
	// past the TLS handshake.
	DialTLS func(network, addr string) (net.Conn, error)

	// TLSClientConfig specifies the TLS configuration to use with
	// tls.Client.
	// If nil, the default configuration is used.
	// If non-nil, HTTP/2 support may not be enabled by default.
	TLSClientConfig *tls.Config

	// TLSHandshakeTimeout specifies the maximum amount of time waiting to
	// wait for a TLS handshake. Zero means no timeout.
	TLSHandshakeTimeout time.Duration

	// DisableKeepAlives, if true, disables HTTP keep-alives and
	// will only use the connection to the server for a single
	// HTTP request.
	//
	// This is unrelated to the similarly named TCP keep-alives.
	DisableKeepAlives bool

	// DisableCompression, if true, prevents the Transport from
	// requesting compression with an "Accept-Encoding: gzip"
	// request header when the Request contains no existing
	// Accept-Encoding value. If the Transport requests gzip on
	// its own and gets a gzipped response, it's transparently
	// decoded in the Response.Body. However, if the user
	// explicitly requested gzip it is not automatically
	// uncompressed.
	DisableCompression bool

	// MaxIdleConns controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	MaxIdleConns int

	// MaxIdleConnsPerHost, if non-zero, controls the maximum idle
	// (keep-alive) connections to keep per-host. If zero,
	// DefaultMaxIdleConnsPerHost is used.
	MaxIdleConnsPerHost int

	// MaxConnsPerHost optionally limits the total number of
	// connections per host, including connections in the dialing,
	// active, and idle states. On limit violation, dials will block.
	//
	// Zero means no limit.
	//
	// For HTTP/2, this currently only controls the number of new
	// connections being created at a time, instead of the total
	// number. In practice, hosts using HTTP/2 only have about one
	// idle connection, though.
	MaxConnsPerHost int

	// IdleConnTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	// Zero means no limit.
	IdleConnTimeout time.Duration

	// ResponseHeaderTimeout, if non-zero, specifies the amount of
	// time to wait for a server's response headers after fully
	// writing the request (including its body, if any). This
	// time does not include the time to read the response body.
	ResponseHeaderTimeout time.Duration

	// ExpectContinueTimeout, if non-zero, specifies the amount of
	// time to wait for a server's first response headers after fully
	// writing the request headers if the request has an
	// "Expect: 100-continue" header. Zero means no timeout and
	// causes the body to be sent immediately, without
	// waiting for the server to approve.
	// This time does not include the time to send the request header.
	ExpectContinueTimeout time.Duration

	// TLSNextProto specifies how the Transport switches to an
	// alternate protocol (such as HTTP/2) after a TLS NPN/ALPN
	// protocol negotiation. If Transport dials an TLS connection
	// with a non-empty protocol name and TLSNextProto contains a
	// map entry for that key (such as "h2"), then the func is
	// called with the request's authority (such as "example.com"
	// or "example.com:1234") and the TLS connection. The function
	// must return a RoundTripper that then handles the request.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	TLSNextProto map[string]func(authority string, c *tls.Conn) RoundTripper

	// ProxyConnectHeader optionally specifies headers to send to
	// proxies during CONNECT requests.
	ProxyConnectHeader Header

	// MaxResponseHeaderBytes specifies a limit on how many
	// response bytes are allowed in the server's response
	// header.
	//
	// Zero means to use a default limit.
	MaxResponseHeaderBytes int64

	// nextProtoOnce guards initialization of TLSNextProto and
	// h2transport (via onceSetNextProtoDefaults)
	nextProtoOnce sync.Once
	h2transport   h2Transport // non-nil if http2 wired up
}

// h2Transport is the interface we expect to be able to call from
// net/http against an *http2.Transport that's either bundled into
// h2_bundle.go or supplied by the user via x/net/http2.
//
// We name it with the "h2" prefix to stay out of the "http2" prefix
// namespace used by x/tools/cmd/bundle for h2_bundle.go.
type h2Transport interface {
	CloseIdleConnections()
}

// onceSetNextProtoDefaults initializes TLSNextProto.
// It must be called via t.nextProtoOnce.Do.
func (t *Transport) onceSetNextProtoDefaults() {
	if strings.Contains(os.Getenv("GODEBUG"), "http2client=0") {
		return
	}

	// If they've already configured http2 with
	// golang.org/x/net/http2 instead of the bundled copy, try to
	// get at its http2.Transport value (via the the "https"
	// altproto map) so we can call CloseIdleConnections on it if
	// requested. (Issue 22891)
	altProto, _ := t.altProto.Load().(map[string]RoundTripper)
	if rv := reflect.ValueOf(altProto["https"]); rv.IsValid() && rv.Type().Kind() == reflect.Struct && rv.Type().NumField() == 1 {
		if v := rv.Field(0); v.CanInterface() {
			if h2i, ok := v.Interface().(h2Transport); ok {
				t.h2transport = h2i
			}
		}
	}

	if t.TLSNextProto != nil {
		// This is the documented way to disable http2 on a
		// Transport.
		return
	}
	if t.TLSClientConfig != nil || t.Dial != nil || t.DialTLS != nil {
		// Be conservative and don't automatically enable
		// http2 if they've specified a custom TLS config or
		// custom dialers. Let them opt-in themselves via
		// http2.ConfigureTransport so we don't surprise them
		// by modifying their tls.Config. Issue 14275.
		return
	}
	t2, err := http2configureTransport(t)
	if err != nil {
		log.Printf("Error enabling Transport HTTP/2 support: %v", err)
		return
	}
	t.h2transport = t2

	// Auto-configure the http2.Transport's MaxHeaderListSize from
	// the http.Transport's MaxResponseHeaderBytes. They don't
	// exactly mean the same thing, but they're close.
	//
	// TODO: also add this to x/net/http2.Configure Transport, behind
	// a +build go1.7 build tag:
	if limit1 := t.MaxResponseHeaderBytes; limit1 != 0 && t2.MaxHeaderListSize == 0 {
		const h2max = 1<<32 - 1
		if limit1 >= h2max {
			t2.MaxHeaderListSize = h2max
		} else {
			t2.MaxHeaderListSize = uint32(limit1)
		}
	}
}

// ProxyFromEnvironment returns the URL of the proxy to use for a
// given request, as indicated by the environment variables
// HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions
// thereof). HTTPS_PROXY takes precedence over HTTP_PROXY for https
// requests.
//
// The environment values may be either a complete URL or a
// "host[:port]", in which case the "http" scheme is assumed.
// An error is returned if the value is a different form.
//
// A nil URL and nil error are returned if no proxy is defined in the
// environment, or a proxy should not be used for the given request,
// as defined by NO_PROXY.
//
// As a special case, if req.URL.Host is "localhost" (with or without
// a port number), then a nil URL and nil error will be returned.
func ProxyFromEnvironment(req *Request) (*url.URL, error) {
	return envProxyFunc()(req.URL)
}

// ProxyURL returns a proxy function (for use in a Transport)
// that always returns the same URL.
func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error) {
	return func(*Request) (*url.URL, error) {
		return fixedURL, nil
	}
}

// transportRequest is a wrapper around a *Request that adds
// optional extra headers to write and stores any error to return
// from roundTrip.
type transportRequest struct {
	*Request                        // original request, not to be mutated
	extra    Header                 // extra headers to write, or nil
	trace    *httptrace.ClientTrace // optional

	mu  sync.Mutex // guards err
	err error      // first setError value for mapRoundTripError to consider
}

func (tr *transportRequest) extraHeaders() Header {
	if tr.extra == nil {
		tr.extra = make(Header)
	}
	return tr.extra
}

func (tr *transportRequest) setError(err error) {
	tr.mu.Lock()
	if tr.err == nil {
		tr.err = err
	}
	tr.mu.Unlock()
}

// roundTrip implements a RoundTripper over HTTP.
func (t *Transport) roundTrip(req *Request) (*Response, error) {
	t.nextProtoOnce.Do(t.onceSetNextProtoDefaults)
	ctx := req.Context()
	trace := httptrace.ContextClientTrace(ctx)

	if req.URL == nil {
		req.closeBody()
		return nil, errors.New("http: nil Request.URL")
	}
	if req.Header == nil {
		req.closeBody()
		return nil, errors.New("http: nil Request.Header")
	}
	scheme := req.URL.Scheme
	isHTTP := scheme == "http" || scheme == "https"
	if isHTTP {
		for k, vv := range req.Header {
			if !httpguts.ValidHeaderFieldName(k) {
				return nil, fmt.Errorf("net/http: invalid header field name %q", k)
			}
			for _, v := range vv {
				if !httpguts.ValidHeaderFieldValue(v) {
					return nil, fmt.Errorf("net/http: invalid header field value %q for key %v", v, k)
				}
			}
		}
	}

	altProto, _ := t.altProto.Load().(map[string]RoundTripper)
	if altRT := altProto[scheme]; altRT != nil {
		if resp, err := altRT.RoundTrip(req); err != ErrSkipAltProtocol {
			return resp, err
		}
	}
	if !isHTTP {
		req.closeBody()
		return nil, &badStringError{"unsupported protocol scheme", scheme}
	}
	if req.Method != "" && !validMethod(req.Method) {
		return nil, fmt.Errorf("net/http: invalid method %q", req.Method)
	}
	if req.URL.Host == "" {
		req.closeBody()
		return nil, errors.New("http: no Host in request URL")
	}

	for {
		select {
		case <-ctx.Done():
			req.closeBody()
			return nil, ctx.Err()
		default:
		}

		// treq gets modified by roundTrip, so we need to recreate for each retry.
		treq := &transportRequest{Request: req, trace: trace}
		cm, err := t.connectMethodForRequest(treq)
		if err != nil {
			req.closeBody()
			return nil, err
		}

		// Get the cached or newly-created connection to either the
		// host (for http or https), the http proxy, or the http proxy
		// pre-CONNECTed to https server. In any case, we'll be ready
		// to send it requests.
		pconn, err := t.getConn(treq, cm)
		if err != nil {
			t.setReqCanceler(req, nil)
			req.closeBody()
			return nil, err
		}

		var resp *Response
		if pconn.alt != nil {
			// HTTP/2 path.
			t.decHostConnCount(cm.key()) // don't count cached http2 conns toward conns per host
			t.setReqCanceler(req, nil)   // not cancelable with CancelRequest
			resp, err = pconn.alt.RoundTrip(req)
		} else {
			resp, err = pconn.roundTrip(treq)
		}
		if err == nil {
			return resp, nil
		}
		if !pconn.shouldRetryRequest(req, err) {
			// Issue 16465: return underlying net.Conn.Read error from peek,
			// as we've historically done.
			if e, ok := err.(transportReadFromServerError); ok {
				err = e.err
			}
			return nil, err
		}
		testHookRoundTripRetried()

		// Rewind the body if we're able to.  (HTTP/2 does this itself so we only
		// need to do it for HTTP/1.1 connections.)
		if req.GetBody != nil && pconn.alt == nil {
			newReq := *req
			var err error
			newReq.Body, err = req.GetBody()
			if err != nil {
				return nil, err
			}
			req = &newReq
		}
	}
}

// shouldRetryRequest reports whether we should retry sending a failed
// HTTP request on a new connection. The non-nil input error is the
// error from roundTrip.
func (pc *persistConn) shouldRetryRequest(req *Request, err error) bool {
	if http2isNoCachedConnError(err) {
		// Issue 16582: if the user started a bunch of
		// requests at once, they can all pick the same conn
		// and violate the server's max concurrent streams.
		// Instead, match the HTTP/1 behavior for now and dial
		// again to get a new TCP connection, rather than failing
		// this request.
		return true
	}
	if err == errMissingHost {
		// User error.
		return false
	}
	if !pc.isReused() {
		// This was a fresh connection. There's no reason the server
		// should've hung up on us.
		//
		// Also, if we retried now, we could loop forever
		// creating new connections and retrying if the server
		// is just hanging up on us because it doesn't like
		// our request (as opposed to sending an error).
		return false
	}
	if _, ok := err.(nothingWrittenError); ok {
		// We never wrote anything, so it's safe to retry, if there's no body or we
		// can "rewind" the body with GetBody.
		return req.outgoingLength() == 0 || req.GetBody != nil
	}
	if !req.isReplayable() {
		// Don't retry non-idempotent requests.
		return false
	}
	if _, ok := err.(transportReadFromServerError); ok {
		// We got some non-EOF net.Conn.Read failure reading
		// the 1st response byte from the server.
		return true
	}
	if err == errServerClosedIdle {
		// The server replied with io.EOF while we were trying to
		// read the response. Probably an unfortunately keep-alive
		// timeout, just as the client was writing a request.
		return true
	}
	return false // conservatively
}

// ErrSkipAltProtocol is a sentinel error value defined by Transport.RegisterProtocol.
var ErrSkipAltProtocol = errors.New("net/http: skip alternate protocol")

// RegisterProtocol registers a new protocol with scheme.
// The Transport will pass requests using the given scheme to rt.
// It is rt's responsibility to simulate HTTP request semantics.
//
// RegisterProtocol can be used by other packages to provide
// implementations of protocol schemes like "ftp" or "file".
//
// If rt.RoundTrip returns ErrSkipAltProtocol, the Transport will
// handle the RoundTrip itself for that one request, as if the
// protocol were not registered.
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper) {
	t.altMu.Lock()
	defer t.altMu.Unlock()
	oldMap, _ := t.altProto.Load().(map[string]RoundTripper)
	if _, exists := oldMap[scheme]; exists {
		panic("protocol " + scheme + " already registered")
	}
	newMap := make(map[string]RoundTripper)
	for k, v := range oldMap {
		newMap[k] = v
	}
	newMap[scheme] = rt
	t.altProto.Store(newMap)
}

// CloseIdleConnections closes any connections which were previously
// connected from previous requests but are now sitting idle in
// a "keep-alive" state. It does not interrupt any connections currently
// in use.
func (t *Transport) CloseIdleConnections() {
	t.nextProtoOnce.Do(t.onceSetNextProtoDefaults)
	t.idleMu.Lock()
	m := t.idleConn
	t.idleConn = nil
	t.idleConnCh = nil
	t.wantIdle = true
	t.idleLRU = connLRU{}
	t.idleMu.Unlock()
	for _, conns := range m {
		for _, pconn := range conns {
			pconn.close(errCloseIdleConns)
		}
	}
	if t2 := t.h2transport; t2 != nil {
		t2.CloseIdleConnections()
	}
}

// CancelRequest cancels an in-flight request by closing its connection.
// CancelRequest should only be called after RoundTrip has returned.
//
// Deprecated: Use Request.WithContext to create a request with a
// cancelable context instead. CancelRequest cannot cancel HTTP/2
// requests.
func (t *Transport) CancelRequest(req *Request) {
	t.cancelRequest(req, errRequestCanceled)
}

// Cancel an in-flight request, recording the error value.
func (t *Transport) cancelRequest(req *Request, err error) {
	t.reqMu.Lock()
	cancel := t.reqCanceler[req]
	delete(t.reqCanceler, req)
	t.reqMu.Unlock()
	if cancel != nil {
		cancel(err)
	}
}

//
// Private implementation past this point.
//

var (
	// proxyConfigOnce guards proxyConfig
	envProxyOnce      sync.Once
	envProxyFuncValue func(*url.URL) (*url.URL, error)
)

// defaultProxyConfig returns a ProxyConfig value looked up
// from the environment. This mitigates expensive lookups
// on some platforms (e.g. Windows).
func envProxyFunc() func(*url.URL) (*url.URL, error) {
	envProxyOnce.Do(func() {
		envProxyFuncValue = httpproxy.FromEnvironment().ProxyFunc()
	})
	return envProxyFuncValue
}

// resetProxyConfig is used by tests.
func resetProxyConfig() {
	envProxyOnce = sync.Once{}
	envProxyFuncValue = nil
}

func (t *Transport) connectMethodForRequest(treq *transportRequest) (cm connectMethod, err error) {
	if port := treq.URL.Port(); !validPort(port) {
		return cm, fmt.Errorf("invalid URL port %q", port)
	}
	cm.targetScheme = treq.URL.Scheme
	cm.targetAddr = canonicalAddr(treq.URL)
	if t.Proxy != nil {
		cm.proxyURL, err = t.Proxy(treq.Request)
		if err == nil && cm.proxyURL != nil {
			if port := cm.proxyURL.Port(); !validPort(port) {
				return cm, fmt.Errorf("invalid proxy URL port %q", port)
			}
		}
	}
	return cm, err
}

// proxyAuth returns the Proxy-Authorization header to set
// on requests, if applicable.
func (cm *connectMethod) proxyAuth() string {
	if cm.proxyURL == nil {
		return ""
	}
	if u := cm.proxyURL.User; u != nil {
		username := u.Username()
		password, _ := u.Password()
		return "Basic " + basicAuth(username, password)
	}
	return ""
}

// error values for debugging and testing, not seen by users.
var (
	errKeepAlivesDisabled = errors.New("http: putIdleConn: keep alives disabled")
	errConnBroken         = errors.New("http: putIdleConn: connection is in bad state")
	errWantIdle           = errors.New("http: putIdleConn: CloseIdleConnections was called")
	errTooManyIdle        = errors.New("http: putIdleConn: too many idle connections")
	errTooManyIdleHost    = errors.New("http: putIdleConn: too many idle connections for host")
	errCloseIdleConns     = errors.New("http: CloseIdleConnections called")
	errReadLoopExiting    = errors.New("http: persistConn.readLoop exiting")
	errIdleConnTimeout    = errors.New("http: idle connection timeout")
	errNotCachingH2Conn   = errors.New("http: not caching alternate protocol's connections")

	// errServerClosedIdle is not seen by users for idempotent requests, but may be
	// seen by a user if the server shuts down an idle connection and sends its FIN
	// in flight with already-written POST body bytes from the client.
	// See https://github.com/golang/go/issues/19943#issuecomment-355607646
	errServerClosedIdle = errors.New("http: server closed idle connection")
)

// transportReadFromServerError is used by Transport.readLoop when the
// 1 byte peek read fails and we're actually anticipating a response.
// Usually this is just due to the inherent keep-alive shut down race,
// where the server closed the connection at the same time the client
// wrote. The underlying err field is usually io.EOF or some
// ECONNRESET sort of thing which varies by platform. But it might be
// the user's custom net.Conn.Read error too, so we carry it along for
// them to return from Transport.RoundTrip.
type transportReadFromServerError struct {
	err error
}

func (e transportReadFromServerError) Error() string {
	return fmt.Sprintf("net/http: Transport failed to read from server: %v", e.err)
}

func (t *Transport) putOrCloseIdleConn(pconn *persistConn) {
	if err := t.tryPutIdleConn(pconn); err != nil {
		pconn.close(err)
	}
}

func (t *Transport) maxIdleConnsPerHost() int {
	if v := t.MaxIdleConnsPerHost; v != 0 {
		return v
	}
	return DefaultMaxIdleConnsPerHost
}

// tryPutIdleConn adds pconn to the list of idle persistent connections awaiting
// a new request.
// If pconn is no longer needed or not in a good state, tryPutIdleConn returns
// an error explaining why it wasn't registered.
// tryPutIdleConn does not close pconn. Use putOrCloseIdleConn instead for that.
func (t *Transport) tryPutIdleConn(pconn *persistConn) error {
	if t.DisableKeepAlives || t.MaxIdleConnsPerHost < 0 {
		return errKeepAlivesDisabled
	}
	if pconn.isBroken() {
		return errConnBroken
	}
	if pconn.alt != nil {
		return errNotCachingH2Conn
	}
	pconn.markReused()
	key := pconn.cacheKey

	t.idleMu.Lock()
	defer t.idleMu.Unlock()

	waitingDialer := t.idleConnCh[key]
	select {
	case waitingDialer <- pconn:
		// We're done with this pconn and somebody else is
		// currently waiting for a conn of this type (they're
		// actively dialing, but this conn is ready
		// first). Chrome calls this socket late binding. See
		// https://insouciant.org/tech/connection-management-in-chromium/
		return nil
	default:
		if waitingDialer != nil {
			// They had populated this, but their dial won
			// first, so we can clean up this map entry.
			delete(t.idleConnCh, key)
		}
	}
	if t.wantIdle {
		return errWantIdle
	}
	if t.idleConn == nil {
		t.idleConn = make(map[connectMethodKey][]*persistConn)
	}
	idles := t.idleConn[key]
	if len(idles) >= t.maxIdleConnsPerHost() {
		return errTooManyIdleHost
	}
	for _, exist := range idles {
		if exist == pconn {
			log.Fatalf("dup idle pconn %p in freelist", pconn)
		}
	}
	t.idleConn[key] = append(idles, pconn)
	t.idleLRU.add(pconn)
	if t.MaxIdleConns != 0 && t.idleLRU.len() > t.MaxIdleConns {
		oldest := t.idleLRU.removeOldest()
		oldest.close(errTooManyIdle)
		t.removeIdleConnLocked(oldest)
	}
	if t.IdleConnTimeout > 0 {
		if pconn.idleTimer != nil {
			pconn.idleTimer.Reset(t.IdleConnTimeout)
		} else {
			pconn.idleTimer = time.AfterFunc(t.IdleConnTimeout, pconn.closeConnIfStillIdle)
		}
	}
	pconn.idleAt = time.Now()
	return nil
}

// getIdleConnCh returns a channel to receive and return idle
// persistent connection for the given connectMethod.
// It may return nil, if persistent connections are not being used.
func (t *Transport) getIdleConnCh(cm connectMethod) chan *persistConn {
	if t.DisableKeepAlives {
		return nil
	}
	key := cm.key()
	t.idleMu.Lock()
	defer t.idleMu.Unlock()
	t.wantIdle = false
	if t.idleConnCh == nil {
		t.idleConnCh = make(map[connectMethodKey]chan *persistConn)
	}
	ch, ok := t.idleConnCh[key]
	if !ok {
		ch = make(chan *persistConn)
		t.idleConnCh[key] = ch
	}
	return ch
}

func (t *Transport) getIdleConn(cm connectMethod) (pconn *persistConn, idleSince time.Time) {
	key := cm.key()
	t.idleMu.Lock()
	defer t.idleMu.Unlock()
	for {
		pconns, ok := t.idleConn[key]
		if !ok {
			return nil, time.Time{}
		}
		if len(pconns) == 1 {
			pconn = pconns[0]
			delete(t.idleConn, key)
		} else {
			// 2 or more cached connections; use the most
			// recently used one at the end.
			pconn = pconns[len(pconns)-1]
			t.idleConn[key] = pconns[:len(pconns)-1]
		}
		t.idleLRU.remove(pconn)
		if pconn.isBroken() {
			// There is a tiny window where this is
			// possible, between the connecting dying and
			// the persistConn readLoop calling
			// Transport.removeIdleConn. Just skip it and
			// carry on.
			continue
		}
		return pconn, pconn.idleAt
	}
}

// removeIdleConn marks pconn as dead.
func (t *Transport) removeIdleConn(pconn *persistConn) {
	t.idleMu.Lock()
	defer t.idleMu.Unlock()
	t.removeIdleConnLocked(pconn)
}

// t.idleMu must be held.
func (t *Transport) removeIdleConnLocked(pconn *persistConn) {
	if pconn.idleTimer != nil {
		pconn.idleTimer.Stop()
	}
	t.idleLRU.remove(pconn)
	key := pconn.cacheKey
	pconns := t.idleConn[key]
	switch len(pconns) {
	case 0:
		// Nothing
	case 1:
		if pconns[0] == pconn {
			delete(t.idleConn, key)
		}
	default:
		for i, v := range pconns {
			if v != pconn {
				continue
			}
			// Slide down, keeping most recently-used
			// conns at the end.
			copy(pconns[i:], pconns[i+1:])
			t.idleConn[key] = pconns[:len(pconns)-1]
			break
		}
	}
}

func (t *Transport) setReqCanceler(r *Request, fn func(error)) {
	t.reqMu.Lock()
	defer t.reqMu.Unlock()
	if t.reqCanceler == nil {
		t.reqCanceler = make(map[*Request]func(error))
	}
	if fn != nil {
		t.reqCanceler[r] = fn
	} else {
		delete(t.reqCanceler, r)
	}
}

// replaceReqCanceler replaces an existing cancel function. If there is no cancel function
// for the request, we don't set the function and return false.
// Since CancelRequest will clear the canceler, we can use the return value to detect if
// the request was canceled since the last setReqCancel call.
func (t *Transport) replaceReqCanceler(r *Request, fn func(error)) bool {
	t.reqMu.Lock()
	defer t.reqMu.Unlock()
	_, ok := t.reqCanceler[r]
	if !ok {
		return false
	}
	if fn != nil {
		t.reqCanceler[r] = fn
	} else {
		delete(t.reqCanceler, r)
	}
	return true
}

var zeroDialer net.Dialer

func (t *Transport) dial(ctx context.Context, network, addr string) (net.Conn, error) {
	if t.DialContext != nil {
		return t.DialContext(ctx, network, addr)
	}
	if t.Dial != nil {
		c, err := t.Dial(network, addr)
		if c == nil && err == nil {
			err = errors.New("net/http: Transport.Dial hook returned (nil, nil)")
		}
		return c, err
	}
	return zeroDialer.DialContext(ctx, network, addr)
}

// getConn dials and creates a new persistConn to the target as
// specified in the connectMethod. This includes doing a proxy CONNECT
// and/or setting up TLS.  If this doesn't return an error, the persistConn
// is ready to write requests to.
func (t *Transport) getConn(treq *transportRequest, cm connectMethod) (*persistConn, error) {
	req := treq.Request
	trace := treq.trace
	ctx := req.Context()
	if trace != nil && trace.GetConn != nil {
		trace.GetConn(cm.addr())
	}
	if pc, idleSince := t.getIdleConn(cm); pc != nil {
		if trace != nil && trace.GotConn != nil {
			trace.GotConn(pc.gotIdleConnTrace(idleSince))
		}
		// set request canceler to some non-nil function so we
		// can detect whether it was cleared between now and when
		// we enter roundTrip
		t.setReqCanceler(req, func(error) {})
		return pc, nil
	}

	type dialRes struct {
		pc  *persistConn
		err error
	}
	dialc := make(chan dialRes)
	cmKey := cm.key()

	// Copy these hooks so we don't race on the postPendingDial in
	// the goroutine we launch. Issue 11136.
	testHookPrePendingDial := testHookPrePendingDial
	testHookPostPendingDial := testHookPostPendingDial

	handlePendingDial := func() {
		testHookPrePendingDial()
		go func() {
			if v := <-dialc; v.err == nil {
				t.putOrCloseIdleConn(v.pc)
			} else {
				t.decHostConnCount(cmKey)
			}
			testHookPostPendingDial()
		}()
	}

	cancelc := make(chan error, 1)
	t.setReqCanceler(req, func(err error) { cancelc <- err })

	if t.MaxConnsPerHost > 0 {
		select {
		case <-t.incHostConnCount(cmKey):
			// count below conn per host limit; proceed
		case pc := <-t.getIdleConnCh(cm):
			if trace != nil && trace.GotConn != nil {
				trace.GotConn(httptrace.GotConnInfo{Conn: pc.conn, Reused: pc.isReused()})
			}
			return pc, nil
		case <-req.Cancel:
			return nil, errRequestCanceledConn
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case err := <-cancelc:
			if err == errRequestCanceled {
				err = errRequestCanceledConn
			}
			return nil, err
		}
	}

	go func() {
		pc, err := t.dialConn(ctx, cm)
		dialc <- dialRes{pc, err}
	}()

	idleConnCh := t.getIdleConnCh(cm)
	select {
	case v := <-dialc:
		// Our dial finished.
		if v.pc != nil {
			if trace != nil && trace.GotConn != nil && v.pc.alt == nil {
				trace.GotConn(httptrace.GotConnInfo{Conn: v.pc.conn})
			}
			return v.pc, nil
		}
		// Our dial failed. See why to return a nicer error
		// value.
		t.decHostConnCount(cmKey)
		select {
		case <-req.Cancel:
			// It was an error due to cancelation, so prioritize that
			// error value. (Issue 16049)
			return nil, errRequestCanceledConn
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case err := <-cancelc:
			if err == errRequestCanceled {
				err = errRequestCanceledConn
			}
			return nil, err
		default:
			// It wasn't an error due to cancelation, so
			// return the original error message:
			return nil, v.err
		}
	case pc := <-idleConnCh:
		// Another request finished first and its net.Conn
		// became available before our dial. Or somebody
		// else's dial that they didn't use.
		// But our dial is still going, so give it away
		// when it finishes:
		handlePendingDial()
		if trace != nil && trace.GotConn != nil {
			trace.GotConn(httptrace.GotConnInfo{Conn: pc.conn, Reused: pc.isReused()})
		}
		return pc, nil
	case <-req.Cancel:
		handlePendingDial()
		return nil, errRequestCanceledConn
	case <-req.Context().Done():
		handlePendingDial()
		return nil, req.Context().Err()
	case err := <-cancelc:
		handlePendingDial()
		if err == errRequestCanceled {
			err = errRequestCanceledConn
		}
		return nil, err
	}
}

// incHostConnCount increments the count of connections for a
// given host. It returns an already-closed channel if the count
// is not at its limit; otherwise it returns a channel which is
// notified when the count is below the limit.
func (t *Transport) incHostConnCount(cmKey connectMethodKey) <-chan struct{} {
	if t.MaxConnsPerHost <= 0 {
		return connsPerHostClosedCh
	}
	t.connCountMu.Lock()
	defer t.connCountMu.Unlock()
	if t.connPerHostCount[cmKey] == t.MaxConnsPerHost {
		if t.connPerHostAvailable == nil {
			t.connPerHostAvailable = make(map[connectMethodKey]chan struct{})
		}
		ch, ok := t.connPerHostAvailable[cmKey]
		if !ok {
			ch = make(chan struct{})
			t.connPerHostAvailable[cmKey] = ch
		}
		return ch
	}
	if t.connPerHostCount == nil {
		t.connPerHostCount = make(map[connectMethodKey]int)
	}
	t.connPerHostCount[cmKey]++
	// return a closed channel to avoid race: if decHostConnCount is called
	// after incHostConnCount and during the nil check, decHostConnCount
	// will delete the channel since it's not being listened on yet.
	return connsPerHostClosedCh
}

// decHostConnCount decrements the count of connections
// for a given host.
// See Transport.MaxConnsPerHost.
func (t *Transport) decHostConnCount(cmKey connectMethodKey) {
	if t.MaxConnsPerHost <= 0 {
		return
	}
	t.connCountMu.Lock()
	defer t.connCountMu.Unlock()
	t.connPerHostCount[cmKey]--
	select {
	case t.connPerHostAvailable[cmKey] <- struct{}{}:
	default:
		// close channel before deleting avoids getConn waiting forever in
		// case getConn has reference to channel but hasn't started waiting.
		// This could lead to more than MaxConnsPerHost in the unlikely case
		// that > 1 go routine has fetched the channel but none started waiting.
		if t.connPerHostAvailable[cmKey] != nil {
			close(t.connPerHostAvailable[cmKey])
		}
		delete(t.connPerHostAvailable, cmKey)
	}
	if t.connPerHostCount[cmKey] == 0 {
		delete(t.connPerHostCount, cmKey)
	}
}

// connCloseListener wraps a connection, the transport that dialed it
// and the connected-to host key so the host connection count can be
// transparently decremented by whatever closes the embedded connection.
type connCloseListener struct {
	net.Conn
	t        *Transport
	cmKey    connectMethodKey
	didClose int32
}

func (c *connCloseListener) Close() error {
	if atomic.AddInt32(&c.didClose, 1) != 1 {
		return nil
	}
	err := c.Conn.Close()
	c.t.decHostConnCount(c.cmKey)
	return err
}

// The connect method and the transport can both specify a TLS
// Host name.  The transport's name takes precedence if present.
func chooseTLSHost(cm connectMethod, t *Transport) string {
	tlsHost := ""
	if t.TLSClientConfig != nil {
		tlsHost = t.TLSClientConfig.ServerName
	}
	if tlsHost == "" {
		tlsHost = cm.tlsHost()
	}
	return tlsHost
}

// Add TLS to a persistent connection, i.e. negotiate a TLS session. If pconn is already a TLS
// tunnel, this function establishes a nested TLS session inside the encrypted channel.
// The remote endpoint's name may be overridden by TLSClientConfig.ServerName.
func (pconn *persistConn) addTLS(name string, trace *httptrace.ClientTrace) error {
	// Initiate TLS and check remote host name against certificate.
	cfg := cloneTLSConfig(pconn.t.TLSClientConfig)
	if cfg.ServerName == "" {
		cfg.ServerName = name
	}
	plainConn := pconn.conn
	tlsConn := tls.Client(plainConn, cfg)
	errc := make(chan error, 2)
	var timer *time.Timer // for canceling TLS handshake
	if d := pconn.t.TLSHandshakeTimeout; d != 0 {
		timer = time.AfterFunc(d, func() {
			errc <- tlsHandshakeTimeoutError{}
		})
	}
	go func() {
		if trace != nil && trace.TLSHandshakeStart != nil {
			trace.TLSHandshakeStart()
		}
		err := tlsConn.Handshake()
		if timer != nil {
			timer.Stop()
		}
		errc <- err
	}()
	if err := <-errc; err != nil {
		plainConn.Close()
		if trace != nil && trace.TLSHandshakeDone != nil {
			trace.TLSHandshakeDone(tls.ConnectionState{}, err)
		}
		return err
	}
	cs := tlsConn.ConnectionState()
	if trace != nil && trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(cs, nil)
	}
	pconn.tlsState = &cs
	pconn.conn = tlsConn
	return nil
}

func (t *Transport) dialConn(ctx context.Context, cm connectMethod) (*persistConn, error) {
	pconn := &persistConn{
		t:             t,
		cacheKey:      cm.key(),
		reqch:         make(chan requestAndChan, 1),
		writech:       make(chan writeRequest, 1),
		closech:       make(chan struct{}),
		writeErrCh:    make(chan error, 1),
		writeLoopDone: make(chan struct{}),
	}
	trace := httptrace.ContextClientTrace(ctx)
	wrapErr := func(err error) error {
		if cm.proxyURL != nil {
			// Return a typed error, per Issue 16997
			return &net.OpError{Op: "proxyconnect", Net: "tcp", Err: err}
		}
		return err
	}
	if cm.scheme() == "https" && t.DialTLS != nil {
		var err error
		pconn.conn, err = t.DialTLS("tcp", cm.addr())
		if err != nil {
			return nil, wrapErr(err)
		}
		if pconn.conn == nil {
			return nil, wrapErr(errors.New("net/http: Transport.DialTLS returned (nil, nil)"))
		}
		if tc, ok := pconn.conn.(*tls.Conn); ok {
			// Handshake here, in case DialTLS didn't. TLSNextProto below
			// depends on it for knowing the connection state.
			if trace != nil && trace.TLSHandshakeStart != nil {
				trace.TLSHandshakeStart()
			}
			if err := tc.Handshake(); err != nil {
				go pconn.conn.Close()
				if trace != nil && trace.TLSHandshakeDone != nil {
					trace.TLSHandshakeDone(tls.ConnectionState{}, err)
				}
				return nil, err
			}
			cs := tc.ConnectionState()
			if trace != nil && trace.TLSHandshakeDone != nil {
				trace.TLSHandshakeDone(cs, nil)
			}
			pconn.tlsState = &cs
		}
	} else {
		conn, err := t.dial(ctx, "tcp", cm.addr())
		if err != nil {
			return nil, wrapErr(err)
		}
		pconn.conn = conn
		if cm.scheme() == "https" {
			var firstTLSHost string
			if firstTLSHost, _, err = net.SplitHostPort(cm.addr()); err != nil {
				return nil, wrapErr(err)
			}
			if err = pconn.addTLS(firstTLSHost, trace); err != nil {
				return nil, wrapErr(err)
			}
		}
	}

	// Proxy setup.
	switch {
	case cm.proxyURL == nil:
		// Do nothing. Not using a proxy.
	case cm.proxyURL.Scheme == "socks5":
		conn := pconn.conn
		d := socksNewDialer("tcp", conn.RemoteAddr().String())
		if u := cm.proxyURL.User; u != nil {
			auth := &socksUsernamePassword{
				Username: u.Username(),
			}
			auth.Password, _ = u.Password()
			d.AuthMethods = []socksAuthMethod{
				socksAuthMethodNotRequired,
				socksAuthMethodUsernamePassword,
			}
			d.Authenticate = auth.Authenticate
		}
		if _, err := d.DialWithConn(ctx, conn, "tcp", cm.targetAddr); err != nil {
			conn.Close()
			return nil, err
		}
	case cm.targetScheme == "http":
		pconn.isProxy = true
		if pa := cm.proxyAuth(); pa != "" {
			pconn.mutateHeaderFunc = func(h Header) {
				h.Set("Proxy-Authorization", pa)
			}
		}
	case cm.targetScheme == "https":
		conn := pconn.conn
		hdr := t.ProxyConnectHeader
		if hdr == nil {
			hdr = make(Header)
		}
		connectReq := &Request{
			Method: "CONNECT",
			URL:    &url.URL{Opaque: cm.targetAddr},
			Host:   cm.targetAddr,
			Header: hdr,
		}
		if pa := cm.proxyAuth(); pa != "" {
			connectReq.Header.Set("Proxy-Authorization", pa)
		}
		connectReq.Write(conn)

		// Read response.
		// Okay to use and discard buffered reader here, because
		// TLS server will not speak until spoken to.
		br := bufio.NewReader(conn)
		resp, err := ReadResponse(br, connectReq)
		if err != nil {
			conn.Close()
			return nil, err
		}
		if resp.StatusCode != 200 {
			f := strings.SplitN(resp.Status, " ", 2)
			conn.Close()
			if len(f) < 2 {
				return nil, errors.New("unknown status code")
			}
			return nil, errors.New(f[1])
		}
	}

	if cm.proxyURL != nil && cm.targetScheme == "https" {
		if err := pconn.addTLS(cm.tlsHost(), trace); err != nil {
			return nil, err
		}
	}

	if s := pconn.tlsState; s != nil && s.NegotiatedProtocolIsMutual && s.NegotiatedProtocol != "" {
		if next, ok := t.TLSNextProto[s.NegotiatedProtocol]; ok {
			return &persistConn{alt: next(cm.targetAddr, pconn.conn.(*tls.Conn))}, nil
		}
	}

	if t.MaxConnsPerHost > 0 {
		pconn.conn = &connCloseListener{Conn: pconn.conn, t: t, cmKey: pconn.cacheKey}
	}
	pconn.br = bufio.NewReader(pconn)
	pconn.bw = bufio.NewWriter(persistConnWriter{pconn})
	go pconn.readLoop()
	go pconn.writeLoop()
	return pconn, nil
}

// persistConnWriter is the io.Writer written to by pc.bw.
// It accumulates the number of bytes written to the underlying conn,
// so the retry logic can determine whether any bytes made it across
// the wire.
// This is exactly 1 pointer field wide so it can go into an interface
// without allocation.
type persistConnWriter struct {
	pc *persistConn
}

func (w persistConnWriter) Write(p []byte) (n int, err error) {
	n, err = w.pc.conn.Write(p)
	w.pc.nwrite += int64(n)
	return
}

// connectMethod is the map key (in its String form) for keeping persistent
// TCP connections alive for subsequent HTTP requests.
//
// A connect method may be of the following types:
//
//	Cache key form                    Description
//	-----------------                 -------------------------
//	|http|foo.com                     http directly to server, no proxy
//	|https|foo.com                    https directly to server, no proxy
//	http://proxy.com|https|foo.com    http to proxy, then CONNECT to foo.com
//	http://proxy.com|http             http to proxy, http to anywhere after that
//	socks5://proxy.com|http|foo.com   socks5 to proxy, then http to foo.com
//	socks5://proxy.com|https|foo.com  socks5 to proxy, then https to foo.com
//	https://proxy.com|https|foo.com   https to proxy, then CONNECT to foo.com
//	https://proxy.com|http            https to proxy, http to anywhere after that
//
type connectMethod struct {
	proxyURL     *url.URL // nil for no proxy, else full proxy URL
	targetScheme string   // "http" or "https"
	// If proxyURL specifies an http or https proxy, and targetScheme is http (not https),
	// then targetAddr is not included in the connect method key, because the socket can
	// be reused for different targetAddr values.
	targetAddr string
}

func (cm *connectMethod) key() connectMethodKey {
	proxyStr := ""
	targetAddr := cm.targetAddr
	if cm.proxyURL != nil {
		proxyStr = cm.proxyURL.String()
		if (cm.proxyURL.Scheme == "http" || cm.proxyURL.Scheme == "https") && cm.targetScheme == "http" {
			targetAddr = ""
		}
	}
	return connectMethodKey{
		proxy:  proxyStr,
		scheme: cm.targetScheme,
		addr:   targetAddr,
	}
}

// scheme returns the first hop scheme: http, https, or socks5
func (cm *connectMethod) scheme() string {
	if cm.proxyURL != nil {
		return cm.proxyURL.Scheme
	}
	return cm.targetScheme
}

// addr returns the first hop "host:port" to which we need to TCP connect.
func (cm *connectMethod) addr() string {
	if cm.proxyURL != nil {
		return canonicalAddr(cm.proxyURL)
	}
	return cm.targetAddr
}

// tlsHost returns the host name to match against the peer's
// TLS certificate.
func (cm *connectMethod) tlsHost() string {
	h := cm.targetAddr
	if hasPort(h) {
		h = h[:strings.LastIndex(h, ":")]
	}
	return h
}

// connectMethodKey is the map key version of connectMethod, with a
// stringified proxy URL (or the empty string) instead of a pointer to
// a URL.
type connectMethodKey struct {
	proxy, scheme, addr string
}

func (k connectMethodKey) String() string {
	// Only used by tests.
	return fmt.Sprintf("%s|%s|%s", k.proxy, k.scheme, k.addr)
}

// persistConn wraps a connection, usually a persistent one
// (but may be used for non-keep-alive requests as well)
type persistConn struct {
	// alt optionally specifies the TLS NextProto RoundTripper.
	// This is used for HTTP/2 today and future protocols later.
	// If it's non-nil, the rest of the fields are unused.
	alt RoundTripper

	t         *Transport
	cacheKey  connectMethodKey
	conn      net.Conn
	tlsState  *tls.ConnectionState
	br        *bufio.Reader       // from conn
	bw        *bufio.Writer       // to conn
	nwrite    int64               // bytes written
	reqch     chan requestAndChan // written by roundTrip; read by readLoop
	writech   chan writeRequest   // written by roundTrip; read by writeLoop
	closech   chan struct{}       // closed when conn closed
	isProxy   bool
	sawEOF    bool  // whether we've seen EOF from conn; owned by readLoop
	readLimit int64 // bytes allowed to be read; owned by readLoop
	// writeErrCh passes the request write error (usually nil)
	// from the writeLoop goroutine to the readLoop which passes
	// it off to the res.Body reader, which then uses it to decide
	// whether or not a connection can be reused. Issue 7569.
	writeErrCh chan error

	writeLoopDone chan struct{} // closed when write loop ends

	// Both guarded by Transport.idleMu:
	idleAt    time.Time   // time it last become idle
	idleTimer *time.Timer // holding an AfterFunc to close it

	mu                   sync.Mutex // guards following fields
	numExpectedResponses int
	closed               error // set non-nil when conn is closed, before closech is closed
	canceledErr          error // set non-nil if conn is canceled
	broken               bool  // an error has happened on this connection; marked broken so it's not reused.
	reused               bool  // whether conn has had successful request/response and is being reused.
	// mutateHeaderFunc is an optional func to modify extra
	// headers on each outbound request before it's written. (the
	// original Request given to RoundTrip is not modified)
	mutateHeaderFunc func(Header)
}

func (pc *persistConn) maxHeaderResponseSize() int64 {
	if v := pc.t.MaxResponseHeaderBytes; v != 0 {
		return v
	}
	return 10 << 20 // conservative default; same as http2
}

func (pc *persistConn) Read(p []byte) (n int, err error) {
	if pc.readLimit <= 0 {
		return 0, fmt.Errorf("read limit of %d bytes exhausted", pc.maxHeaderResponseSize())
	}
	if int64(len(p)) > pc.readLimit {
		p = p[:pc.readLimit]
	}
	n, err = pc.conn.Read(p)
	if err == io.EOF {
		pc.sawEOF = true
	}
	pc.readLimit -= int64(n)
	return
}

// isBroken reports whether this connection is in a known broken state.
func (pc *persistConn) isBroken() bool {
	pc.mu.Lock()
	b := pc.closed != nil
	pc.mu.Unlock()
	return b
}

// canceled returns non-nil if the connection was closed due to
// CancelRequest or due to context cancelation.
func (pc *persistConn) canceled() error {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	return pc.canceledErr
}

// isReused reports whether this connection is in a known broken state.
func (pc *persistConn) isReused() bool {
	pc.mu.Lock()
	r := pc.reused
	pc.mu.Unlock()
	return r
}

func (pc *persistConn) gotIdleConnTrace(idleAt time.Time) (t httptrace.GotConnInfo) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	t.Reused = pc.reused
	t.Conn = pc.conn
	t.WasIdle = true
	if !idleAt.IsZero() {
		t.IdleTime = time.Since(idleAt)
	}
	return
}

func (pc *persistConn) cancelRequest(err error) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.canceledErr = err
	pc.closeLocked(errRequestCanceled)
}

// closeConnIfStillIdle closes the connection if it's still sitting idle.
// This is what's called by the persistConn's idleTimer, and is run in its
// own goroutine.
func (pc *persistConn) closeConnIfStillIdle() {
	t := pc.t
	t.idleMu.Lock()
	defer t.idleMu.Unlock()
	if _, ok := t.idleLRU.m[pc]; !ok {
		// Not idle.
		return
	}
	t.removeIdleConnLocked(pc)
	pc.close(errIdleConnTimeout)
}

// mapRoundTripError returns the appropriate error value for
// persistConn.roundTrip.
//
// The provided err is the first error that (*persistConn).roundTrip
// happened to receive from its select statement.
//
// The startBytesWritten value should be the value of pc.nwrite before the roundTrip
// started writing the request.
func (pc *persistConn) mapRoundTripError(req *transportRequest, startBytesWritten int64, err error) error {
	if err == nil {
		return nil
	}

	// If the request was canceled, that's better than network
	// failures that were likely the result of tearing down the
	// connection.
	if cerr := pc.canceled(); cerr != nil {
		return cerr
	}

	// See if an error was set explicitly.
	req.mu.Lock()
	reqErr := req.err
	req.mu.Unlock()
	if reqErr != nil {
		return reqErr
	}

	if err == errServerClosedIdle {
		// Don't decorate
		return err
	}

	if _, ok := err.(transportReadFromServerError); ok {
		// Don't decorate
		return err
	}
	if pc.isBroken() {
		<-pc.writeLoopDone
		if pc.nwrite == startBytesWritten {
			return nothingWrittenError{err}
		}
		return fmt.Errorf("net/http: HTTP/1.x transport connection broken: %v", err)
	}
	return err
}

// errCallerOwnsConn is an internal sentinel error used when we hand
// off a writable response.Body to the caller. We use this to prevent
// closing a net.Conn that is now owned by the caller.
var errCallerOwnsConn = errors.New("read loop ending; caller owns writable underlying conn")

func (pc *persistConn) readLoop() {
	closeErr := errReadLoopExiting // default value, if not changed below
	defer func() {
		pc.close(closeErr)
		pc.t.removeIdleConn(pc)
	}()

	tryPutIdleConn := func(trace *httptrace.ClientTrace) bool {
		if err := pc.t.tryPutIdleConn(pc); err != nil {
			closeErr = err
			if trace != nil && trace.PutIdleConn != nil && err != errKeepAlivesDisabled {
				trace.PutIdleConn(err)
			}
			return false
		}
		if trace != nil && trace.PutIdleConn != nil {
			trace.PutIdleConn(nil)
		}
		return true
	}

	// eofc is used to block caller goroutines reading from Response.Body
	// at EOF until this goroutines has (potentially) added the connection
	// back to the idle pool.
	eofc := make(chan struct{})
	defer close(eofc) // unblock reader on errors

	// Read this once, before loop starts. (to avoid races in tests)
	testHookMu.Lock()
	testHookReadLoopBeforeNextRead := testHookReadLoopBeforeNextRead
	testHookMu.Unlock()

	alive := true
	for alive {
		pc.readLimit = pc.maxHeaderResponseSize()
		_, err := pc.br.Peek(1)

		pc.mu.Lock()
		if pc.numExpectedResponses == 0 {
			pc.readLoopPeekFailLocked(err)
			pc.mu.Unlock()
			return
		}
		pc.mu.Unlock()

		rc := <-pc.reqch
		trace := httptrace.ContextClientTrace(rc.req.Context())

		var resp *Response
		if err == nil {
			resp, err = pc.readResponse(rc, trace)
		} else {
			err = transportReadFromServerError{err}
			closeErr = err
		}

		if err != nil {
			if pc.readLimit <= 0 {
				err = fmt.Errorf("net/http: server response headers exceeded %d bytes; aborted", pc.maxHeaderResponseSize())
			}

			select {
			case rc.ch <- responseAndError{err: err}:
			case <-rc.callerGone:
				return
			}
			return
		}
		pc.readLimit = maxInt64 // effictively no limit for response bodies

		pc.mu.Lock()
		pc.numExpectedResponses--
		pc.mu.Unlock()

		bodyWritable := resp.bodyIsWritable()
		hasBody := rc.req.Method != "HEAD" && resp.ContentLength != 0

		if resp.Close || rc.req.Close || resp.StatusCode <= 199 || bodyWritable {
			// Don't do keep-alive on error if either party requested a close
			// or we get an unexpected informational (1xx) response.
			// StatusCode 100 is already handled above.
			alive = false
		}

		if !hasBody {
			pc.t.setReqCanceler(rc.req, nil)

			// Put the idle conn back into the pool before we send the response
			// so if they process it quickly and make another request, they'll
			// get this same conn. But we use the unbuffered channel 'rc'
			// to guarantee that persistConn.roundTrip got out of its select
			// potentially waiting for this persistConn to close.
			// but after
			alive = alive &&
				!pc.sawEOF &&
				pc.wroteRequest() &&
				tryPutIdleConn(trace)

			if bodyWritable {
				closeErr = errCallerOwnsConn
			}

			select {
			case rc.ch <- responseAndError{res: resp}:
			case <-rc.callerGone:
				return
			}

			// Now that they've read from the unbuffered channel, they're safely
			// out of the select that also waits on this goroutine to die, so
			// we're allowed to exit now if needed (if alive is false)
			testHookReadLoopBeforeNextRead()
			continue
		}

		waitForBodyRead := make(chan bool, 2)
		body := &bodyEOFSignal{
			body: resp.Body,
			earlyCloseFn: func() error {
				waitForBodyRead <- false
				<-eofc // will be closed by deferred call at the end of the function
				return nil

			},
			fn: func(err error) error {
				isEOF := err == io.EOF
				waitForBodyRead <- isEOF
				if isEOF {
					<-eofc // see comment above eofc declaration
				} else if err != nil {
					if cerr := pc.canceled(); cerr != nil {
						return cerr
					}
				}
				return err
			},
		}

		resp.Body = body
		if rc.addedGzip && strings.EqualFold(resp.Header.Get("Content-Encoding"), "gzip") {
			resp.Body = &gzipReader{body: body}
			resp.Header.Del("Content-Encoding")
			resp.Header.Del("Content-Length")
			resp.ContentLength = -1
			resp.Uncompressed = true
		}

		select {
		case rc.ch <- responseAndError{res: resp}:
		case <-rc.callerGone:
			return
		}

		// Before looping back to the top of this function and peeking on
		// the bufio.Reader, wait for the caller goroutine to finish
		// reading the response body. (or for cancelation or death)
		select {
		case bodyEOF := <-waitForBodyRead:
			pc.t.setReqCanceler(rc.req, nil) // before pc might return to idle pool
			alive = alive &&
				bodyEOF &&
				!pc.sawEOF &&
				pc.wroteRequest() &&
				tryPutIdleConn(trace)
			if bodyEOF {
				eofc <- struct{}{}
			}
		case <-rc.req.Cancel:
			alive = false
			pc.t.CancelRequest(rc.req)
		case <-rc.req.Context().Done():
			alive = false
			pc.t.cancelRequest(rc.req, rc.req.Context().Err())
		case <-pc.closech:
			alive = false
		}

		testHookReadLoopBeforeNextRead()
	}
}

func (pc *persistConn) readLoopPeekFailLocked(peekErr error) {
	if pc.closed != nil {
		return
	}
	if n := pc.br.Buffered(); n > 0 {
		buf, _ := pc.br.Peek(n)
		log.Printf("Unsolicited response received on idle HTTP channel starting with %q; err=%v", buf, peekErr)
	}
	if peekErr == io.EOF {
		// common case.
		pc.closeLocked(errServerClosedIdle)
	} else {
		pc.closeLocked(fmt.Errorf("readLoopPeekFailLocked: %v", peekErr))
	}
}

// readResponse reads an HTTP response (or two, in the case of "Expect:
// 100-continue") from the server. It returns the final non-100 one.
// trace is optional.
func (pc *persistConn) readResponse(rc requestAndChan, trace *httptrace.ClientTrace) (resp *Response, err error) {
	if trace != nil && trace.GotFirstResponseByte != nil {
		if peek, err := pc.br.Peek(1); err == nil && len(peek) == 1 {
			trace.GotFirstResponseByte()
		}
	}
	num1xx := 0               // number of informational 1xx headers received
	const max1xxResponses = 5 // arbitrary bound on number of informational responses

	continueCh := rc.continueCh
	for {
		resp, err = ReadResponse(pc.br, rc.req)
		if err != nil {
			return
		}
		resCode := resp.StatusCode
		if continueCh != nil {
			if resCode == 100 {
				if trace != nil && trace.Got100Continue != nil {
					trace.Got100Continue()
				}
				continueCh <- struct{}{}
				continueCh = nil
			} else if resCode >= 200 {
				close(continueCh)
				continueCh = nil
			}
		}
		is1xx := 100 <= resCode && resCode <= 199
		// treat 101 as a terminal status, see issue 26161
		is1xxNonTerminal := is1xx && resCode != StatusSwitchingProtocols
		if is1xxNonTerminal {
			num1xx++
			if num1xx > max1xxResponses {
				return nil, errors.New("net/http: too many 1xx informational responses")
			}
			pc.readLimit = pc.maxHeaderResponseSize() // reset the limit
			if trace != nil && trace.Got1xxResponse != nil {
				if err := trace.Got1xxResponse(resCode, textproto.MIMEHeader(resp.Header)); err != nil {
					return nil, err
				}
			}
			continue
		}
		break
	}
	if resp.isProtocolSwitch() {
		resp.Body = newReadWriteCloserBody(pc.br, pc.conn)
	}

	resp.TLS = pc.tlsState
	return
}

// waitForContinue returns the function to block until
// any response, timeout or connection close. After any of them,
// the function returns a bool which indicates if the body should be sent.
func (pc *persistConn) waitForContinue(continueCh <-chan struct{}) func() bool {
	if continueCh == nil {
		return nil
	}
	return func() bool {
		timer := time.NewTimer(pc.t.ExpectContinueTimeout)
		defer timer.Stop()

		select {
		case _, ok := <-continueCh:
			return ok
		case <-timer.C:
			return true
		case <-pc.closech:
			return false
		}
	}
}

func newReadWriteCloserBody(br *bufio.Reader, rwc io.ReadWriteCloser) io.ReadWriteCloser {
	body := &readWriteCloserBody{ReadWriteCloser: rwc}
	if br.Buffered() != 0 {
		body.br = br
	}
	return body
}

// readWriteCloserBody is the Response.Body type used when we want to
// give users write access to the Body through the underlying
// connection (TCP, unless using custom dialers). This is then
// the concrete type for a Response.Body on the 101 Switching
// Protocols response, as used by WebSockets, h2c, etc.
type readWriteCloserBody struct {
	br *bufio.Reader // used until empty
	io.ReadWriteCloser
}

func (b *readWriteCloserBody) Read(p []byte) (n int, err error) {
	if b.br != nil {
		if n := b.br.Buffered(); len(p) > n {
			p = p[:n]
		}
		n, err = b.br.Read(p)
		if b.br.Buffered() == 0 {
			b.br = nil
		}
		return n, err
	}
	return b.ReadWriteCloser.Read(p)
}

// nothingWrittenError wraps a write errors which ended up writing zero bytes.
type nothingWrittenError struct {
	error
}

func (pc *persistConn) writeLoop() {
	defer close(pc.writeLoopDone)
	for {
		select {
		case wr := <-pc.writech:
			startBytesWritten := pc.nwrite
			err := wr.req.Request.write(pc.bw, pc.isProxy, wr.req.extra, pc.waitForContinue(wr.continueCh))
			if bre, ok := err.(requestBodyReadError); ok {
				err = bre.error
				// Errors reading from the user's
				// Request.Body are high priority.
				// Set it here before sending on the
				// channels below or calling
				// pc.close() which tears town
				// connections and causes other
				// errors.
				wr.req.setError(err)
			}
			if err == nil {
				err = pc.bw.Flush()
			}
			if err != nil {
				wr.req.Request.closeBody()
				if pc.nwrite == startBytesWritten {
					err = nothingWrittenError{err}
				}
			}
			pc.writeErrCh <- err // to the body reader, which might recycle us
			wr.ch <- err         // to the roundTrip function
			if err != nil {
				pc.close(err)
				return
			}
		case <-pc.closech:
			return
		}
	}
}

// maxWriteWaitBeforeConnReuse is how long the a Transport RoundTrip
// will wait to see the Request's Body.Write result after getting a
// response from the server. See comments in (*persistConn).wroteRequest.
const maxWriteWaitBeforeConnReuse = 50 * time.Millisecond

// wroteRequest is a check before recycling a connection that the previous write
// (from writeLoop above) happened and was successful.
func (pc *persistConn) wroteRequest() bool {
	select {
	case err := <-pc.writeErrCh:
		// Common case: the write happened well before the response, so
		// avoid creating a timer.
		return err == nil
	default:
		// Rare case: the request was written in writeLoop above but
		// before it could send to pc.writeErrCh, the reader read it
		// all, processed it, and called us here. In this case, give the
		// write goroutine a bit of time to finish its send.
		//
		// Less rare case: We also get here in the legitimate case of
		// Issue 7569, where the writer is still writing (or stalled),
		// but the server has already replied. In this case, we don't
		// want to wait too long, and we want to return false so this
		// connection isn't re-used.
		select {
		case err := <-pc.writeErrCh:
			return err == nil
		case <-time.After(maxWriteWaitBeforeConnReuse):
			return false
		}
	}
}

// responseAndError is how the goroutine reading from an HTTP/1 server
// communicates with the goroutine doing the RoundTrip.
type responseAndError struct {
	res *Response // else use this response (see res method)
	err error
}

type requestAndChan struct {
	req *Request
	ch  chan responseAndError // unbuffered; always send in select on callerGone

	// whether the Transport (as opposed to the user client code)
	// added the Accept-Encoding gzip header. If the Transport
	// set it, only then do we transparently decode the gzip.
	addedGzip bool

	// Optional blocking chan for Expect: 100-continue (for send).
	// If the request has an "Expect: 100-continue" header and
	// the server responds 100 Continue, readLoop send a value
	// to writeLoop via this chan.
	continueCh chan<- struct{}

	callerGone <-chan struct{} // closed when roundTrip caller has returned
}

// A writeRequest is sent by the readLoop's goroutine to the
// writeLoop's goroutine to write a request while the read loop
// concurrently waits on both the write response and the server's
// reply.
type writeRequest struct {
	req *transportRequest
	ch  chan<- error

	// Optional blocking chan for Expect: 100-continue (for receive).
	// If not nil, writeLoop blocks sending request body until
	// it receives from this chan.
	continueCh <-chan struct{}
}

type httpError struct {
	err     string
	timeout bool
}

func (e *httpError) Error() string   { return e.err }
func (e *httpError) Timeout() bool   { return e.timeout }
func (e *httpError) Temporary() bool { return true }

var errTimeout error = &httpError{err: "net/http: timeout awaiting response headers", timeout: true}
var errRequestCanceled = errors.New("net/http: request canceled")
var errRequestCanceledConn = errors.New("net/http: request canceled while waiting for connection") // TODO: unify?

func nop() {}

// testHooks. Always non-nil.
var (
	testHookEnterRoundTrip   = nop
	testHookWaitResLoop      = nop
	testHookRoundTripRetried = nop
	testHookPrePendingDial   = nop
	testHookPostPendingDial  = nop

	testHookMu                     sync.Locker = fakeLocker{} // guards following
	testHookReadLoopBeforeNextRead             = nop
)

func (pc *persistConn) roundTrip(req *transportRequest) (resp *Response, err error) {
	testHookEnterRoundTrip()
	if !pc.t.replaceReqCanceler(req.Request, pc.cancelRequest) {
		pc.t.putOrCloseIdleConn(pc)
		return nil, errRequestCanceled
	}
	pc.mu.Lock()
	pc.numExpectedResponses++
	headerFn := pc.mutateHeaderFunc
	pc.mu.Unlock()

	if headerFn != nil {
		headerFn(req.extraHeaders())
	}

	// Ask for a compressed version if the caller didn't set their
	// own value for Accept-Encoding. We only attempt to
	// uncompress the gzip stream if we were the layer that
	// requested it.
	requestedGzip := false
	if !pc.t.DisableCompression &&
		req.Header.Get("Accept-Encoding") == "" &&
		req.Header.Get("Range") == "" &&
		req.Method != "HEAD" {
		// Request gzip only, not deflate. Deflate is ambiguous and
		// not as universally supported anyway.
		// See: http://www.gzip.org/zlib/zlib_faq.html#faq38
		//
		// Note that we don't request this for HEAD requests,
		// due to a bug in nginx:
		//   https://trac.nginx.org/nginx/ticket/358
		//   https://golang.org/issue/5522
		//
		// We don't request gzip if the request is for a range, since
		// auto-decoding a portion of a gzipped document will just fail
		// anyway. See https://golang.org/issue/8923
		requestedGzip = true
		req.extraHeaders().Set("Accept-Encoding", "gzip")
	}

	var continueCh chan struct{}
	if req.ProtoAtLeast(1, 1) && req.Body != nil && req.expectsContinue() {
		continueCh = make(chan struct{}, 1)
	}

	if pc.t.DisableKeepAlives {
		req.extraHeaders().Set("Connection", "close")
	}

	gone := make(chan struct{})
	defer close(gone)

	defer func() {
		if err != nil {
			pc.t.setReqCanceler(req.Request, nil)
		}
	}()

	const debugRoundTrip = false

	// Write the request concurrently with waiting for a response,
	// in case the server decides to reply before reading our full
	// request body.
	startBytesWritten := pc.nwrite
	writeErrCh := make(chan error, 1)
	pc.writech <- writeRequest{req, writeErrCh, continueCh}

	resc := make(chan responseAndError)
	pc.reqch <- requestAndChan{
		req:        req.Request,
		ch:         resc,
		addedGzip:  requestedGzip,
		continueCh: continueCh,
		callerGone: gone,
	}

	var respHeaderTimer <-chan time.Time
	cancelChan := req.Request.Cancel
	ctxDoneChan := req.Context().Done()
	for {
		testHookWaitResLoop()
		select {
		case err := <-writeErrCh:
			if debugRoundTrip {
				req.logf("writeErrCh resv: %T/%#v", err, err)
			}
			if err != nil {
				pc.close(fmt.Errorf("write error: %v", err))
				return nil, pc.mapRoundTripError(req, startBytesWritten, err)
			}
			if d := pc.t.ResponseHeaderTimeout; d > 0 {
				if debugRoundTrip {
					req.logf("starting timer for %v", d)
				}
				timer := time.NewTimer(d)
				defer timer.Stop() // prevent leaks
				respHeaderTimer = timer.C
			}
		case <-pc.closech:
			if debugRoundTrip {
				req.logf("closech recv: %T %#v", pc.closed, pc.closed)
			}
			return nil, pc.mapRoundTripError(req, startBytesWritten, pc.closed)
		case <-respHeaderTimer:
			if debugRoundTrip {
				req.logf("timeout waiting for response headers.")
			}
			pc.close(errTimeout)
			return nil, errTimeout
		case re := <-resc:
			if (re.res == nil) == (re.err == nil) {
				panic(fmt.Sprintf("internal error: exactly one of res or err should be set; nil=%v", re.res == nil))
			}
			if debugRoundTrip {
				req.logf("resc recv: %p, %T/%#v", re.res, re.err, re.err)
			}
			if re.err != nil {
				return nil, pc.mapRoundTripError(req, startBytesWritten, re.err)
			}
			return re.res, nil
		case <-cancelChan:
			pc.t.CancelRequest(req.Request)
			cancelChan = nil
		case <-ctxDoneChan:
			pc.t.cancelRequest(req.Request, req.Context().Err())
			cancelChan = nil
			ctxDoneChan = nil
		}
	}
}

// tLogKey is a context WithValue key for test debugging contexts containing
// a t.Logf func. See export_test.go's Request.WithT method.
type tLogKey struct{}

func (tr *transportRequest) logf(format string, args ...interface{}) {
	if logf, ok := tr.Request.Context().Value(tLogKey{}).(func(string, ...interface{})); ok {
		logf(time.Now().Format(time.RFC3339Nano)+": "+format, args...)
	}
}

// markReused marks this connection as having been successfully used for a
// request and response.
func (pc *persistConn) markReused() {
	pc.mu.Lock()
	pc.reused = true
	pc.mu.Unlock()
}

// close closes the underlying TCP connection and closes
// the pc.closech channel.
//
// The provided err is only for testing and debugging; in normal
// circumstances it should never be seen by users.
func (pc *persistConn) close(err error) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.closeLocked(err)
}

func (pc *persistConn) closeLocked(err error) {
	if err == nil {
		panic("nil error")
	}
	pc.broken = true
	if pc.closed == nil {
		pc.closed = err
		if pc.alt != nil {
			// Do nothing; can only get here via getConn's
			// handlePendingDial's putOrCloseIdleConn when
			// it turns out the abandoned connection in
			// flight ended up negotiating an alternate
			// protocol. We don't use the connection
			// freelist for http2. That's done by the
			// alternate protocol's RoundTripper.
		} else {
			if err != errCallerOwnsConn {
				pc.conn.Close()
			}
			close(pc.closech)
		}
	}
	pc.mutateHeaderFunc = nil
}

var portMap = map[string]string{
	"http":   "80",
	"https":  "443",
	"socks5": "1080",
}

// canonicalAddr returns url.Host but always with a ":port" suffix
func canonicalAddr(url *url.URL) string {
	addr := url.Hostname()
	if v, err := idnaASCII(addr); err == nil {
		addr = v
	}
	port := url.Port()
	if port == "" {
		port = portMap[url.Scheme]
	}
	return net.JoinHostPort(addr, port)
}

// bodyEOFSignal is used by the HTTP/1 transport when reading response
// bodies to make sure we see the end of a response body before
// proceeding and reading on the connection again.
//
// It wraps a ReadCloser but runs fn (if non-nil) at most
// once, right before its final (error-producing) Read or Close call
// returns. fn should return the new error to return from Read or Close.
//
// If earlyCloseFn is non-nil and Close is called before io.EOF is
// seen, earlyCloseFn is called instead of fn, and its return value is
// the return value from Close.
type bodyEOFSignal struct {
	body         io.ReadCloser
	mu           sync.Mutex        // guards following 4 fields
	closed       bool              // whether Close has been called
	rerr         error             // sticky Read error
	fn           func(error) error // err will be nil on Read io.EOF
	earlyCloseFn func() error      // optional alt Close func used if io.EOF not seen
}

var errReadOnClosedResBody = errors.New("http: read on closed response body")

func (es *bodyEOFSignal) Read(p []byte) (n int, err error) {
	es.mu.Lock()
	closed, rerr := es.closed, es.rerr
	es.mu.Unlock()
	if closed {
		return 0, errReadOnClosedResBody
	}
	if rerr != nil {
		return 0, rerr
	}

	n, err = es.body.Read(p)
	if err != nil {
		es.mu.Lock()
		defer es.mu.Unlock()
		if es.rerr == nil {
			es.rerr = err
		}
		err = es.condfn(err)
	}
	return
}

func (es *bodyEOFSignal) Close() error {
	es.mu.Lock()
	defer es.mu.Unlock()
	if es.closed {
		return nil
	}
	es.closed = true
	if es.earlyCloseFn != nil && es.rerr != io.EOF {
		return es.earlyCloseFn()
	}
	err := es.body.Close()
	return es.condfn(err)
}

// caller must hold es.mu.
func (es *bodyEOFSignal) condfn(err error) error {
	if es.fn == nil {
		return err
	}
	err = es.fn(err)
	es.fn = nil
	return err
}

// gzipReader wraps a response body so it can lazily
// call gzip.NewReader on the first call to Read
type gzipReader struct {
	body *bodyEOFSignal // underlying HTTP/1 response body framing
	zr   *gzip.Reader   // lazily-initialized gzip reader
	zerr error          // any error from gzip.NewReader; sticky
}

func (gz *gzipReader) Read(p []byte) (n int, err error) {
	if gz.zr == nil {
		if gz.zerr == nil {
			gz.zr, gz.zerr = gzip.NewReader(gz.body)
		}
		if gz.zerr != nil {
			return 0, gz.zerr
		}
	}

	gz.body.mu.Lock()
	if gz.body.closed {
		err = errReadOnClosedResBody
	}
	gz.body.mu.Unlock()

	if err != nil {
		return 0, err
	}
	return gz.zr.Read(p)
}

func (gz *gzipReader) Close() error {
	return gz.body.Close()
}

type readerAndCloser struct {
	io.Reader
	io.Closer
}

type tlsHandshakeTimeoutError struct{}

func (tlsHandshakeTimeoutError) Timeout() bool   { return true }
func (tlsHandshakeTimeoutError) Temporary() bool { return true }
func (tlsHandshakeTimeoutError) Error() string   { return "net/http: TLS handshake timeout" }

// fakeLocker is a sync.Locker which does nothing. It's used to guard
// test-only fields when not under test, to avoid runtime atomic
// overhead.
type fakeLocker struct{}

func (fakeLocker) Lock()   {}
func (fakeLocker) Unlock() {}

// clneTLSConfig returns a shallow clone of cfg, or a new zero tls.Config if
// cfg is nil. This is safe to call even if cfg is in active use by a TLS
// client or server.
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return cfg.Clone()
}

type connLRU struct {
	ll *list.List // list.Element.Value type is of *persistConn
	m  map[*persistConn]*list.Element
}

// add adds pc to the head of the linked list.
func (cl *connLRU) add(pc *persistConn) {
	if cl.ll == nil {
		cl.ll = list.New()
		cl.m = make(map[*persistConn]*list.Element)
	}
	ele := cl.ll.PushFront(pc)
	if _, ok := cl.m[pc]; ok {
		panic("persistConn was already in LRU")
	}
	cl.m[pc] = ele
}

func (cl *connLRU) removeOldest() *persistConn {
	ele := cl.ll.Back()
	pc := ele.Value.(*persistConn)
	cl.ll.Remove(ele)
	delete(cl.m, pc)
	return pc
}

// remove removes pc from cl.
func (cl *connLRU) remove(pc *persistConn) {
	if ele, ok := cl.m[pc]; ok {
		cl.ll.Remove(ele)
		delete(cl.m, pc)
	}
}

// len returns the number of items in the cache.
func (cl *connLRU) len() int {
	return len(cl.m)
}

// validPort reports whether p (without the colon) is a valid port in
// a URL, per RFC 3986 Section 3.2.3, which says the port may be
// empty, or only contain digits.
func validPort(p string) bool {
	for _, r := range []byte(p) {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
`

func Benchmark(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		it, err := g.Go.Tokenise(nil, lexerBenchSource)
		assert.NoError(b, err)
		for t := it(); t != chroma.EOF; t = it() {
		}
	}
}
