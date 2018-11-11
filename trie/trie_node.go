package trie

import (
	"fmt"
)

// Children defines a map from a rune to trie node and represents children of a trie node
type Children map[rune]*node

// node represents a node in the trie
type node struct {
	value    rune
	parent   *node
	children Children

	isTerm bool
	isRoot bool
}

// Value gives the rune in the trie node
func (tn *node) Value() rune {
	return tn.value
}

// Parent gives the parent of the trie node
func (tn *node) Parent() *node {
	return tn.parent
}

// Children gives the children node of the trie node
func (tn *node) Children() Children {
	return tn.children
}

// IsTerm returns true if the trie node is a terminating
func (tn *node) IsTerm() bool {
	return tn.isTerm
}

// IsRoot resturn true if the trie node is the root node
func (tn *node) IsRoot() bool {
	return tn.isRoot
}

// AddChild adds a node for a rune and returns that node. If the rune is empty or if the
// child node already exists an error is returned
func (tn *node) AddChild(r rune) (*node, error) {
	if r == 0 {
		return nil, fmt.Errorf("rune is required")
	}
	if _, ok := tn.children[r]; ok {
		return nil, fmt.Errorf("child node for %c alread exists", r)
	}

	tn.children[r] = &node{
		value:    r,
		parent:   tn,
		children: make(Children),
	}

	return tn.children[r], nil
}

// MakeTerm amrks the trie node as terminating
func (tn *node) MakeTerm() {
	tn.isTerm = true
}

// ###UNTESTED CODE###

func isSame(a, b *node) error {
	// Compare values in the current node
	if err := hasSameValue(a, b); err != nil {
		return fmt.Errorf("values are not same: %s", err)
	}

	// Compare values in the parent node
	if err := hasSameValue(a.parent, b.parent); err != nil {
		return fmt.Errorf("parent's values are not same: %s", err)
	}

	// Are all of a's children in b and have the same value?
	for r, cNode := range a.children {
		bChildNode, ok := b.children[r]
		if !ok {
			return fmt.Errorf("node missing in second object: %c", r)
		}
		if err := hasSameValue(cNode, bChildNode); err != nil {
			return fmt.Errorf("nodes with same key %c do not match in values: %s", r, err)
		}
	}

	// Are all of b's children in a?
	for r := range b.children {
		if _, ok := a.children[r]; !ok {
			return fmt.Errorf("extra node in second object: %c", r)
		}
		// No need to check the children as they have been checked in the previous loop
	}

	return nil
}

func hasSameValue(a, b *node) error {
	if a.value != b.value {
		return fmt.Errorf("values do not match: %c vs %c", a.value, b.value)
	}

	if a.isTerm != b.isTerm {
		return fmt.Errorf("terminal values do not match: %t vs %t", a.isTerm, b.isTerm)
	}

	if a.isRoot != b.isRoot {
		return fmt.Errorf("root values do not match: %t vs %t", a.isRoot, b.isRoot)
	}

	return nil
}

func deepCopy(src *node, parent *node) *node {
	copy := &node{
		value:    src.value,
		parent:   parent,
		children: make(Children),

		isRoot: src.isRoot,
		isTerm: src.isTerm,
	}

	for _, cNode := range src.children {
		copy.children[cNode.value] = deepCopy(cNode, copy)
	}

	return copy
}
