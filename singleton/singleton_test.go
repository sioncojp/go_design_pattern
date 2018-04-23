package singleton

import "testing"

func TestSingleton(t *testing.T) {
	ch := make(chan interface{})
	defer close(ch)

	books := []Book{"ruby", "perl", "ruby"}

	for _, b := range books {
		go run(ch, b)
	}

	for i := 1; i <= len(books); i++ {
		<-ch
	}
}

func run(ch chan interface{}, b Book) {
	s := GetLendingList()
	ch <- s
	s.Lending(b)
}
