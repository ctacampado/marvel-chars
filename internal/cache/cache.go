package cache

import (
	"fmt"
)

// Elem represents the element being stored
type Elem interface {
	Print()
	GetID() string
}

// Store represents our cache which is a key-value store
type Store struct {
	Data map[string]Elem `json:"Data"`
}

// DATAFILE contains the filename
var DATAFILE = "./datafile.json"

// New returns a new store
func New() *Store {
	return &Store{Data: make(map[string]Elem)}
}

// LOOKUP returns element by key
func (s *Store) LOOKUP(k string) *Elem {
	_, ok := s.Data[k]
	if ok {
		n := s.Data[k]
		return &n
	}
	return nil
}

// ADD elements to the store
func (s *Store) ADD(k string, n Elem) bool {
	if k == "" {
		return false
	}

	if s.LOOKUP(k) == nil {
		s.Data[k] = n
		return true
	}
	return false
}

// DELETE removes elem from store
func (s *Store) DELETE(k string) bool {
	if s.LOOKUP(k) != nil {
		delete(s.Data, k)
		return true
	}
	return false
}

// CHANGE modifies the content of store located in data[k]
func (s *Store) CHANGE(k string, n Elem) bool {
	s.Data[k] = n
	return true
}

// PRINT displays the content of store
func (s *Store) PRINT() {
	for k, v := range s.Data {
		fmt.Printf("k: %s | v:", k)
		v.Print()
	}
}
