package p

import (
    . "github.com/alecthomas/chroma" // nolint
    "github.com/alecthomas/chroma/lexers/internal"
)

// nolint

// Lexer for the Plutus Core Languages (version 2.1)
//
// including both Typed- and Untyped- versions
// based on “Formal Specification of the Plutus Core Language (version 2.1)”, published 6th April 2021:
// https://hydra.iohk.io/build/8205579/download/1/plutus-core-specification.pdf

var PlutusCoreLang = internal.Register(MustNewLazyLexer(
    &Config{
        Name:      "Plutus Core",
        Aliases:   []string{"plutus-core", "plc"},
        Filenames: []string{"*.plc"},
        MimeTypes: []string{"text/x-plutus-core", "application/x-plutus-core"},
    },
    plutusCoreRules,
))

func plutusCoreRules() Rules {
    return Rules{
        "root": {
            {`\s+`, Text, nil},
            {`(\(|\))`, Punctuation, nil},
            {`(\[|\])`, Punctuation, nil},
            {`({|})`, Punctuation, nil},

            // Constants. Figure 1.
            // For version, see handling of (program ...) below.
            {`([+-]?\d+)`, LiteralNumberInteger, nil},
            {`(#([a-fA-F0-9][a-fA-F0-9])+)`, LiteralString, nil},
            {`(\(\))`, NameConstant, nil},
            {`(True|False)`, NameConstant, nil},

            // Keywords. Figures 2 and 15.
            // Special handling for program because it is followed by a version.
            {`(con |abs |iwrap |unwrap |lam |builtin |delay |force |error)`, Keyword, nil},
            {`(fun |all |ifix |lam |con )`, Keyword, nil},
            {`(type|fun )`, Keyword, nil},
            {`(program )(\S+)`, ByGroups(Keyword, LiteralString), nil},

            // Built-in Types. Figure 12.
            {`(unit|boolean|integer|bytestring|str)`, KeywordType, nil},

            // Built-ins Functions. Figure 14.
            {`(ifThenElse)`, NameBuiltin, nil},
            {`(addInteger|subtractInteger|multiplyInteger|divideInteger|modInteger|quotientInteger|remainderInteger)`, NameBuiltin, nil},
            {`(lessThanInteger|lessThanEqualsInteger|greaterThanInteger|greaterThanEqualsInteger|equalsInteger)`, NameBuiltin, nil},
            {`(concatenate|equalsByteString|lessThanByteString|greaterThanByteString|takeByteString|dropByteString|sha2_256|sha3_256)`, NameBuiltin, nil},
            {`(verifySignature)`, NameBuiltin, nil},

            // Name. Figure 1.
            {`([a-zA-Z][a-zA-Z0-9_']*)`, Name, nil},
        },
    }
}
