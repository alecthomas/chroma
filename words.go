package chroma

import (
	"bytes"
	"errors"
	"slices"
	"unicode"
	"unicode/utf8"

	"github.com/dlclark/regexp2/v2"
	"github.com/dlclark/regexp2/v2/syntax"
)

// wordList is a rule's words attribute, matched with a trie rather than a regex.
type wordList struct {
	prefix, suffix string
	words          []string
}

// wordTrie is a double-array trie over the UTF-8 bytes of its words.
//
// State 0 is the root. A transition from state s on byte b goes to
// t = base[s] + codes[b], valid only if check[t] == s+1; a code of 0 means b
// occurs in no word.
type wordTrie struct {
	codes    [256]uint8
	base     []int32
	check    []int32
	terminal []uint64
}

func newWordTrie(words [][]rune) wordTrie {
	encoded := make([][]byte, len(words))
	for i, word := range words {
		encoded[i] = []byte(string(word))
	}
	slices.SortFunc(encoded, bytes.Compare)
	encoded = slices.CompactFunc(encoded, bytes.Equal)
	b := trieBuilder{free: 1}
	t := &b.trie
	var seen [256]bool
	for _, word := range encoded {
		for _, c := range word {
			seen[c] = true
		}
	}
	var code uint8
	for c, ok := range seen {
		if ok {
			code++
			t.codes[c] = code
		}
	}
	t.grow(1)
	t.check[0] = -1
	b.place(0, encoded, 0)
	t.base = slices.Clip(t.base)
	t.check = slices.Clip(t.check)
	t.terminal = slices.Clip(t.terminal)
	return *t
}

type trieBuilder struct {
	trie wordTrie
	// free is a lower bound on the lowest unused slot, so base searches skip the packed prefix.
	free int32
}

func (t *wordTrie) grow(size int) {
	if extra := size - len(t.check); extra > 0 {
		t.base = append(t.base, make([]int32, extra)...)
		t.check = append(t.check, make([]int32, extra)...)
	}
	if words := (size + 63) / 64; len(t.terminal) < words {
		t.terminal = append(t.terminal, make([]uint64, words-len(t.terminal))...)
	}
}

// place lays out the children of state, whose words share their first depth bytes.
func (b *trieBuilder) place(state int32, words [][]byte, depth int) {
	t := &b.trie
	if len(words[0]) == depth {
		t.terminal[state/64] |= 1 << (state % 64)
		words = words[1:]
	}
	if len(words) == 0 {
		return
	}
	var groups [][][]byte
	var codes []int32
	for i := 0; i < len(words); {
		j := i + 1
		for j < len(words) && words[j][depth] == words[i][depth] {
			j++
		}
		groups = append(groups, words[i:j])
		codes = append(codes, int32(t.codes[words[i][depth]]))
		i = j
	}
	base := b.freeBase(codes)
	t.base[state] = base
	for _, code := range codes {
		t.check[base+code] = state + 1
	}
	for int(b.free) < len(t.check) && t.check[b.free] != 0 {
		b.free++
	}
	for i, group := range groups {
		b.place(base+codes[i], group, depth+1)
	}
}

// freeBase returns the lowest base at which every child slot is unused.
func (b *trieBuilder) freeBase(codes []int32) int32 {
	t := &b.trie
	for base := max(0, b.free-codes[0]); ; base++ {
		t.grow(int(base + codes[len(codes)-1] + 1))
		free := true
		for _, code := range codes {
			if t.check[base+code] != 0 {
				free = false
				break
			}
		}
		if free {
			return base
		}
	}
}

func (t *wordTrie) step(state int32, b byte) (int32, bool) {
	code := t.codes[b]
	if code == 0 {
		return 0, false
	}
	next := t.base[state] + int32(code)
	if int(next) >= len(t.check) || t.check[next] != state+1 {
		return 0, false
	}
	return next, true
}

func (t *wordTrie) isTerminal(state int32) bool {
	return t.terminal[state/64]&(1<<(state%64)) != 0
}

type wordMatcher struct {
	trie           wordTrie
	foldCase       bool
	prefix, suffix *regexp2.Regexp
}

func newWordMatcher(list *wordList, flags string, foldCase bool) (*wordMatcher, error) {
	m := &wordMatcher{foldCase: foldCase}
	words := make([][]rune, len(list.words))
	for i, word := range list.words {
		words[i] = []rune(word)
		if foldCase {
			for j, r := range words[i] {
				words[i][j] = unicode.ToLower(r)
			}
		}
	}
	m.trie = newWordTrie(words)
	var err error
	if m.prefix, err = compileWordAssertion(list.prefix, flags); err != nil {
		return nil, err
	}
	if m.suffix, err = compileWordAssertion(list.suffix, flags); err != nil {
		return nil, err
	}
	return m, nil
}

// compileWordAssertion compiles a words prefix or suffix, which must match without consuming
// text or capturing, so that a trie match produces the same groups as the equivalent regex.
func compileWordAssertion(pattern, flags string) (*regexp2.Regexp, error) {
	if pattern == "" {
		return nil, nil
	}
	tree, err := syntax.Parse(pattern, syntax.ParseOptions{})
	if err != nil {
		return nil, err
	}
	if tree.Captop > 1 {
		return nil, errors.New("words prefix and suffix must not capture: " + pattern)
	}
	if !zeroWidth(tree.Root) {
		return nil, errors.New("words prefix and suffix must be zero-width: " + pattern)
	}
	if flags != "" {
		pattern = "(?" + flags + ")" + pattern
	}
	return regexp2.Compile(`\G(?:`+pattern+`)`, regexp2.None)
}

func zeroWidth(node *syntax.RegexNode) bool {
	switch node.T {
	case syntax.NtBol, syntax.NtEol, syntax.NtBoundary, syntax.NtNonboundary, syntax.NtBeginning,
		syntax.NtStart, syntax.NtEndZ, syntax.NtEnd, syntax.NtEmpty, syntax.NtECMABoundary,
		syntax.NtNonECMABoundary, syntax.NtPosLook, syntax.NtNegLook:
		return true
	case syntax.NtConcatenate, syntax.NtAlternate, syntax.NtGroup, syntax.NtAtomic, syntax.NtCapture,
		syntax.NtLoop, syntax.NtLazyloop:
		for _, child := range node.Children {
			if !zeroWidth(child) {
				return false
			}
		}
		return true
	}
	return false
}

// match returns the end of the longest word at pos whose prefix and suffix also match, or -1.
func (m *wordMatcher) match(text []rune, pos int) (int, error) {
	var buf [16]int
	ends := buf[:0]
	state := int32(0)
walk:
	for i := pos; i < len(text); i++ {
		r := text[i]
		if m.foldCase {
			r = unicode.ToLower(r)
		}
		var ok bool
		if r < utf8.RuneSelf {
			if state, ok = m.trie.step(state, byte(r)); !ok {
				break
			}
		} else {
			var encoded [utf8.UTFMax]byte
			for _, b := range encoded[:utf8.EncodeRune(encoded[:], r)] {
				if state, ok = m.trie.step(state, b); !ok {
					break walk
				}
			}
		}
		if m.trie.isTerminal(state) {
			ends = append(ends, i+1)
		}
	}
	if len(ends) == 0 {
		return -1, nil
	}
	if ok, err := assertAt(m.prefix, text, pos); !ok || err != nil {
		return -1, err
	}
	for i := len(ends) - 1; i >= 0; i-- {
		ok, err := assertAt(m.suffix, text, ends[i])
		if err != nil {
			return -1, err
		}
		if ok {
			return ends[i], nil
		}
	}
	return -1, nil
}

func assertAt(re *regexp2.Regexp, text []rune, pos int) (bool, error) {
	if re == nil {
		return true, nil
	}
	match, err := re.FindRunesMatchStartingAt(text, pos)
	return match != nil && match.RuneIndex == pos, err
}
