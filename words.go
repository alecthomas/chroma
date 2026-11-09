package chroma

import (
	"bytes"
	"errors"
	"math/bits"
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
	size := 0
	for _, word := range words {
		for _, r := range word {
			size += utf8.RuneLen(r)
		}
	}
	// One buffer holds every encoded word; its capacity is exact so the slices into it stay valid.
	buf := make([]byte, 0, size)
	encoded := make([][]byte, len(words))
	for i, word := range words {
		start := len(buf)
		for _, r := range word {
			buf = utf8.AppendRune(buf, r)
		}
		encoded[i] = buf[start:len(buf):len(buf)]
	}
	slices.SortFunc(encoded, bytes.Compare)
	encoded = slices.CompactFunc(encoded, bytes.Equal)
	b := trieBuilder{free: 1, size: 1}
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
	b.ensure(1)
	b.used[0] = 1
	b.place(0, encoded, 0)
	t.base = exactCopy(t.base[:b.size])
	t.check = exactCopy(t.check[:b.size])
	t.terminal = exactCopy(t.terminal[:(b.size+63)/64])
	return *t
}

// exactCopy drops the spare capacity left by growing the arrays geometrically.
func exactCopy[T any](s []T) []T {
	return append(make([]T, 0, len(s)), s...)
}

type trieBuilder struct {
	trie wordTrie
	// used marks occupied slots, so base searches can jump to the next free one.
	used []uint64
	// size is one past the highest occupied slot.
	size int32
	// free is a lower bound on the lowest unused slot, so base searches skip the packed prefix.
	free int32
	// starts and codes are stacks of each pending node's child groups, shared to avoid
	// allocating per node; place only appends past its caller's entries.
	starts, codes []int32
}

// ensure grows the arrays geometrically to hold at least size slots.
func (b *trieBuilder) ensure(size int) {
	t := &b.trie
	if size <= len(t.check) {
		return
	}
	size = max(size, 2*len(t.check))
	extra := size - len(t.check)
	t.base = append(t.base, make([]int32, extra)...)
	t.check = append(t.check, make([]int32, extra)...)
	words := (size + 63) / 64
	t.terminal = append(t.terminal, make([]uint64, words-len(t.terminal))...)
	b.used = append(b.used, make([]uint64, words-len(b.used))...)
}

func (b *trieBuilder) isUsed(slot int32) bool {
	w := int(slot / 64)
	return w < len(b.used) && b.used[w]&(1<<(slot%64)) != 0
}

// nextFree returns the lowest unused slot at or after from.
func (b *trieBuilder) nextFree(from int32) int32 {
	for w := int(from / 64); w < len(b.used); w++ {
		unused := ^b.used[w]
		if w == int(from/64) {
			unused &^= 1<<(from%64) - 1
		}
		if unused != 0 {
			return int32(w*64 + bits.TrailingZeros64(unused)) //nolint:gosec
		}
	}
	return max(from, int32(len(b.used)*64)) //nolint:gosec
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
	mark := len(b.starts)
	for i := 0; i < len(words); {
		b.starts = append(b.starts, int32(i))
		b.codes = append(b.codes, int32(t.codes[words[i][depth]]))
		j := i + 1
		for j < len(words) && words[j][depth] == words[i][depth] {
			j++
		}
		i = j
	}
	n := len(b.starts) - mark
	base := b.freeBase(b.codes[mark:])
	t.base[state] = base
	for _, code := range b.codes[mark:] {
		slot := base + code
		t.check[slot] = state + 1
		b.used[slot/64] |= 1 << (slot % 64)
		b.size = max(b.size, slot+1)
	}
	b.free = b.nextFree(b.free)
	for k := range n {
		end := len(words)
		if k+1 < n {
			end = int(b.starts[mark+k+1])
		}
		b.place(base+b.codes[mark+k], words[b.starts[mark+k]:end], depth+1)
	}
	b.starts = b.starts[:mark]
	b.codes = b.codes[:mark]
}

// freeBase returns the lowest base at which every child slot is unused.
func (b *trieBuilder) freeBase(codes []int32) int32 {
	first, last := codes[0], codes[len(codes)-1]
	for slot := b.nextFree(max(b.free, first)); ; slot = b.nextFree(slot + 1) {
		base := slot - first
		b.ensure(int(base + last + 1))
		free := true
		for _, code := range codes[1:] {
			if b.isUsed(base + code) {
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
			if state, ok = m.trie.step(state, byte(r)); !ok { //nolint:gosec
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
	for _, end := range slices.Backward(ends) {
		ok, err := assertAt(m.suffix, text, end)
		if err != nil {
			return -1, err
		}
		if ok {
			return end, nil
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
