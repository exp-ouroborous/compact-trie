[![GoDoc](https://godoc.org/github.com/exp-ouroborous/compact-trie?status.svg)](https://godoc.org/github.com/exp-ouroborous/compact-trie)

# Compact-Trie (beta)
A data structure to compactly store words or more generally ordered rune slices. See [Trie](https://en.wikipedia.org/wiki/Trie) for more info. 

## Usage

Create a Trie with:

```Go
t := trie.New("Trie_Name")
```

Add words with:

```Go
t.Add("word")
```

Check if a word is in the trie:

```Go
// A nil error means that the word was found
err := t.Find("word")
```

Visualize the trie using a linux tree:

```Go
fmt.Println(t.String())
```

## TODOs
- Add a method to remove a word
- Add metadata for the trie

## License
MIT
