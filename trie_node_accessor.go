package trie

// NodeAccessor defines an interface to the node
type NodeAccessor interface {
	Value() rune
	Parent() *Node
	Children() ChildNodeMap
	IsTerm() bool
	IsRoot() bool

	AddChild(r rune) *NodeResult
	MakeTerm()

	Equal(tn *Node) bool
}
