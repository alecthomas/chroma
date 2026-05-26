package quick_test

import (
	"log"
	"os"

	"github.com/alecthomas/chroma/v2/quick"
)

func Example() {
	code := `package main

func main() { }
`
	err := quick.Highlight(os.Stdout, code, "go", "html", "monokai")
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// <html>
	// <style type="text/css">
	// /* Background */ .bg.dark { color: #f8f8f2; background-color: #272822; }
	// /* PreWrapper */ .chroma.dark { color: #f8f8f2; background-color: #272822; -webkit-text-size-adjust: none; }
	// /* Error */ .chroma.dark .err { color: #960050; background-color: #1e0010 }
	// /* LineLink */ .chroma.dark .lnlinks { outline: none; text-decoration: none; color: inherit }
	// /* LineTableTD */ .chroma.dark .lntd { vertical-align: top; padding: 0; margin: 0; border: 0; }
	// /* LineTable */ .chroma.dark .lntable { border-spacing: 0; padding: 0; margin: 0; border: 0; }
	// /* LineHighlight */ .chroma.dark .hl { background-color: #3c3d38 }
	// /* LineNumbersTable */ .chroma.dark .lnt { white-space: pre; -webkit-user-select: none; user-select: none; margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
	// /* LineNumbers */ .chroma.dark .ln { white-space: pre; -webkit-user-select: none; user-select: none; margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
	// /* Line */ .chroma.dark .line { display: flex; }
	// /* Keyword */ .chroma.dark .k { color: #66d9ef }
	// /* KeywordConstant */ .chroma.dark .kc { color: #66d9ef }
	// /* KeywordDeclaration */ .chroma.dark .kd { color: #66d9ef }
	// /* KeywordNamespace */ .chroma.dark .kn { color: #f92672 }
	// /* KeywordPseudo */ .chroma.dark .kp { color: #66d9ef }
	// /* KeywordReserved */ .chroma.dark .kr { color: #66d9ef }
	// /* KeywordType */ .chroma.dark .kt { color: #66d9ef }
	// /* NameAttribute */ .chroma.dark .na { color: #a6e22e }
	// /* NameClass */ .chroma.dark .nc { color: #a6e22e }
	// /* NameConstant */ .chroma.dark .no { color: #66d9ef }
	// /* NameDecorator */ .chroma.dark .nd { color: #a6e22e }
	// /* NameException */ .chroma.dark .ne { color: #a6e22e }
	// /* NameOther */ .chroma.dark .nx { color: #a6e22e }
	// /* NameTag */ .chroma.dark .nt { color: #f92672 }
	// /* NameFunction */ .chroma.dark .nf { color: #a6e22e }
	// /* NameFunctionMagic */ .chroma.dark .fm { color: #a6e22e }
	// /* Literal */ .chroma.dark .l { color: #ae81ff }
	// /* LiteralDate */ .chroma.dark .ld { color: #e6db74 }
	// /* LiteralString */ .chroma.dark .s { color: #e6db74 }
	// /* LiteralStringAffix */ .chroma.dark .sa { color: #e6db74 }
	// /* LiteralStringBacktick */ .chroma.dark .sb { color: #e6db74 }
	// /* LiteralStringChar */ .chroma.dark .sc { color: #e6db74 }
	// /* LiteralStringDelimiter */ .chroma.dark .dl { color: #e6db74 }
	// /* LiteralStringDoc */ .chroma.dark .sd { color: #e6db74 }
	// /* LiteralStringDouble */ .chroma.dark .s2 { color: #e6db74 }
	// /* LiteralStringEscape */ .chroma.dark .se { color: #ae81ff }
	// /* LiteralStringHeredoc */ .chroma.dark .sh { color: #e6db74 }
	// /* LiteralStringInterpol */ .chroma.dark .si { color: #e6db74 }
	// /* LiteralStringOther */ .chroma.dark .sx { color: #e6db74 }
	// /* LiteralStringRegex */ .chroma.dark .sr { color: #e6db74 }
	// /* LiteralStringSingle */ .chroma.dark .s1 { color: #e6db74 }
	// /* LiteralStringSymbol */ .chroma.dark .ss { color: #e6db74 }
	// /* LiteralNumber */ .chroma.dark .m { color: #ae81ff }
	// /* LiteralNumberBin */ .chroma.dark .mb { color: #ae81ff }
	// /* LiteralNumberFloat */ .chroma.dark .mf { color: #ae81ff }
	// /* LiteralNumberHex */ .chroma.dark .mh { color: #ae81ff }
	// /* LiteralNumberInteger */ .chroma.dark .mi { color: #ae81ff }
	// /* LiteralNumberIntegerLong */ .chroma.dark .il { color: #ae81ff }
	// /* LiteralNumberOct */ .chroma.dark .mo { color: #ae81ff }
	// /* Operator */ .chroma.dark .o { color: #f92672 }
	// /* OperatorWord */ .chroma.dark .ow { color: #f92672 }
	// /* OperatorReserved */ .chroma.dark .or { color: #f92672 }
	// /* Comment */ .chroma.dark .c { color: #75715e }
	// /* CommentHashbang */ .chroma.dark .ch { color: #75715e }
	// /* CommentMultiline */ .chroma.dark .cm { color: #75715e }
	// /* CommentSingle */ .chroma.dark .c1 { color: #75715e }
	// /* CommentSpecial */ .chroma.dark .cs { color: #75715e }
	// /* CommentPreproc */ .chroma.dark .cp { color: #75715e }
	// /* CommentPreprocFile */ .chroma.dark .cpf { color: #75715e }
	// /* GenericDeleted */ .chroma.dark .gd { color: #f92672 }
	// /* GenericEmph */ .chroma.dark .ge { font-style: italic }
	// /* GenericInserted */ .chroma.dark .gi { color: #a6e22e }
	// /* GenericStrong */ .chroma.dark .gs { font-weight: bold }
	// /* GenericSubheading */ .chroma.dark .gu { color: #75715e }
	// body { color:#f8f8f2;background-color:#272822;; }
	// </style><body class="bg dark">
	// <pre class="chroma dark"><code><span class="line"><span class="cl"><span class="kn">package</span><span class="w"> </span><span class="nx">main</span><span class="w">
	// </span></span></span><span class="line"><span class="cl"><span class="w">
	// </span></span></span><span class="line"><span class="cl"><span class="kd">func</span><span class="w"> </span><span class="nf">main</span><span class="p">()</span><span class="w"> </span><span class="p">{</span><span class="w"> </span><span class="p">}</span><span class="w">
	// </span></span></span></code></pre>
	// </body>
	// </html>
}
