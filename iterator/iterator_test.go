package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	b := &Bookshelf{}
	books := []Book{
		{"a"},
		{"b"},
		{"c"},
	}
	b.Books = books
	for i := b.Iterator(); i.HasNext(); {
		v := i.Next()
		fmt.Println(v)
	}
}
