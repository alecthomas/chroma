
package styles

import (
    "github.com/alecthomas/chroma"
)

// Default style.
var Default = Register(chroma.NewStyle("default", chroma.StyleEntries{
    chroma.TextWhitespace: "#bbbbbb",
    chroma.Comment: "italic #408080",
    chroma.CommentPreproc: "noitalic #BC7A00",
    chroma.Keyword: "bold #008000",
    chroma.KeywordPseudo: "nobold",
    chroma.KeywordType: "nobold #B00040",
    chroma.Operator: "#666666",
    chroma.OperatorWord: "bold #AA22FF",
    chroma.NameBuiltin: "#008000",
    chroma.NameFunction: "#0000FF",
    chroma.NameClass: "bold #0000FF",
    chroma.NameNamespace: "bold #0000FF",
    chroma.NameException: "bold #D2413A",
    chroma.NameVariable: "#19177C",
    chroma.NameConstant: "#880000",
    chroma.NameLabel: "#A0A000",
    chroma.NameEntity: "bold #999999",
    chroma.NameAttribute: "#7D9029",
    chroma.NameTag: "bold #008000",
    chroma.NameDecorator: "#AA22FF",
    chroma.LiteralString: "#BA2121",
    chroma.LiteralStringDoc: "italic",
    chroma.LiteralStringInterpol: "bold #BB6688",
    chroma.LiteralStringEscape: "bold #BB6622",
    chroma.LiteralStringRegex: "#BB6688",
    chroma.LiteralStringSymbol: "#19177C",
    chroma.LiteralStringOther: "#008000",
    chroma.LiteralNumber: "#666666",
    chroma.GenericHeading: "bold #000080",
    chroma.GenericSubheading: "bold #800080",
    chroma.GenericDeleted: "#A00000",
    chroma.GenericInserted: "#00A000",
    chroma.GenericError: "#FF0000",
    chroma.GenericEmph: "italic",
    chroma.GenericStrong: "bold",
    chroma.GenericPrompt: "bold #000080",
    chroma.GenericOutput: "#888",
    chroma.GenericTraceback: "#04D",
    chroma.Error: "border:#FF0000",
    chroma.Background: " bg:#f8f8f8",
}))

