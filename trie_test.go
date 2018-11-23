package trie

import (
	"fmt"
	"testing"

	"github.com/disiqueira/gotree"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHasWord(t *testing.T) {

	tr := New("")
	tr.Root.children['a'] = &node{
		value:    'a',
		parent:   tr.Root,
		children: make(childNodeMap),
	}
	tr.Root.children['a'].children['b'] = &node{
		value:  'b',
		parent: tr.Root.children['a'],
		isTerm: true,
	}
	tr.Root.children['b'] = &node{
		value:  'b',
		parent: tr.Root,
		isTerm: true,
	}
	inTrie := "ab"
	notInTrie := "abcd"
	inTrieButNotTerm := "a"

	var cases = []struct {
		Name      string
		Input     string
		ExpectErr string
	}{
		{
			Name: "Word in trie is found successfully",
		},
		{
			Name:      "empty word throws error",
			Input:     "empty",
			ExpectErr: "no string to find",
		},
		{
			Name:      "word not in trie throws error",
			Input:     "not-in-trie",
			ExpectErr: fmt.Sprintf("word %s not found", notInTrie),
		},
		{
			Name:      "word in trie but not as terminating throws error",
			Input:     "in-trie-but-not-term",
			ExpectErr: fmt.Sprintf("word %s not found", inTrieButNotTerm),
		},
	}

	for _, test := range cases {
		var err error

		ipWord := inTrie

		if test.Input == "empty" {
			ipWord = ""
		} else if test.Input == "not-in-trie" {
			ipWord = notInTrie
		} else if test.Input == "in-trie-but-not-term" {
			ipWord = inTrieButNotTerm
		}
		err = tr.Find(ipWord)

		if test.ExpectErr != "" {
			require.NotEmpty(t, err)
			assert.Contains(t, err.Error(), test.ExpectErr)
			continue
		}
		assert.Empty(t, err)

	}
}

func TestAddWord(t *testing.T) {
	inTrie := "ab"
	notInTrie := "abcd"
	inTrieButNotTerm := "a"

	var cases = []struct {
		Name      string
		Input     string
		ExpectErr string
	}{
		{
			Name: "Word not in trie is added successfully",
		},
		{
			Name:      "empty word throws error",
			Input:     "empty",
			ExpectErr: "no string to add",
		},
		{
			Name:      "word in trie throws error",
			Input:     "in-trie",
			ExpectErr: "word already exists in trie",
		},
		{
			Name:  "word in trie but not as terminating is added successfully",
			Input: "in-trie-but-not-term",
		},
	}

	for _, test := range cases {
		var err error
		tr := New("test")
		tr.Root.children['a'] = &node{
			value:    'a',
			parent:   tr.Root,
			children: make(childNodeMap),
		}
		tr.Root.children['a'].children['b'] = &node{
			value:  'b',
			parent: tr.Root.children['a'],
			isTerm: true,
		}
		tr.Root.children['b'] = &node{
			value:  'b',
			parent: tr.Root,
			isTerm: true,
		}

		ipWord := notInTrie

		if test.Input == "empty" {
			ipWord = ""
		} else if test.Input == "in-trie" {
			ipWord = inTrie
		} else if test.Input == "in-trie-but-not-term" {
			ipWord = inTrieButNotTerm
		}
		err = tr.Add(ipWord)

		if test.ExpectErr != "" {
			require.NotEmpty(t, err)
			assert.Contains(t, err.Error(), test.ExpectErr, test.Name)
			continue
		}
		assert.Empty(t, err, test.Name)

	}
}

func TestHasWords(t *testing.T) {
	var cases = []struct {
		Name string
	}{
		{
			Name: "all words in the trie are returned successfully",
		},
	}

	for _, test := range cases {
		tr := New("test")
		tr.Root.children['a'] = &node{
			value:    'a',
			parent:   tr.Root,
			children: make(childNodeMap),
		}
		tr.Root.children['a'].children['b'] = &node{
			value:  'b',
			parent: tr.Root.children['a'],
			isTerm: true,
		}
		tr.Root.children['b'] = &node{
			value:  'b',
			parent: tr.Root,
			isTerm: true,
		}

		words := tr.Words()
		expWords := []string{"ab", "b"}
		assert.ElementsMatch(t, expWords, words, test.Name)

	}
}

func TestTree(t *testing.T) {
	var cases = []struct {
		Name string
	}{
		{
			Name: "tree is generated correctly",
		},
	}

	for _, test := range cases {
		trieName := "test"
		tr := New(trieName)
		tr.Root.children['a'] = &node{
			value:    'a',
			parent:   tr.Root,
			children: make(childNodeMap),
		}
		tr.Root.children['a'].children['b'] = &node{
			value:  'b',
			parent: tr.Root.children['a'],
			isTerm: true,
		}
		tr.Root.children['b'] = &node{
			value:  'b',
			parent: tr.Root,
			isTerm: true,
		}

		expTree := gotree.New(trieName)
		expTree.Add("a").Add("b")
		expTree.Add("b")
		tree := tr.Tree()

		// TODO: This seems to throw a comparison error once in a while
		assert.Equal(t, expTree, tree, test.Name)
	}
}
