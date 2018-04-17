package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	b := &Bookshelf{}
	books := []string{"a", "b", "c"}
	for _, v := range books {
		b.Add(Book{v})
	}
	for i := b.Iterator(); i.HasNext(); {
		v := i.Next()
		fmt.Println(v)
	}
}
