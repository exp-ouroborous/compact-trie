package trie

// Children defines a map from a rune to trie node and represents children of a trie node
type Children map[rune]*node

const nodeFound = "FOUND"
const nodeAdded = "ADDED"

// node represents a node in the trie
type node struct {
	value    rune
	parent   *node
	children Children

	isTerm bool
	isRoot bool
}

type nodeResult struct {
	result string
	*node
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

// AddChild attempts to add a node and returns a nodeResult encapsulating results of the action
func (tn *node) AddChild(r rune) *nodeResult {
	if tn.children == nil {
		tn.children = make(Children)
	}

	if foundN, ok := tn.children[r]; ok {
		return &nodeResult{
			result: nodeFound,
			node:   foundN,
		}
	}

	tn.children[r] = &node{
		value:    r,
		parent:   tn,
		children: make(Children),
		isTerm:   true,
		isRoot:   false,
	}
	tn.isTerm = false

	return &nodeResult{
		result: nodeAdded,
		node:   tn.children[r],
	}
}

// MakeTerm rks the trie node as terminating
func (tn *node) MakeTerm() {
	tn.isTerm = true
}

func (tn *node) Equal(tn2 *node) bool {
	// Two nils are equal
	if tn == nil || tn2 == nil {
		if tn == nil && tn2 == nil {
			return true
		}
		return false
	}

	// Compare node specific values
	if tn.value != tn2.value || tn.isRoot != tn2.isRoot || tn.isTerm != tn2.isTerm {
		return false
	}

	// Compare parents
	if tn.parent == nil || tn2.parent == nil {
		if tn.parent == nil && tn2.parent == nil {
			return true
		}
		return false
	}
	if tn.parent.value != tn2.parent.value {
		return false
	}

	// Compare children
	for cRune, cNode := range tn.children {
		c2Node, ok := tn2.children[cRune]
		if !ok {
			return false
		}
		if cNode.value != c2Node.value {
			return false
		}
	}
	for c2Rune := range tn2.children {
		_, ok := tn.children[c2Rune]
		if !ok {
			return false
		}
	}

	return true
}
