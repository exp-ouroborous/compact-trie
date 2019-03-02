package trie

// ChildNodeMap defines a map from a rune to trie node and represents children of a trie node
type ChildNodeMap map[rune]*node

const nodeFound = "FOUND"
const nodeAdded = "ADDED"

// Node defines an interface to the node
type Node interface {
	Value() rune
	Parent() *node
	Children() ChildNodeMap
	IsTerm() bool
	IsRoot() bool

	AddChild(r rune) *NodeResult
	MakeTerm()

	Equal(tn *node) bool
}

// node represents a node in the trie. isTerm flags this node as terminating node i.e. there is
// a word that ends at this node, it can still have children. isRoot flags this node as the root
// node.
type node struct {
	value    rune
	parent   *node
	children ChildNodeMap
	data     interface{}

	isTerm     bool
	isRoot     bool
	childCount int
}

// NodeResult represents the result of adding a node. It includes the node found/added as well as
// a the result string which tells you whether the node was found or added
type NodeResult struct {
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

// Children gives the child nodes of the trie node
func (tn *node) Children() ChildNodeMap {
	return tn.children
}

// Data gives the data in the node
func (tn *node) Data() interface{} {
	return tn.data
}

// SetData sets the data in the node
func (tn *node) SetData(v interface{}) {
	tn.data = v
}

// HasChildren returns true is the node has children
func (tn *node) HasChildren() bool {
	return len(tn.children) > 0
}

// IsTerm returns true if the trie node is a terminating node
func (tn *node) IsTerm() bool {
	return tn.isTerm
}

// IsRoot resturn true if the trie node is the root node
func (tn *node) IsRoot() bool {
	return tn.isRoot
}

// AddChild attempts to add a node and returns a nodeResult encapsulating results of the action
func (tn *node) AddChild(r rune) *NodeResult {
	if tn.children == nil {
		tn.children = make(ChildNodeMap)
	}

	if foundN, ok := tn.children[r]; ok {
		return &NodeResult{
			result: nodeFound,
			node:   foundN,
		}
	}

	tn.children[r] = &node{
		value:    r,
		parent:   tn,
		children: make(ChildNodeMap),
		isTerm:   true,
		isRoot:   false,
	}
	tn.isTerm = false

	return &NodeResult{
		result: nodeAdded,
		node:   tn.children[r],
	}
}

// MakeTerm marks the trie node as terminating
func (tn *node) MakeTerm() {
	tn.isTerm = true
}

// Equal returns true if the receiver node and the compared to node satisfy all of the following
// - same value and flags OR both are nil
// - parents which are both nil OR have the same value
// - have the same number of children indexed using same subscripts and have same values
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
