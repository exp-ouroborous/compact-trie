[![GoDoc](https://godoc.org/github.com/exp-ouroborous/compact-trie?status.svg)](https://godoc.org/github.com/exp-ouroborous/compact-trie)

# Compact-Trie (beta)
A data structure to compactly store words or more generally ordered rune slices. See [Trie](https://en.wikipedia.org/wiki/Trie) 
for more info.

## Usage

Create a Trie with:

```Go
t := trie.New("Trie_Name")
```

Create a Trie from file of newline delimited words:

```Go
t := trie.NewFromFile("file_name","Trie_Name")
```

Add words (and data) with:

```Go
t.Add("word",data)
```

Check if a word is in the trie:

```Go
// A nil error means that the word was found
n, err := t.Find("word")
```

Remove a word from the trie:

```Go
// A nil error means that the word was successfully removed
err := t.Remove("word")
```

Visualize the trie using a linux tree:

```Go
fmt.Println(t.String())
```

## TODO
- Add fuzzy word search

## License
MIT
