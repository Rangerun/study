package store

import (
	mystore "bookserver/store"
	//factory "bookserver/store/factory"
	"sync"
)

type MemStore struct {
	sync.RWMutex
	Books map[string]*mystore.Book
}

func (mem *MemStore) Create(book *mystore.Book) error {
	if _, ok := mem.Books[book.Id]; ok {
		return mystore.ErrExist
	}
	nBook := *book
	mem.Books[book.Id] = &nBook
	return nil
}

func (mem *MemStore) GetBook(id string) (*mystore.Book, error) {
	if book, ok := mem.Books[id]; ok {
		return book, nil
	}
	return nil, mystore.ErrNotFound
}