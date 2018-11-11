package trie

// NodeAccessor defines an interface to the node
type NodeAccessor interface {
	Value() rune
	Parent() *node
	Children() Children
	IsTerm() bool
	IsRoot() bool

	AddChild(r rune) *node
	MakeTerm()
}
