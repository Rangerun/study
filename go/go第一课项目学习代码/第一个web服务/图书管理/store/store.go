package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	Id 		string	`json:"id"`
	Name	string	`json:"name"`
}

type Store interface {
	Create(*Book) error
	GetBook(string) (*Book, error)
}