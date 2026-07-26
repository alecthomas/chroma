package lexers

import (
	"iter"
	"strings"

	. "github.com/alecthomas/chroma/v3" // nolint
)

// Plan9RC lexer.
var Plan9RC = Register(&plan9RCLexer{
	config: &Config{
		Name:      "Plan 9 rc",
		Aliases:   []string{"plan9rc", "rc"},
		Filenames: []string{"*.rc"},
		MimeTypes: []string{"text/x-plan9-rc"},
		EnsureNL:  true,
	},
})

type plan9RCLexer struct {
	config   *Config
	analyser func(text string) float32
}

func (l *plan9RCLexer) Config() *Config { return l.config }

func (l *plan9RCLexer) Tokenise(options *TokeniseOptions, text string) (iter.Seq[Token], error) {
	if options == nil {
		options = &TokeniseOptions{EnsureLF: true}
	}
	if options.EnsureLF {
		text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n")
	}
	if l.config.EnsureNL && text != "" && !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	scanner := &plan9RCScanner{text: text, commandStart: true}
	scanner.scan(false)
	return Literator(scanner.tokens...), nil
}

func (l *plan9RCLexer) SetRegistry(registry *LexerRegistry) Lexer {
	_ = registry
	return l
}

func (l *plan9RCLexer) SetAnalyser(analyser func(text string) float32) Lexer {
	l.analyser = analyser
	return l
}

func (l *plan9RCLexer) AnalyseText(text string) float32 {
	if l.analyser != nil {
		return l.analyser(text)
	}
	if strings.HasPrefix(text, "#!/bin/rc\n") || strings.HasPrefix(text, "#!/usr/bin/env rc\n") {
		return 1.0
	}
	if strings.Contains(text, "`{") && strings.Contains(text, "fn ") {
		return 0.4
	}
	return 0
}

type plan9RCHereDoc struct {
	marker string
	quoted bool
}

type plan9RCScanner struct {
	text         string
	pos          int
	tokens       []Token
	hereDocs     []plan9RCHereDoc
	commandStart bool
	commandParen bool
}

func (s *plan9RCScanner) scan(stopAtBrace bool) {
	for s.pos < len(s.text) {
		if stopAtBrace && s.text[s.pos] == '}' {
			s.emit(LiteralStringBacktick, s.text[s.pos:s.pos+1])
			s.pos++
			return
		}

		switch {
		case s.pos == 0 && strings.HasPrefix(s.text, "#!"):
			s.emitLine(CommentHashbang)
		case s.text[s.pos] == '#':
			s.emitLine(CommentSingle)
		case strings.HasPrefix(s.text[s.pos:], "\\\n"):
			s.emit(LiteralStringEscape, s.text[s.pos:s.pos+2])
			s.pos += 2
		case isPlan9RCWhitespace(s.text[s.pos]):
			s.scanWhitespace()
		case s.text[s.pos] == '\'':
			s.scanSingleQuoted()
		case strings.HasPrefix(s.text[s.pos:], "`{"):
			s.emit(LiteralStringBacktick, "`{")
			s.pos += 2
			s.scanSubstitution()
		case strings.HasPrefix(s.text[s.pos:], "<{") || strings.HasPrefix(s.text[s.pos:], ">{"):
			s.emit(Operator, s.text[s.pos:s.pos+2])
			s.pos += 2
			s.scanSubstitution()
		case s.text[s.pos] == '$':
			s.scanVariable()
		case s.scanRedirection():
		case s.scanPipe():
		case s.scanOperator():
		default:
			s.scanWord()
		}
	}
}

func (s *plan9RCScanner) scanWhitespace() {
	start := s.pos
	for s.pos < len(s.text) && isPlan9RCWhitespace(s.text[s.pos]) {
		if s.text[s.pos] == '\n' {
			s.pos++
			s.emit(Text, s.text[start:s.pos])
			s.scanPendingHereDocs()
			s.commandStart = true
			return
		}
		s.pos++
	}
	s.emit(Text, s.text[start:s.pos])
}

func (s *plan9RCScanner) scanSingleQuoted() {
	start := s.pos
	s.pos++
	for s.pos < len(s.text) {
		if s.text[s.pos] != '\'' {
			s.pos++
			continue
		}
		s.pos++
		if s.pos < len(s.text) && s.text[s.pos] == '\'' {
			s.pos++
			continue
		}
		break
	}
	s.emit(LiteralStringSingle, s.text[start:s.pos])
	s.commandStart = false
}

func (s *plan9RCScanner) scanVariable() {
	start := s.pos
	s.pos++
	if s.pos < len(s.text) && (s.text[s.pos] == '#' || s.text[s.pos] == '"') {
		s.pos++
	}
	s.emit(Operator, s.text[start:s.pos])

	if s.pos >= len(s.text) {
		return
	}
	switch {
	case s.text[s.pos] == '$':
		return
	case s.text[s.pos] == '\'':
		s.scanSingleQuoted()
	case strings.HasPrefix(s.text[s.pos:], "`{"):
		s.emit(LiteralStringBacktick, "`{")
		s.pos += 2
		s.scanSubstitution()
	case strings.HasPrefix(s.text[s.pos:], "<{") || strings.HasPrefix(s.text[s.pos:], ">{"):
		s.emit(Operator, s.text[s.pos:s.pos+2])
		s.pos += 2
		s.scanSubstitution()
	case s.text[s.pos] == '(':
		s.scanOperator()
	case isPlan9RCVarChar(s.text[s.pos]):
		start = s.pos
		for s.pos < len(s.text) && isPlan9RCVarChar(s.text[s.pos]) {
			s.pos++
		}
		s.emit(NameVariable, s.text[start:s.pos])
		s.commandStart = false
	}
}

func (s *plan9RCScanner) scanSubstitution() {
	braceDepth := 0
	for s.pos < len(s.text) {
		if s.text[s.pos] == '}' {
			if braceDepth == 0 {
				s.emit(LiteralStringBacktick, s.text[s.pos:s.pos+1])
				s.pos++
				return
			}
			s.emit(Operator, s.text[s.pos:s.pos+1])
			s.pos++
			braceDepth--
			continue
		}
		if s.text[s.pos] == '{' {
			s.emit(Operator, s.text[s.pos:s.pos+1])
			s.pos++
			braceDepth++
			s.commandStart = true
			continue
		}

		switch {
		case s.text[s.pos] == '#':
			s.emitLine(CommentSingle)
		case strings.HasPrefix(s.text[s.pos:], "\\\n"):
			s.emit(LiteralStringEscape, s.text[s.pos:s.pos+2])
			s.pos += 2
		case isPlan9RCWhitespace(s.text[s.pos]):
			s.scanWhitespace()
		case s.text[s.pos] == '\'':
			s.scanSingleQuoted()
		case strings.HasPrefix(s.text[s.pos:], "`{"):
			s.emit(LiteralStringBacktick, "`{")
			s.pos += 2
			s.scanSubstitution()
		case strings.HasPrefix(s.text[s.pos:], "<{") || strings.HasPrefix(s.text[s.pos:], ">{"):
			s.emit(Operator, s.text[s.pos:s.pos+2])
			s.pos += 2
			s.scanSubstitution()
		case s.text[s.pos] == '$':
			s.scanVariable()
		case s.scanRedirection():
		case s.scanPipe():
		case s.scanOperator():
		default:
			s.scanWord()
		}
	}
}

func (s *plan9RCScanner) scanRedirection() bool {
	if s.pos >= len(s.text) || (s.text[s.pos] != '<' && s.text[s.pos] != '>') {
		return false
	}
	start := s.pos
	switch {
	case strings.HasPrefix(s.text[s.pos:], ">>"),
		strings.HasPrefix(s.text[s.pos:], "<<"),
		strings.HasPrefix(s.text[s.pos:], "<>"):
		s.pos += 2
	default:
		s.pos++
	}
	op := s.text[start:s.pos]
	s.emit(Operator, op)
	if s.pos < len(s.text) && s.text[s.pos] == '[' {
		s.scanDecorator()
	}
	if op == "<<" {
		s.scanHereDocMarker()
	}
	return true
}

func (s *plan9RCScanner) scanPipe() bool {
	if s.pos >= len(s.text) || s.text[s.pos] != '|' || strings.HasPrefix(s.text[s.pos:], "||") {
		return false
	}
	s.pos++
	s.emit(Operator, "|")
	if s.pos < len(s.text) && s.text[s.pos] == '[' {
		s.scanDecorator()
	}
	s.commandStart = true
	return true
}

func (s *plan9RCScanner) scanDecorator() {
	start := s.pos
	s.pos++
	for s.pos < len(s.text) && s.text[s.pos] != ']' && s.text[s.pos] != '\n' {
		s.pos++
	}
	if s.pos < len(s.text) && s.text[s.pos] == ']' {
		s.pos++
	}
	s.emit(NameDecorator, s.text[start:s.pos])
}

func (s *plan9RCScanner) scanHereDocMarker() {
	start := s.pos
	for s.pos < len(s.text) && (s.text[s.pos] == ' ' || s.text[s.pos] == '\t') {
		s.pos++
	}
	if s.pos > start {
		s.emit(Text, s.text[start:s.pos])
	}
	if s.pos >= len(s.text) || s.text[s.pos] == '\n' {
		return
	}

	quoted := false
	marker := ""
	if s.text[s.pos] == '\'' {
		quoted = true
		start = s.pos
		s.pos++
		var b strings.Builder
		for s.pos < len(s.text) {
			if s.text[s.pos] != '\'' {
				b.WriteByte(s.text[s.pos])
				s.pos++
				continue
			}
			s.pos++
			if s.pos < len(s.text) && s.text[s.pos] == '\'' {
				b.WriteByte('\'')
				s.pos++
				continue
			}
			break
		}
		s.emit(LiteralStringSingle, s.text[start:s.pos])
		marker = b.String()
	} else {
		start = s.pos
		for s.pos < len(s.text) && !isPlan9RCWhitespace(s.text[s.pos]) && !isPlan9RCSpecial(s.text[s.pos]) {
			s.pos++
		}
		marker = s.text[start:s.pos]
		s.emit(Text, marker)
	}
	if marker != "" {
		s.hereDocs = append(s.hereDocs, plan9RCHereDoc{marker: marker, quoted: quoted})
	}
}

func (s *plan9RCScanner) scanPendingHereDocs() {
	for len(s.hereDocs) > 0 && s.pos < len(s.text) {
		hereDoc := s.hereDocs[0]
		s.hereDocs = s.hereDocs[1:]
		s.scanHereDocBody(hereDoc)
	}
}

func (s *plan9RCScanner) scanHereDocBody(hereDoc plan9RCHereDoc) {
	bodyStart := s.pos
	for s.pos < len(s.text) {
		lineStart := s.pos
		lineEnd := strings.IndexByte(s.text[lineStart:], '\n')
		if lineEnd < 0 {
			s.pos = len(s.text)
			s.emitHereDocBody(hereDoc, s.text[bodyStart:s.pos])
			return
		}
		lineEnd += lineStart
		line := strings.TrimSuffix(s.text[lineStart:lineEnd], "\r")
		if line == hereDoc.marker {
			s.emitHereDocBody(hereDoc, s.text[bodyStart:lineStart])
			s.emit(Text, s.text[lineStart:lineEnd+1])
			s.pos = lineEnd + 1
			return
		}
		s.pos = lineEnd + 1
	}
	s.emitHereDocBody(hereDoc, s.text[bodyStart:s.pos])
}

func (s *plan9RCScanner) emitHereDocBody(hereDoc plan9RCHereDoc, body string) {
	if body == "" {
		return
	}
	if hereDoc.quoted {
		s.emit(LiteralStringHeredoc, body)
		return
	}
	start := 0
	for start < len(body) {
		i := strings.IndexByte(body[start:], '$')
		if i < 0 {
			s.emit(LiteralStringHeredoc, body[start:])
			return
		}
		i += start
		if i > start {
			s.emit(LiteralStringHeredoc, body[start:i])
		}
		j := i + 1
		if j < len(body) && (body[j] == '#' || body[j] == '"') {
			j++
		}
		s.emit(Operator, body[i:j])
		k := j
		for k < len(body) && isPlan9RCVarChar(body[k]) {
			k++
		}
		if k > j {
			s.emit(NameVariable, body[j:k])
			if k < len(body) && body[k] == '^' {
				s.emit(Operator, body[k:k+1])
				k++
			}
		}
		start = k
	}
}

func (s *plan9RCScanner) scanOperator() bool {
	if strings.HasPrefix(s.text[s.pos:], "&&") || strings.HasPrefix(s.text[s.pos:], "||") {
		s.emit(Operator, s.text[s.pos:s.pos+2])
		s.pos += 2
		s.commandStart = true
		return true
	}
	if !strings.ContainsRune(";&^=(){}!@~", rune(s.text[s.pos])) {
		return false
	}
	ch := s.text[s.pos]
	tokenType := Operator
	if ch == '!' || ch == '@' || ch == '~' {
		tokenType = Keyword
	}
	s.emit(tokenType, s.text[s.pos:s.pos+1])
	s.pos++
	switch ch {
	case ';', '&', '{', '}':
		s.commandStart = true
	case '(':
		s.commandStart = s.commandParen
		s.commandParen = false
	case ')':
		s.commandStart = true
	default:
		s.commandStart = false
	}
	return true
}

func (s *plan9RCScanner) scanWord() {
	start := s.pos
	for s.pos < len(s.text) && !isPlan9RCWhitespace(s.text[s.pos]) && !isPlan9RCSpecial(s.text[s.pos]) && !strings.HasPrefix(s.text[s.pos:], "\\\n") {
		s.pos++
	}
	word := s.text[start:s.pos]
	switch {
	case s.commandStart && word == ".":
		s.emit(NameBuiltin, word)
	case isPlan9RCKeyword(word):
		s.emit(Keyword, word)
		switch word {
		case "if", "while":
			s.commandParen = true
			s.commandStart = false
		case "not":
			s.commandStart = true
		default:
			s.commandStart = false
		}
		return
	case isPlan9RCBuiltin(word):
		s.emit(NameBuiltin, word)
	case s.pos < len(s.text) && s.text[s.pos] == '=' && isPlan9RCName(word):
		s.emit(NameVariable, word)
	default:
		s.emit(Text, word)
	}
	s.commandStart = false
}

func (s *plan9RCScanner) emitLine(tokenType TokenType) {
	start := s.pos
	for s.pos < len(s.text) && s.text[s.pos] != '\n' {
		s.pos++
	}
	if s.pos < len(s.text) {
		s.pos++
	}
	s.emit(tokenType, s.text[start:s.pos])
	if tokenType == CommentSingle {
		s.scanPendingHereDocs()
	}
	s.commandStart = true
}

func (s *plan9RCScanner) emit(tokenType TokenType, value string) {
	if value == "" {
		return
	}
	s.tokens = append(s.tokens, Token{Type: tokenType, Value: value})
}

func isPlan9RCWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isPlan9RCSpecial(ch byte) bool {
	return strings.ContainsRune("#;&|^$=`'{}()<>", rune(ch))
}

func isPlan9RCVarChar(ch byte) bool {
	return ch == '*' || ch == '_' || ('0' <= ch && ch <= '9') || ('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z')
}

func isPlan9RCName(s string) bool {
	if s == "" {
		return false
	}
	for i := range len(s) {
		if !isPlan9RCVarChar(s[i]) {
			return false
		}
	}
	return true
}

func isPlan9RCKeyword(s string) bool {
	switch s {
	case "for", "in", "while", "if", "not", "switch", "fn", "case":
		return true
	}
	return false
}

func isPlan9RCBuiltin(s string) bool {
	switch s {
	case "builtin", "cd", "eval", "exec", "flag", "exit", "rfork", "shift", "wait", "whatis":
		return true
	}
	return false
}
