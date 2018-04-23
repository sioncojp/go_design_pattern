package singleton

import (
	"fmt"
	"sync"
)

type lendingList struct {
	Books []Book
}

var lendingListInstance *lendingList
var once sync.Once

func GetLendingList() *lendingList {
	once.Do(func() {
		// 初期化
		lendingListInstance = &lendingList{}
	})
	return lendingListInstance
}

type Book string

func (l *lendingList) Lending(b Book) {
	for _, v := range l.Books {
		if v == b {
			fmt.Printf("%v貸りられてるよ\n", b)
			return
		}
	}
	l.Books = append(l.Books, b)
	fmt.Printf("%vを借りるね\n", b)
}
