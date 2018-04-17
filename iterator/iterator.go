package iterator

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	Next() interface{}
	HasNext() bool
}

type Book struct {
	Title string
}

type Bookshelf struct {
	Books []Book
}

type BookshelfIterator struct {
	Bookshelf *Bookshelf
	index     int
}

func (s *Bookshelf) Add(b Book) {
	s.Books = append(s.Books, b)
}

func (s *Bookshelf) Iterator() Iterator {
	return &BookshelfIterator{Bookshelf: s}
}

func (s *BookshelfIterator) HasNext() bool {
	return s.index < len(s.Bookshelf.Books)
}

func (s *BookshelfIterator) Next() interface{} {
	b := s.Bookshelf.Books[s.index]
	s.index++
	return b.Title
}
