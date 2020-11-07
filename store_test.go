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

func (s *storeSuit) TestCreateBook(){
	s.store.CreateBooks(&Book{
		Title: "Things fall apart",
		Category: "tale",
		Author: "Chinauh Achebe",
		Ratings: 5,
	})
}