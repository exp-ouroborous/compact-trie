package main

import (
	"compact-trie/trie"
	"fmt"
)

func main() {
	t := trie.NewTrie()

	t.Add("a")
	t.Add("ab")
	t.Add("abba")
	t.Add("abbc")

	fmt.Println(t.Words())
}
