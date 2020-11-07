package main
import (
	"fmt"
	"database/sql"
)

const (
	insertBooksQuery = "INSERT INTO bookStore( title, category, description, ratings) VALUES ($1, $2, $3, $4)"
	 getBooksQuery = "SELECT title, category, description, ratings from bookStore"
)

// Store ..
type Store interface {
	CreateBooks(books *Book) error
	GetBooks() ([]*Book, error)
}

type dbStore struct {
	db *sql.DB
}

func (ds *dbStore) CreateBooks(book *Book) error {
	_, err := ds.db.Query(insertBooksQuery, book.Title, book.Category, book.Description, book.Ratings)
	return err
}

func(ds *dbStore) GetBooks() ([]*Book, error) {
	rows, err := ds.db.Query(getBooksQuery)
	if err != nil {
		return nil, fmt.Errorf("cannot execute GetBooks query: %w",err)
	}
	defer rows.Close()

	books :=[]*Book{}
	for rows.Next() {
		book := &Book{}
		if err := rows.Scan(&book.Title, book.Category, book.Description, book.Ratings); err != nil {
			return nil,fmt.Errorf("cannot scan rows: %w", err)
		}
		books = append(books, book)
	}
	return books, nil
}

var store Store

// NewStore ...
func NewStore (s Store){
	store = s
}