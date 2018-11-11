package trie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TrieNodeTestSuite struct {
	suite.Suite

	nodeValue  rune
	nodeIsRoot bool
	nodeIsTerm bool

	tn *node

	compareTo *node
}

func (s *TrieNodeTestSuite) SetupTest() {

	// Setup the node
	s.nodeValue = 'a'
	s.tn = &node{
		value:    s.nodeValue,
		children: make(Children),
	}
	tnParent := &node{
		children: Children{
			s.nodeValue: s.tn,
		},
	}
	s.tn.parent = tnParent
	tnChild1 := &node{
		value: 'a',
	}
	tnChild2 := &node{
		value: 'b',
	}
	tnChildren := make(Children)
	tnChildren[tnChild1.value] = tnChild1
	tnChildren[tnChild2.value] = tnChild2
	s.tn.children = tnChildren
}

func (s *TrieNodeTestSuite) TestValue() {
	s.Equal(s.nodeValue, s.tn.Value())
}

func (s *TrieNodeTestSuite) TestParent() {
	s.Equal(s.tn.parent, s.tn.Parent())
}

func (s *TrieNodeTestSuite) TestChildren() {
	s.Equal(s.tn.children, s.tn.Children())
}

func (s *TrieNodeTestSuite) TestIsRoot() {
	s.Equal(s.tn.isRoot, s.tn.IsRoot())
}

func (s *TrieNodeTestSuite) TestIsTerm() {
	s.Equal(s.tn.isTerm, s.tn.IsTerm())
}

func (s *TrieNodeTestSuite) TestAddChild() {
	s.tn.AddChild('z')
	s.Equal('z', s.tn.Children()['z'].value)

	_, err := s.tn.AddChild(0)
	s.Equal("rune is required", err.Error())

	_, err = s.tn.AddChild('z')
	s.Equal(fmt.Sprintf("child node for %c alread exists", 'z'), err.Error())
}
func (s *TrieNodeTestSuite) TestMakeTerm() {
	s.nodeIsTerm = true
	s.tn.MakeTerm()

	s.Equal(s.nodeIsTerm, s.tn.IsTerm())
}

func TestTrieNodeTestSuite(t *testing.T) {
	suite.Run(t, new(TrieNodeTestSuite))
}
