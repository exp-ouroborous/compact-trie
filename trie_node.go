package trie

// ChildNodeMap defines a map from a rune to trie node and represents children of a trie node
type ChildNodeMap map[rune]*Node

const nodeFound = "FOUND"
const nodeAdded = "ADDED"

// Node represents a Node in the trie. isTerm flags this Node as terminating Node i.e. there is
// a word that ends at this Node. This Node can still have children. isRoot flags this Node as
// the root Node.
type Node struct {
	value    rune
	parent   *Node
	children ChildNodeMap
	data     interface{}

	isTerm bool
	isRoot bool
}

// NodeResult represents the result of adding a node. It includes the node found/added as well as
// a the result string which tells you whether the node was found or added
type NodeResult struct {
	result string
	*Node
}

// Value gives the rune in the trie node
func (tn *Node) Value() rune {
	return tn.value
}

// Parent gives the parent of the trie node
func (tn *Node) Parent() *Node {
	return tn.parent
}

// Children gives the child nodes of the trie node
func (tn *Node) Children() ChildNodeMap {
	return tn.children
}

// Data gives the data in the node
func (tn *Node) Data() interface{} {
	return tn.data
}

// SetData sets the data in the node
func (tn *Node) SetData(v interface{}) {
	tn.data = v
}

// HasChildren returns true is the node has children
func (tn *Node) HasChildren() bool {
	return len(tn.children) > 0
}

// IsTerm returns true if the trie node is a terminating node
func (tn *Node) IsTerm() bool {
	return tn.isTerm
}

// IsRoot resturn true if the trie node is the root node
func (tn *Node) IsRoot() bool {
	return tn.isRoot
}

// AddChild attempts to add a node and returns a nodeResult encapsulating results of the action
func (tn *Node) AddChild(r rune) *NodeResult {
	if tn.children == nil {
		tn.children = make(ChildNodeMap)
	}

	if foundN, ok := tn.children[r]; ok {
		return &NodeResult{
			result: nodeFound,
			Node:   foundN,
		}
	}

	tn.children[r] = &Node{
		value:    r,
		parent:   tn,
		children: make(ChildNodeMap),
		isTerm:   true,
		isRoot:   false,
	}
	tn.isTerm = false

	return &NodeResult{
		result: nodeAdded,
		Node:   tn.children[r],
	}
}

// MakeTerm marks the trie node as terminating
func (tn *Node) MakeTerm() {
	tn.isTerm = true
}

// Equal returns true if the receiver node and the compared to node satisfy all of the following
// - same value and flags OR both are nil
// - parents which are both nil OR have the same value
// - have the same number of children indexed using same subscripts and have same values
func (tn *Node) Equal(tn2 *Node) bool {
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
