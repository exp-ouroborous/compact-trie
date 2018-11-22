package main

import (
	"compact-trie/trie"
	"fmt"
)

func main() {
	t := trie.NewTrie("Trie")
	arr := []string{"abba", "cat", "cab", "can", "abb"}
	addWords(arr, t)

	fmt.Println(t.Words())
	fmt.Println(t.String())
}

func addWords(words []string, t *trie.Trie) {
	for i := range words {
		err := t.Add(words[i])
		if err != nil {
			fmt.Printf("could not add %s: %s\n", words[i], err)
		}
	}
}
