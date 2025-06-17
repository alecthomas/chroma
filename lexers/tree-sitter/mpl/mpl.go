package mpl

// #cgo CFLAGS: -I${SRCDIR}/../../../../vendor/tree-sitter/include
// #cgo LDFLAGS: -L${SRCDIR}/../../../../vendor/tree-sitter/lib -ltree-sitter
// #include "tree_sitter/parser.h"
import "C"
import (
	"unsafe"

	sitter "github.com/smacker/go-tree-sitter"
)

// GetLanguage returns the MPL language.
func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_mpl())
	return sitter.NewLanguage(ptr)
}
