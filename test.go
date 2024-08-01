package main

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// // TestGetBooks tests the getBooks endpoint.
// func TestGetBooks(t *testing.T) {
// 	// Set up the router and load sample data
// 	loadBooksFromJSON("books.json")
// 	router := gin.Default()
// 	router.GET("/books", getBooks)

// 	// Create a request to the /books endpoint
// 	req, err := http.NewRequest("GET", "/books", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v\n", err)
// 	}

// 	// Create a ResponseRecorder to record the response
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Check the status code
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// Check the response body
// 	expected := `[{"bookId":"bb329a31-6b1e-4daa-87ee-71631aa05866","authorId":"e0d91f68-a183-477d-8aa4-1f44ccc78a70","publisherId":"2f7b19e9-b268-4440-a15b-bed8177ed607","title":"The Great Gatsby","publicationDate":"1925-04-10","isbn":"9780743273565","pages":180,"genre":"Novel","description":"Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.","price":15.99,"quantity":5},{"bookId":"c0a79e07-8c6a-41a3-93d5-b0a8b03d0f0d","authorId":"4d1e4477-d658-4f93-9a58-82fd56beef6f","publisherId":"6fa72e20-8dd2-4b69-b08f-50603c38d122","title":"To Kill a Mockingbird","publicationDate":"1960-07-11","isbn":"9780061120084","pages":324,"genre":"Fiction","description":"A novel about the serious issues of rape and racial inequality.","price":18.99,"quantity":7},{"bookId":"91e79a02-4379-4d29-8bff-8e4d0dcf0ecb","authorId":"c3968e91-baf7-47b6-85a2-2f5f2ef6e7f4","publisherId":"3c1d60f4-ec52-4979-9c4b-671b70f7a56d","title":"1984","publicationDate":"1949-06-08","isbn":"9780451524935","pages":328,"genre":"Dystopian","description":"A novel that tells the story of a dystopian society under totalitarian rule.","price":14.99,"quantity":10},{"bookId":"8e0d94a7-cc91-4ab1-96d1-4d0d6f81d0e3","authorId":"7b6f0b3b-0b69-4311-82db-b56e3b88a8a3","publisherId":"7c8a792d-3f0f-4d2e-b828-9d801c0e2b94","title":"Pride and Prejudice","publicationDate":"1813-01-28","isbn":"9781503290563","pages":279,"genre":"Romance","description":"A romantic novel that critiques the British landed gentry at the end of the 18th century.","price":12.99,"quantity":3},{"bookId":"f5b1fbc7-9934-4c3e-826f-f5f8056d8f23","authorId":"94c61f62-28a5-4f11-9916-3d2e9fbc1e7e","publisherId":"8a072c63-6d72-4b7f-b7de-3c72a169d67f","title":"The Catcher in the Rye","publicationDate":"1951-07-16","isbn":"9780316769488","pages":214,"genre":"Fiction","description":"A novel about the events in the life of a young boy named Holden Caulfield.","price":10.99,"quantity":8}]`
// 	assert.JSONEq(t, expected, w.Body.String())
// }
