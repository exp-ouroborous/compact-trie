package trie

import (
	"fmt"
	"testing"

	"github.com/disiqueira/gotree/v3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFromFile(t *testing.T) {
	expTr := New("test")
	expTr.Root.Children()['a'] = &node{
		value:    'a',
		parent:   expTr.Root,
		children: make(childNodeMap),
	}
	expTr.Root.Children()['a'].Children()['b'] = &node{
		value:  'b',
		parent: expTr.Root.Children()['a'],
		isTerm: true,
	}
	expTr.Root.Children()['b'] = &node{
		value:  'b',
		parent: expTr.Root,
		isTerm: true,
	}

	file := "testdata/wordtest.txt"
	unreadableFile := "testdata/this-file-is-not-here"

	var cases = []struct {
		Name      string
		File      string //empty,unreadable
		ExpectErr string
	}{
		{
			Name: "trie is correctly loaded from file",
		},
		{
			Name:      "empty file name throws an error",
			File:      "empty",
			ExpectErr: "file is required",
		},
		{
			Name:      "unreadable file throws an error",
			File:      "unreadable",
			ExpectErr: "could not read file",
		},
	}

	for _, test := range cases {
		var err error
		ipFile := file

		if test.File == "empty" {
			ipFile = ""
		} else if test.File == "unreadable" {
			ipFile = unreadableFile
		}
		tr, err := NewFromFile(ipFile, "test")

		if test.ExpectErr != "" {
			require.NotEmpty(t, err, test.Name)
			assert.Contains(t, err.Error(), test.ExpectErr, test.Name)
			continue
		}
		assert.Empty(t, err, test.Name)

		// TODO Is it kosher to use another function in the unit test??
		assert.Equal(t, true, tr.Equal(expTr), test.Name)
	}
}

func TestFind(t *testing.T) {
	tr := New("")
	tr.Root.Children()['a'] = &node{
		value:    'a',
		parent:   tr.Root,
		children: make(childNodeMap),
	}
	tr.Root.Children()['a'].Children()['b'] = &node{
		value:  'b',
		parent: tr.Root.Children()['a'],
		isTerm: true,
	}
	tr.Root.Children()['b'] = &node{
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
		_, err = tr.Find(ipWord)

		if test.ExpectErr != "" {
			require.NotEmpty(t, err)
			assert.Contains(t, err.Error(), test.ExpectErr)
			continue
		}
		assert.Empty(t, err)

	}
}

func TestAdd(t *testing.T) {
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
		tr, err := NewFromFile("testdata/wordtest.txt", "test")
		require.Empty(t, err, test.Name)

		ipWord := notInTrie

		if test.Input == "empty" {
			ipWord = ""
		} else if test.Input == "in-trie" {
			ipWord = inTrie
		} else if test.Input == "in-trie-but-not-term" {
			ipWord = inTrieButNotTerm
		}
		_, err = tr.Add(ipWord, "")

		if test.ExpectErr != "" {
			require.NotEmpty(t, err)
			assert.Contains(t, err.Error(), test.ExpectErr, test.Name)
			continue
		}
		assert.Empty(t, err, test.Name)

	}
}

func TestRemove(t *testing.T) {
	inTrie := "abc"
	inTrieButNotTerm := "a"

	var cases = []struct {
		Name      string
		Input     string
		ExpectErr string
	}{
		{
			Name: "Word in trie is removed successfully",
		},
		{
			Name:      "empty word throws error",
			Input:     "empty",
			ExpectErr: "could not find word",
		},
		{
			Name:      "word not in trie throws error",
			Input:     "in-trie-but-not-term",
			ExpectErr: "could not find word ",
		},
	}

	for _, test := range cases {
		var err error
		tr := New("test")
		tr.Root.Children()['a'] = &node{
			value:    'a',
			parent:   tr.Root,
			children: make(childNodeMap),
		}
		tr.Root.Children()['a'].Children()['b'] = &node{
			value:    'b',
			parent:   tr.Root.Children()['a'],
			isTerm:   true,
			children: make(childNodeMap),
		}
		tr.Root.Children()['a'].Children()['b'].Children()['c'] = &node{
			value:  'c',
			parent: tr.Root.Children()['a'].Children()['b'],
			isTerm: true,
		}
		tr.Root.Children()['b'] = &node{
			value:  'b',
			parent: tr.Root,
			isTerm: true,
		}

		ipWord := inTrie

		if test.Input == "empty" {
			ipWord = ""
		} else if test.Input == "in-trie-but-not-term" {
			ipWord = inTrieButNotTerm
		}
		err = tr.Remove(ipWord)

		if test.ExpectErr != "" {
			require.NotEmpty(t, err, test.Name+": "+test.ExpectErr)
			assert.Contains(t, err.Error(), test.ExpectErr, test.Name)
			continue
		}
		assert.Empty(t, err, test.Name)

		_, err = tr.Find(ipWord)
		require.NotEmpty(t, err, test.Name+": "+"did not expect to find removed word")
	}
}

func TestWords(t *testing.T) {
	var cases = []struct {
		Name string
	}{
		{
			Name: "all words in the trie are returned successfully",
		},
	}

	for _, test := range cases {
		tr := New("test")
		tr.Root.Children()['a'] = &node{
			value:    'a',
			parent:   tr.Root,
			children: make(childNodeMap),
		}
		tr.Root.Children()['a'].Children()['b'] = &node{
			value:  'b',
			parent: tr.Root.Children()['a'],
			isTerm: true,
		}
		tr.Root.Children()['b'] = &node{
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
		Name       string
		StringTest bool
	}{
		{
			Name: "tree is generated correctly",
		},
		{
			Name:       "tree string is generated correctly",
			StringTest: true,
		},
	}

	for _, test := range cases {
		trieName := "test"
		tr := New(trieName)
		tr.Root.Children()['a'] = &node{
			value:      'a',
			parent:     tr.Root,
			children:   make(childNodeMap),
			childCount: 1,
		}
		tr.Root.Children()['a'].Children()['b'] = &node{
			value:      'b',
			parent:     tr.Root.Children()['a'],
			childCount: 1,
			isTerm:     true,
		}
		tr.Root.Children()['b'] = &node{
			value:      'b',
			parent:     tr.Root,
			childCount: 1,
			isTerm:     true,
		}

		expTree := gotree.New(trieName)
		expTree.Add("a").Add("b")
		expTree.Add("b")
		tree := tr.Tree()

		if test.StringTest {
			expTree := "test\n├── a\n│   └── b\n└── b\n"
			tree := tr.String()
			assert.Equal(t, expTree, tree, test.Name)
		}

		assert.Equal(t, expTree, tree, test.Name)
	}
}
