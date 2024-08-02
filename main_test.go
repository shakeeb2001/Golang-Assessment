package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.GET("/books/:bookId", getBookByBookId)
	router.PUT("/books/:bookId", updateBookByBookId)
	router.DELETE("/books/:bookId", deleteBookByBookId)
	router.GET("/books/search", searchBooks)
	return router
}

func loadTestData() {
	books = []Book{
		{"bb329a31-6b1e-4daa-87ee-71631aa05866", "e0d91f68-a183-477d-8aa4-1f44ccc78a70", "2f7b19e9-b268-4440-a15b-bed8177ed607", "The Great Gatsby", "1925-04-10", "9780743273565", 180, "Novel", "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.", 15.99, 5},
		{"c0a79e07-8c6a-41a3-93d5-b0a8b03d0f0d", "4d1e4477-d658-4f93-9a58-82fd56beef6f", "6fa72e20-8dd2-4b69-b08f-50603c38d122", "To Kill a Mockingbird", "1960-07-11", "9780061120084", 324, "Fiction", "A novel about the serious issues of rape and racial inequality.", 18.99, 7},
		{"91e79a02-4379-4d29-8bff-8e4d0dcf0ecb", "c3968e91-baf7-47b6-85a2-2f5f2ef6e7f4", "3c1d60f4-ec52-4979-9c4b-671b70f7a56d", "1984", "1949-06-08", "9780451524935", 328, "Dystopian", "A novel that tells the story of a dystopian society under totalitarian rule.", 14.99, 10},
		{"8e0d94a7-cc91-4ab1-96d1-4d0d6f81d0e3", "7b6f0b3b-0b69-4311-82db-b56e3b88a8a3", "7c8a792d-3f0f-4d2e-b828-9d801c0e2b94", "Pride and Prejudice", "1813-01-28", "9781503290563", 279, "Romance", "A romantic novel that critiques the British landed gentry at the end of the 18th century.", 12.99, 3},
		{"f5b1fbc7-9934-4c3e-826f-f5f8056d8f23", "94c61f62-28a5-4f11-9916-3d2e9fbc1e7e", "8a072c63-6d72-4b7f-b7de-3c72a169d67f", "The Catcher in the Rye", "1951-07-16", "9780316769488", 214, "Fiction", "A novel about the events in the life of a young boy named Holden Caulfield.", 10.99, 8},
	}
}

// Test the Getbooks with pagination.
func TestGetBooks(t *testing.T) {

	loadTestData()
	router := setupRouter()

	req, err := http.NewRequest("GET", "/books?limit=2&offset=0", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `[{"bookId":"bb329a31-6b1e-4daa-87ee-71631aa05866","authorId":"e0d91f68-a183-477d-8aa4-1f44ccc78a70","publisherId":"2f7b19e9-b268-4440-a15b-bed8177ed607","title":"The Great Gatsby","publicationDate":"1925-04-10","isbn":"9780743273565","pages":180,"genre":"Novel","description":"Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.","price":15.99,"quantity":5},{"bookId":"c0a79e07-8c6a-41a3-93d5-b0a8b03d0f0d","authorId":"4d1e4477-d658-4f93-9a58-82fd56beef6f","publisherId":"6fa72e20-8dd2-4b69-b08f-50603c38d122","title":"To Kill a Mockingbird","publicationDate":"1960-07-11","isbn":"9780061120084","pages":324,"genre":"Fiction","description":"A novel about the serious issues of rape and racial inequality.","price":18.99,"quantity":7}]`
	assert.JSONEq(t, expected, w.Body.String())
}

// TestGetBookByBookId
func TestGetBookByBookId(t *testing.T) {
	loadTestData()
	router := setupRouter()

	req, err := http.NewRequest("GET", "/books/bb329a31-6b1e-4daa-87ee-71631aa05866", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `{"bookId":"bb329a31-6b1e-4daa-87ee-71631aa05866","authorId":"e0d91f68-a183-477d-8aa4-1f44ccc78a70","publisherId":"2f7b19e9-b268-4440-a15b-bed8177ed607","title":"The Great Gatsby","publicationDate":"1925-04-10","isbn":"9780743273565","pages":180,"genre":"Novel","description":"Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.","price":15.99,"quantity":5}`
	assert.JSONEq(t, expected, w.Body.String())
}

// TestPostBooks tests
func TestPostBooks(t *testing.T) {
	loadTestData()
	router := setupRouter()

	newBook := Book{
		BookId:          "6c8d8e89-4e1f-4ef8-8f9d-2f48b1c5ef56",
		AuthorId:        "12345678-1234-1234-1234-123456789012",
		PublisherId:     "87654321-4321-4321-4321-210987654321",
		Title:           "New Book",
		PublicationDate: "2023-01-01",
		Isbn:            "9781234567890",
		Pages:           100,
		Genre:           "Fiction",
		Description:     "A new book for testing.",
		Price:           9.99,
		Quantity:        1,
	}

	jsonValue, _ := json.Marshal(newBook)
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	expected := `{"bookId":"6c8d8e89-4e1f-4ef8-8f9d-2f48b1c5ef56","authorId":"12345678-1234-1234-1234-123456789012","publisherId":"87654321-4321-4321-4321-210987654321","title":"New Book","publicationDate":"2023-01-01","isbn":"9781234567890","pages":100,"genre":"Fiction","description":"A new book for testing.","price":9.99,"quantity":1}`
	assert.JSONEq(t, expected, w.Body.String())
}

// TestUpdateBookByBookId test
func TestUpdateBookByBookId(t *testing.T) {
	loadTestData()
	router := setupRouter()

	updatedBook := Book{
		BookId:          "bb329a31-6b1e-4daa-87ee-71631aa05866",
		AuthorId:        "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
		PublisherId:     "2f7b19e9-b268-4440-a15b-bed8177ed607",
		Title:           "The Great Gatsby (Updated)",
		PublicationDate: "1925-04-10",
		Isbn:            "9780743273565",
		Pages:           180,
		Genre:           "Novel",
		Description:     "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
		Price:           15.99,
		Quantity:        5,
	}

	jsonValue, _ := json.Marshal(updatedBook)
	req, err := http.NewRequest("PUT", "/books/bb329a31-6b1e-4daa-87ee-71631aa05866", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `{"bookId":"bb329a31-6b1e-4daa-87ee-71631aa05866","authorId":"e0d91f68-a183-477d-8aa4-1f44ccc78a70","publisherId":"2f7b19e9-b268-4440-a15b-bed8177ed607","title":"The Great Gatsby (Updated)","publicationDate":"1925-04-10","isbn":"9780743273565","pages":180,"genre":"Novel","description":"Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.","price":15.99,"quantity":5}`
	assert.JSONEq(t, expected, w.Body.String())
}

// TestDeleteBookByBookId test
func TestDeleteBookByBookId(t *testing.T) {
	loadTestData()
	router := setupRouter()

	req, err := http.NewRequest("DELETE", "/books/bb329a31-6b1e-4daa-87ee-71631aa05866", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `{"message":"book deleted"}`
	assert.JSONEq(t, expected, w.Body.String())
}

// TestSearchBooks test
func TestSearchBooks(t *testing.T) {
	loadTestData()
	router := setupRouter()

	req, err := http.NewRequest("GET", "/books/search?q=great", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `[{"bookId":"bb329a31-6b1e-4daa-87ee-71631aa05866","authorId":"e0d91f68-a183-477d-8aa4-1f44ccc78a70","publisherId":"2f7b19e9-b268-4440-a15b-bed8177ed607","title":"The Great Gatsby","publicationDate":"1925-04-10","isbn":"9780743273565","pages":180,"genre":"Novel","description":"Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.","price":15.99,"quantity":5}]`
	assert.JSONEq(t, expected, w.Body.String())
}
