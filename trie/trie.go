package trie

import (
	"fmt"

	"github.com/disiqueira/gotree"
)

type wordArray struct {
	words []string
}

func (w *wordArray) add(word string) {
	w.words = append(w.words, word)
}

// Trie defines a trie
type Trie struct {
	Root *node
}

// NewTrie creates a trie
func NewTrie() *Trie {
	return &Trie{
		Root: &node{
			children: make(Children),
			isRoot:   true,
		},
	}
}

// HasWord finds the terminating node for the word specified. If no node is found and error is returned
func (t *Trie) HasWord(word string) error {
	if len(word) == 0 {
		return fmt.Errorf("no string to find")
	}

	runes := []rune(word)

	_, err := t.findAtNode(t.Root, runes, 0)
	if err != nil {
		return fmt.Errorf("word %s not found: %s", word, err)
	}

	return nil
}

// findAtNode gets the node beginning from specified node where the runes terminate
func (t *Trie) findAtNode(n *node, runes []rune, pos int) (*node, error) {

	if pos >= len(runes) {
		return nil, nil
	}

	r := runes[pos]
	cNode, ok := n.Children()[r]
	if !ok {
		var prefix string
		if pos > 0 {
			prefix = string(runes[0 : pos-1])
		}
		return nil, fmt.Errorf("string %s not found, longest prefix found: %s", string(runes), prefix)
	}

	if pos < (len(runes) - 1) {
		pos = pos + 1
	} else {
		// This was the last character, check if the node is terminating
		if !cNode.IsTerm() {
			return nil, fmt.Errorf("string %s not found but exists as a non-terminated path", string(runes))
		}
		return cNode, nil
	}

	return t.findAtNode(cNode, runes, pos)
}

// Add adds a word to the trie. If the word already exists in the trie and error is returned
func (t *Trie) Add(word string) error {
	if len(word) == 0 {
		return fmt.Errorf("no string to add")
	}

	runes := []rune(word)

	return t.addAtNode(t.Root, runes)
}

// addAtNode adds runes starting at node specified
func (t *Trie) addAtNode(n *node, runes []rune) error {

	if len(runes) == 0 {
		return nil
	}

	r := runes[0]
	cNode, err := n.AddChild(r)
	if err != nil {
		return fmt.Errorf("unable to add child: %s", err)
	}

	var cRunes []rune
	if len(runes) > 1 {
		cRunes = runes[1:]
	} else {
		// This was the last character so we should check if this is a terminator
		if cNode.IsTerm() {
			return fmt.Errorf("word already exists in trie")
		}
		cNode.MakeTerm()
		return nil
	}

	return t.addAtNode(cNode, cRunes)
}

// Words returns an array of words in the trie
func (t *Trie) Words() []string {
	words := &wordArray{
		words: []string{},
	}

	t.wordsAtNode(t.Root, "", words)

	return words.words
}

// wordsAtNode returns all words that occur after the node specified
func (t *Trie) wordsAtNode(n *node, tillThis string, words *wordArray) {
	if n.IsTerm() {
		words.add(tillThis)
	}

	for r, cNode := range n.Children() {
		t.wordsAtNode(cNode, tillThis+string(r), words)
	}
}

// Tree gives a goTree for the trie
func (t *Trie) Tree() gotree.Tree {
	tree := gotree.New("HEAD")

	t.treeAtNode(t.Root, tree)

	return tree
}

// treeAtNode gives the tree beginning from the node specified
func (t *Trie) treeAtNode(n *node, tree gotree.Tree) {
	for r, cNode := range n.Children() {
		leaf := tree.Add(string(r))
		t.treeAtNode(cNode, leaf)
	}
}

// String returns the tree as a string
func (t *Trie) String() string {
	return t.Tree().Print()
}
