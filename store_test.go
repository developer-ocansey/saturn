package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type storeSuit struct {
	suite.Suite

	store *dbStore
	db *sql.DB
}

func (s *storeSuit) SetupSuit(){
	conn := "dbname=testBookStore sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *storeSuit) startTest(){
	if _, err := s.db.Query("DELETE from bookStore");err!=nil{
		s.T().Fatal(err)
	}
}

func (s *storeSuit) endTestSuite() {
	s.db.Close()
}

func TestStoreSuite(t *testing.T){
	s:= &storeSuit{}
	suite.Run(t, s)
}

func (s *storeSuit) TestCreateBook() {
	s.store.CreateBook(&Book{
		Title: "Things fall apart",
		Category: "tale",
		Author: "Chinauh Achebe",
		Description: "new books",
		Ratings: 5,
	})
	res, err := s.db.Query(`SELECT COUNT(*) from bookStore where title = 'Things fall apart' and author = 'Chinauh Achebe'`)
	if err != nil {
		s.T().Fatal(err)
	}

	var count int 
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Fatal(err)
		}
	}

	if count != 1 {
		s.T().Errorf("expected 1 but got %d", count)
	}
}

func (s *storeSuit) TestGetBook() {
	_, err := s.db.Query(`INSERT INTO bookStore(title, category, description, ratings) VALUES ('Things fall apart', 'tale', 'Chinauh Achebe', 'new books', '5')`)
	if err != nil {
		s.T().Fatal(err)
	}
	books, err := s.store.GetBooks()
	if err != nil {
		s.T().Fatal(err)
	}

	numberOfBooks := len(books)
	if numberOfBooks != 1 {
		s.T().Errorf("expected 1 book but got %d", numberOfBooks)
	}

	expectedBooks := Book{"Things fall apart", "tale", "Chinauh Achebe", "new books", 5 }
	if *books[0] != expectedBooks {
		s.T().Errorf("expected %v but got %v", expectedBooks, *books[0])
	}

}