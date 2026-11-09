package chroma

import (
	"errors"
	"slices"
	"unicode"

	"github.com/dlclark/regexp2/v2"
	"github.com/dlclark/regexp2/v2/syntax"
)

// wordList is a rule's words attribute, matched with a trie rather than a regex.
type wordList struct {
	prefix, suffix string
	words          []string
}

// trieNode's edges are labels[first:first+count] and targets[first:first+count], sorted by label.
type trieNode struct {
	first, count uint32
	terminal     bool
}

type wordTrie struct {
	nodes   []trieNode
	labels  []rune
	targets []uint32
}

func newWordTrie(words [][]rune) wordTrie {
	words = slices.Clone(words)
	slices.SortFunc(words, slices.Compare)
	words = slices.CompactFunc(words, slices.Equal)
	t := wordTrie{}
	t.build(words, 0)
	return t
}

// build adds the node for words, which share their first depth runes, and returns its index.
func (t *wordTrie) build(words [][]rune, depth int) uint32 {
	index := uint32(len(t.nodes))
	t.nodes = append(t.nodes, trieNode{})
	if len(words[0]) == depth {
		t.nodes[index].terminal = true
		words = words[1:]
	}
	first := uint32(len(t.labels))
	var groups [][][]rune
	for i := 0; i < len(words); {
		j := i + 1
		for j < len(words) && words[j][depth] == words[i][depth] {
			j++
		}
		t.labels = append(t.labels, words[i][depth])
		t.targets = append(t.targets, 0)
		groups = append(groups, words[i:j])
		i = j
	}
	t.nodes[index].first = first
	t.nodes[index].count = uint32(len(groups))
	for i, group := range groups {
		t.targets[first+uint32(i)] = t.build(group, depth+1)
	}
	return index
}

func (t *wordTrie) child(node uint32, r rune) (uint32, bool) {
	n := t.nodes[node]
	labels := t.labels[n.first : n.first+n.count]
	if len(labels) <= 8 {
		for i, label := range labels {
			if label == r {
				return t.targets[n.first+uint32(i)], true
			}
		}
		return 0, false
	}
	i, found := slices.BinarySearch(labels, r)
	if !found {
		return 0, false
	}
	return t.targets[n.first+uint32(i)], true
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
	node := uint32(0)
	for i := pos; i < len(text); i++ {
		r := text[i]
		if m.foldCase {
			r = unicode.ToLower(r)
		}
		next, ok := m.trie.child(node, r)
		if !ok {
			break
		}
		node = next
		if m.trie.nodes[node].terminal {
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
