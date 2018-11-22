package trie

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/google/go-cmp/cmp"
)

func TestNodeFun(t *testing.T) {
	value := 'a'
	parent := &node{}
	children := Children{
		'c': &node{},
	}
	isTerm := false
	isRoot := false
	n := &node{
		value:    value,
		parent:   parent,
		children: children,
		isTerm:   isTerm,
		isRoot:   isRoot,
	}

	var cases = []struct {
		Name string
		Fun  string
		Out  interface{}
	}{
		{
			Name: "Value works correctly",
			Fun:  "Value",
			Out:  value,
		},
		{
			Name: "Parent works correctly",
			Fun:  "Parent",
			Out:  parent,
		},
		{
			Name: "Children works correctly",
			Fun:  "Children",
			Out:  children,
		},
		{
			Name: "IsTerm works correctly",
			Fun:  "IsTerm",
			Out:  isTerm,
		},
		{
			Name: "IsRoot works correctly",
			Fun:  "IsRoot",
			Out:  isRoot,
		},
		{
			Name: "MakeTerm works correctly",
			Fun:  "MakeTerm",
			Out:  true,
		},
	}

	for _, test := range cases {
		var op interface{}

		if test.Fun == "Value" {
			op = n.Value()
		} else if test.Fun == "Parent" {
			op = n.Parent()
		} else if test.Fun == "Children" {
			op = n.Children()
		} else if test.Fun == "IsTerm" {
			op = n.IsTerm()
		} else if test.Fun == "IsRoot" {
			op = n.IsRoot()
		} else if test.Fun == "MakeTerm" {
			n.MakeTerm()
			op = n.isTerm
		}

		if !cmp.Equal(test.Out, op) {
			t.Fatalf("expected and actual do not match: %s", cmp.Diff(test.Out, op))
		}
	}
}

func TestNodeAddChild(t *testing.T) {
	var cases = []struct {
		Name         string
		ExistingRune bool
		Result       string
	}{
		{
			Name:   "new rune is added successfully",
			Result: nodeAdded,
		},
		{
			Name:         "existing rune is found successfully",
			ExistingRune: true,
			Result:       nodeFound,
		},
	}
	for _, test := range cases {
		n := &node{
			value:    'a',
			children: make(Children),
			isTerm:   true,
		}
		cRune := 'b'
		ipRune := cRune

		if test.ExistingRune {
			n = &node{
				value:    'a',
				children: make(Children),
				isTerm:   false,
			}
			n.children[cRune] = &node{
				value:  cRune,
				parent: n,
				isTerm: true,
			}
		}

		nResult := n.AddChild(ipRune)

		assert.Equal(t, test.Result, nResult.result)
		if _, ok := n.children[cRune]; !ok {
			t.Fatalf("child for %c does not exist", cRune)
		}
		if n.children[cRune].value != cRune {
			t.Fatalf("child for %c does not have correct value", cRune)
		}
		if n.isTerm != false {
			t.Fatalf("isTerm is not false")
		}
	}
}

func TestNodeEqual(t *testing.T) {
	var cases = []struct {
		Name       string
		NilNode    string //both,one
		InNodeDiff string //value,root,term
		ParentDiff string //nil,value
		ChildDiff  string //empty,value,extra
		IsEqual    bool
		ExpectErr  string
	}{
		{
			Name:    "equal nodes are equal",
			IsEqual: true,
		},
		{
			Name:    "nil nodes are equal",
			NilNode: "both",
			IsEqual: true,
		},
		{
			Name:    "nil and non-nil nodes are not equal",
			NilNode: "one",
			IsEqual: false,
		},
		{
			Name:    "nil and non-nil nodes are not equal",
			NilNode: "one",
			IsEqual: false,
		},
		{
			Name:       "nodes with different values are not equal",
			InNodeDiff: "value",
			IsEqual:    false,
		},
		{
			Name:       "nodes with different root flags are not equal",
			InNodeDiff: "root",
			IsEqual:    false,
		},
		{
			Name:       "nodes with different term flags are not equal",
			InNodeDiff: "term",
			IsEqual:    false,
		},
		{
			Name:       "one node with nil parent is not equal",
			ParentDiff: "nil",
			IsEqual:    false,
		},
		{
			Name:       "node with different parent is not equal",
			ParentDiff: "value",
			IsEqual:    false,
		},
		{
			Name:      "node with no child is not equal",
			ChildDiff: "empty",
			IsEqual:   false,
		},
		{
			Name:      "node with different child is not equal",
			ChildDiff: "value",
			IsEqual:   false,
		},
		{
			Name:      "node with extra child is not equal",
			ChildDiff: "extra",
			IsEqual:   false,
		},
	}

	for _, test := range cases {
		var err error

		nRune := 'n'
		pRune := 'p'
		cRune := 'c'
		tn1 := &node{
			value:    nRune,
			parent:   &node{value: pRune},
			children: Children{cRune: &node{value: cRune}},
		}
		tn2 := &node{
			value:    nRune,
			parent:   &node{value: pRune},
			children: Children{cRune: &node{value: cRune}},
		}

		if test.NilNode == "both" {
			tn1 = nil
			tn2 = nil
		} else if test.NilNode == "one" {
			tn2 = nil
		}

		if test.InNodeDiff == "value" {
			tn2.value = 'x'
		} else if test.InNodeDiff == "root" {
			tn2.isRoot = true
		} else if test.InNodeDiff == "term" {
			tn2.isTerm = true
		}

		if test.ParentDiff == "nil" {
			tn2.parent = nil
		} else if test.ParentDiff == "value" {
			tn2.parent.value = 'x'
		}

		if test.ChildDiff == "empty" {
			tn2.children = make(Children)
		} else if test.ChildDiff == "value" {
			tn2.children['c'] = &node{value: 'd'}
		} else if test.ChildDiff == "extra" {
			tn2.children['d'] = &node{value: 'd'}
		}

		if test.ExpectErr != "" {
			require.NotEmpty(t, err)
			assert.Contains(t, err.Error(), test.ExpectErr)
			return
		}
		assert.Empty(t, err)

		op := tn1.Equal(tn2)
		assert.Equal(t, test.IsEqual, op)
	}
}
