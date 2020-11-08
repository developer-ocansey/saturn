package main

import (
	"github.com/stretchr/testify/mock"
)

// MockStore ..
type MockStore struct {
	mock.Mock
}
// CreateBook ..
func (m *MockStore) CreateBook(book *Book) error {
	c := m.Called(book)
	return c.Error(0)
} 

// GetBooks ..
func (m *MockStore) GetBooks() ([]*Book, error) {
	c := m.Called()
	return c.Get(0).([]*Book), c.Error(1)
}

// NewMockStore constructor
func NewMockStore() *MockStore{
	s := new(MockStore)
	store = s
	return s
}