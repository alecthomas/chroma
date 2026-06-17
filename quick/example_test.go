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
	// /* Background */ .bg { color: #f8f8f2; background-color: #272822; }
	// /* PreWrapper */ .chroma { color: #f8f8f2; background-color: #272822; -webkit-text-size-adjust: none; }
	// /* Error */ .chroma .err { color: #960050; background-color: #1e0010 }
	// /* LineLink */ .chroma .lnlinks { outline: none; text-decoration: none; color: inherit }
	// /* LineTableTD */ .chroma .lntd { vertical-align: top; padding: 0; margin: 0; border: 0; }
	// /* LineTable */ .chroma .lntable { border-spacing: 0; padding: 0; margin: 0; border: 0; }
	// /* LineHighlight */ .chroma .hl { background-color: #3c3d38 }
	// /* LineNumbersTable */ .chroma .lnt { white-space: pre; -webkit-user-select: none; user-select: none; margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
	// /* LineNumbers */ .chroma .ln { white-space: pre; -webkit-user-select: none; user-select: none; margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
	// /* Line */ .chroma .line { display: flex; }
	// /* Keyword */ .chroma .k { color: #66d9ef }
	// /* KeywordConstant */ .chroma .kc { color: #66d9ef }
	// /* KeywordDeclaration */ .chroma .kd { color: #66d9ef }
	// /* KeywordNamespace */ .chroma .kn { color: #f92672 }
	// /* KeywordPseudo */ .chroma .kp { color: #66d9ef }
	// /* KeywordReserved */ .chroma .kr { color: #66d9ef }
	// /* KeywordType */ .chroma .kt { color: #66d9ef }
	// /* NameAttribute */ .chroma .na { color: #a6e22e }
	// /* NameClass */ .chroma .nc { color: #a6e22e }
	// /* NameConstant */ .chroma .no { color: #66d9ef }
	// /* NameDecorator */ .chroma .nd { color: #a6e22e }
	// /* NameException */ .chroma .ne { color: #a6e22e }
	// /* NameOther */ .chroma .nx { color: #a6e22e }
	// /* NameTag */ .chroma .nt { color: #f92672 }
	// /* NameFunction */ .chroma .nf { color: #a6e22e }
	// /* NameFunctionMagic */ .chroma .fm { color: #a6e22e }
	// /* Literal */ .chroma .l { color: #ae81ff }
	// /* LiteralDate */ .chroma .ld { color: #e6db74 }
	// /* LiteralString */ .chroma .s { color: #e6db74 }
	// /* LiteralStringAffix */ .chroma .sa { color: #e6db74 }
	// /* LiteralStringBacktick */ .chroma .sb { color: #e6db74 }
	// /* LiteralStringChar */ .chroma .sc { color: #e6db74 }
	// /* LiteralStringDelimiter */ .chroma .dl { color: #e6db74 }
	// /* LiteralStringDoc */ .chroma .sd { color: #e6db74 }
	// /* LiteralStringDouble */ .chroma .s2 { color: #e6db74 }
	// /* LiteralStringEscape */ .chroma .se { color: #ae81ff }
	// /* LiteralStringHeredoc */ .chroma .sh { color: #e6db74 }
	// /* LiteralStringInterpol */ .chroma .si { color: #e6db74 }
	// /* LiteralStringOther */ .chroma .sx { color: #e6db74 }
	// /* LiteralStringRegex */ .chroma .sr { color: #e6db74 }
	// /* LiteralStringSingle */ .chroma .s1 { color: #e6db74 }
	// /* LiteralStringSymbol */ .chroma .ss { color: #e6db74 }
	// /* LiteralNumber */ .chroma .m { color: #ae81ff }
	// /* LiteralNumberBin */ .chroma .mb { color: #ae81ff }
	// /* LiteralNumberFloat */ .chroma .mf { color: #ae81ff }
	// /* LiteralNumberHex */ .chroma .mh { color: #ae81ff }
	// /* LiteralNumberInteger */ .chroma .mi { color: #ae81ff }
	// /* LiteralNumberIntegerLong */ .chroma .il { color: #ae81ff }
	// /* LiteralNumberOct */ .chroma .mo { color: #ae81ff }
	// /* Operator */ .chroma .o { color: #f92672 }
	// /* OperatorWord */ .chroma .ow { color: #f92672 }
	// /* OperatorReserved */ .chroma .or { color: #f92672 }
	// /* Comment */ .chroma .c { color: #75715e }
	// /* CommentHashbang */ .chroma .ch { color: #75715e }
	// /* CommentMultiline */ .chroma .cm { color: #75715e }
	// /* CommentSingle */ .chroma .c1 { color: #75715e }
	// /* CommentSpecial */ .chroma .cs { color: #75715e }
	// /* CommentPreproc */ .chroma .cp { color: #75715e }
	// /* CommentPreprocFile */ .chroma .cpf { color: #75715e }
	// /* GenericDeleted */ .chroma .gd { color: #f92672 }
	// /* GenericEmph */ .chroma .ge { font-style: italic }
	// /* GenericInserted */ .chroma .gi { color: #a6e22e }
	// /* GenericStrong */ .chroma .gs { font-weight: bold }
	// /* GenericSubheading */ .chroma .gu { color: #75715e }
	// body { color:#f8f8f2;background-color:#272822;; }
	// </style><body class="bg">
	// <pre class="chroma"><code><span class="line"><span class="cl"><span class="kn">package</span><span class="w"> </span><span class="nx">main</span><span class="w">
	// </span></span></span><span class="line"><span class="cl"><span class="w">
	// </span></span></span><span class="line"><span class="cl"><span class="kd">func</span><span class="w"> </span><span class="nf">main</span><span class="p">()</span><span class="w"> </span><span class="p">{</span><span class="w"> </span><span class="p">}</span><span class="w">
	// </span></span></span></code></pre>
	// </body>
	// </html>
}
