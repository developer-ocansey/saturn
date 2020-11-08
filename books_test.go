package main

import (
	"bytes"
	"encoding/json"
	"testing"
	"net/http"
	"net/url"
	"strconv"
	"net/http/httptest"
)
func TestGetBookHandler(t *testing.T) {
	mock := NewMockStore()

	mock.On("GetBooks").Return([]*Book{
		{
			"test book",
			"category",
			"emmanuel antonio",
			"description",
			5,
		},
	}, nil).Once()

	req, err :=http.NewRequest("GET", "", nil)
	if err != nil{
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	h := http.HandlerFunc(getBookHandler)
	h.ServeHTTP(rec, req)
	if ok := rec.Code; ok != http.StatusOK {
		t.Errorf("expected status code %v but got got %v", ok, http.StatusOK)
	}

	expected := Book{
		"test book",
		"category",
		"emmanuel antonio",
		"description",
		5,
	}
	b := []Book{}
	err = json.NewDecoder(rec.Body).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]
	if actual != expected {
		t.Errorf("expected %v but got %v", actual, expected)
	}
	mock.AssertExpectations(t)
}

func TestCreateBooksHandler(t *testing.T){
	mock := NewMockStore()

	mock.On("reateBook", &Book{
		"test book",
		"category",
		"emmanuel antonio",
		"description",
		5,
	}).Return(nil)

	form:= newCreateBookForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") // Convert to json/encoding
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	hf := http.HandlerFunc(createBookHandler)

	hf.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	mock.AssertExpectations(t)


}

func newCreateBookForm() *url.Values {
	form := url.Values{}
	form.Set("title", "no the only one")
	form.Set("category", "drama")
	form.Set("author", "Emmanuel Antoio")
	form.Set("description", "description")
	form.Set("ratings", "5")
	return &form
}