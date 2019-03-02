package trie

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNodeFun(t *testing.T) {
	value := 'a'
	parent := &node{}
	children := childNodeMap{
		'c': &node{},
	}
	data := "some-data"
	changeData := "some-other-data"
	isTerm := false
	isRoot := false
	n := &node{
		value:    value,
		parent:   parent,
		children: children,
		data:     data,

		isTerm: isTerm,
		isRoot: isRoot,
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
			Name: "HasChildren works correctly",
			Fun:  "HasChildren",
			Out:  true,
		},
		{
			Name: "Data works correctly",
			Fun:  "Data",
			Out:  data,
		},
		{
			Name: "SetData works correctly",
			Fun:  "SetData",
			Out:  changeData,
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
			Name: "SetTerm works correctly",
			Fun:  "SetTerm",
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
		} else if test.Fun == "HasChildren" {
			op = n.HasChildren()
		} else if test.Fun == "Data" {
			op = n.Data()
		} else if test.Fun == "SetData" {
			n.SetData(changeData)
			op = n.data
		} else if test.Fun == "IsTerm" {
			op = n.IsTerm()
		} else if test.Fun == "IsRoot" {
			op = n.IsRoot()
		} else if test.Fun == "SetTerm" {
			n.SetTerm(true)
			op = n.isTerm
		}

		if !cmp.Equal(test.Out, op) {
			t.Fatalf("%s: expected and actual do not match: %s", test.Name, cmp.Diff(test.Out, op))
		}
	}
}

func TestNodeAddRemoveChild(t *testing.T) {
	var cases = []struct {
		Name         string
		NoChildMap   bool
		ExistingRune bool
		RemoveChild  bool
		Result       string
	}{
		{
			Name:   "new rune is added successfully",
			Result: nodeAdded,
		},
		{
			Name:       "new rune is added successfully when child map has not been initialized",
			NoChildMap: true,
			Result:     nodeAdded,
		},
		{
			Name:         "existing rune is found successfully",
			ExistingRune: true,
			Result:       nodeFound,
		},
		{
			Name:        "existing rune is removed succesfully",
			RemoveChild: true,
		},
	}
	for _, test := range cases {
		n := &node{
			value:    'a',
			children: make(childNodeMap),
			isTerm:   true,
		}
		cRune := 'b'
		ipRune := cRune

		if test.NoChildMap {
			n.children = nil
		}

		if test.ExistingRune {
			n = &node{
				value:    'a',
				children: make(childNodeMap),
				isTerm:   false,
			}
			n.children[cRune] = &node{
				value:  cRune,
				parent: n,
				isTerm: true,
			}
		}

		if test.RemoveChild {
			n.children['c'] = &node{}
			n.children['d'] = &node{}

			exp := &node{
				value:    'a',
				children: make(childNodeMap),
				isTerm:   true,
			}
			exp.children['d'] = &node{}

			n.RemoveChild('c')
			assert.Equal(t, exp, n)
			return
		}
		nResult := n.AddChild(ipRune)

		assert.Equal(t, test.Result, nResult.result)
		if _, ok := n.children[cRune]; !ok {
			t.Fatalf("child for %c does not exist", cRune)
		}
		if n.children[cRune].Value() != cRune {
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
			children: childNodeMap{cRune: &node{value: cRune}},
		}
		tn2 := &node{
			value:    nRune,
			parent:   &node{value: pRune},
			children: childNodeMap{cRune: &node{value: cRune}},
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
			pNode := tn2.parent.(*node)
			pNode.value = 'x'
		}

		if test.ChildDiff == "empty" {
			tn2.children = make(childNodeMap)
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

		var tn2Ip Node
		if tn2 != nil {
			tn2Ip = tn2
		}
		op := tn1.Equal(tn2Ip)
		assert.Equal(t, test.IsEqual, op)
	}
}
