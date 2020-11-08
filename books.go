package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
)

// Book ... create a constructor
type Book struct {
	Title string `json:"title"`
	Category string `json:"category"`
	Author string `json:"author"`
	Description string `json:"description"`
	Ratings int `json:"ratings"`
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
	book := &Book{}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("An error occurred: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rt, err := strconv.Atoi(r.Form.Get("ratings"))
	if err != nil {
		fmt.Println(fmt.Errorf("An error occurred: %v", err))
		w.WriteHeader(http.StatusInternalServerError) // Hide errors from users
		return
	}

	book.Title = r.Form.Get("title")
	book.Category = r.Form.Get("category")
	book.Author = r.Form.Get("author")
	book.Description = r.Form.Get("description")
	book.Ratings = rt
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	books, err := store.GetBooks()
	if err != nil {
		fmt.Println(fmt.Errorf("An error occurred: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	b, err := json.Marshal(books)
	w.Write(b)
}