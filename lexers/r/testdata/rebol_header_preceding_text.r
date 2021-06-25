preface.... everything what is before header is not evaluated
so this should not be colorized:
1 + 2

REBOL [] ;<- this is minimal header, everything behind it must be colorized

;## String tests ##
print "Hello ^"World" ;<- with escaped char
multiline-string: {
    bla bla "bla" {bla}
}
char-a: #"a"
escaped-a: #"^(61)"
new-line: #"^/"

;## Binaries ##
print decompress 64#{eJzLSM3JyQcABiwCFQUAAAA=}
;2#{0000 00000} ;<- this one is invalid!
2#{}
#{FF00}

;##Date + time ##
1-Feb-2009
1-Feb-2009/2:24:46+1:0
1:0 1:1:1 -0:1.1
