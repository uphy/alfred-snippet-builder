# Alfred Snippets Builder

```go
package main

import "github.com/uphy/alfred-snippet-builder"

func main() {
	s := snippet.New("prefix", "suffix")
	s.Add("name", "snippet", "keyword")
	if err := s.Save("sample.alfredsnippets"); err != nil {
		panic(err)
	}
}

```