// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	router.GET("/books", getBooks)
// 	router.POST("/books", postBooks)
// 	router.GET("/albums/:bookId", getBookByBookId)
// 	router.GET("/albums/:bookId", UpdateByBookId)
// 	router.DELETE("/albums/:bookId", DeleteByBookId)
// 	router.Run("localhost:8080")
// }

// // getBooks responds with the list of all books as JSON.
// func getBooks(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, books)
// }

// // postAlbums adds an album from JSON received in the request body.
// func postBooks(c *gin.Context) {
// 	var newBooks Book

// 	// newAlbum.
// 	if err := c.BindJSON(&newBooks); err != nil {
// 		return
// 	}

// 	books = append(books, newBooks)
// 	c.IndentedJSON(http.StatusCreated, newBooks)
// }

// // Book represents data about a book.
// type Book struct {
// 	BookId          string  `json:"bookId"`
// 	AuthorId        string  `json:"authorId"`
// 	PublisherId     string  `json:"publisherId"`
// 	Title           string  `json:"title"`
// 	PublicationDate string  `json:"publicationDate"`
// 	Isbn            string  `json:"isbn"`
// 	Pages           int     `json:"pages"`
// 	Genre           string  `json:"genre"`
// 	Description     string  `json:"description"`
// 	Price           float64 `json:"price"`
// 	Quantity        int     `json:"quantity"`
// }

// // books slice to seed book data.
// var books = []Book{
// 	{
// 		BookId:          "bb329a31-6b1e-4daa-87ee-71631aa05866",
// 		AuthorId:        "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
// 		PublisherId:     "2f7b19e9-b268-4440-a15b-bed8177ed607",
// 		Title:           "The Great Gatsby",
// 		PublicationDate: "1925-04-10",
// 		Isbn:            "9780743273565",
// 		Pages:           180,
// 		Genre:           "Novel",
// 		Description:     "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
// 		Price:           15.99,
// 		Quantity:        5,
// 	},
// }

// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	router.GET("/books", getBooks)
// 	router.POST("/books", postBooks)
// 	router.GET("/books/:bookId", getBookByBookId)
// 	router.PUT("/books/:bookId", updateBookByBookId)
// 	router.DELETE("/books/:bookId", deleteBookByBookId)
// 	router.Run("localhost:8080")
// }

// // Book represents data about a book.
// type Book struct {
// 	BookId          string  `json:"bookId"`
// 	AuthorId        string  `json:"authorId"`
// 	PublisherId     string  `json:"publisherId"`
// 	Title           string  `json:"title"`
// 	PublicationDate string  `json:"publicationDate"`
// 	Isbn            string  `json:"isbn"`
// 	Pages           int     `json:"pages"`
// 	Genre           string  `json:"genre"`
// 	Description     string  `json:"description"`
// 	Price           float64 `json:"price"`
// 	Quantity        int     `json:"quantity"`
// }

// // books slice to seed book data.
// var books = []Book{
// 	{
// 		BookId:          "bb329a31-6b1e-4daa-87ee-71631aa05866",
// 		AuthorId:        "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
// 		PublisherId:     "2f7b19e9-b268-4440-a15b-bed8177ed607",
// 		Title:           "The Great Gatsby",
// 		PublicationDate: "1925-04-10",
// 		Isbn:            "9780743273565",
// 		Pages:           180,
// 		Genre:           "Novel",
// 		Description:     "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
// 		Price:           15.99,
// 		Quantity:        5,
// 	},
// }

// // getBooks responds with the list of all books as JSON.
// func getBooks(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, books)
// }

// // postBooks adds a book from JSON received in the request body.
// func postBooks(c *gin.Context) {
// 	var newBook Book

// 	// Bind received JSON to newBook.
// 	if err := c.BindJSON(&newBook); err != nil {
// 		return
// 	}

// 	books = append(books, newBook)
// 	c.IndentedJSON(http.StatusCreated, newBook)
// }

// // getBookByBookId responds with a book that matches the provided bookId.
// func getBookByBookId(c *gin.Context) {
// 	bookId := c.Param("bookId")

// 	for _, book := range books {
// 		if book.BookId == bookId {
// 			c.IndentedJSON(http.StatusOK, book)
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
// }

// // updateBookByBookId updates the book that matches the provided bookId.
// func updateBookByBookId(c *gin.Context) {
// 	bookId := c.Param("bookId")
// 	var updatedBook Book

// 	if err := c.BindJSON(&updatedBook); err != nil {
// 		return
// 	}

// 	for i, book := range books {
// 		if book.BookId == bookId {
// 			books[i] = updatedBook
// 			c.IndentedJSON(http.StatusOK, updatedBook)
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
// }

// // deleteBookByBookId deletes the book that matches the provided bookId.
// func deleteBookByBookId(c *gin.Context) {
// 	bookId := c.Param("bookId")

// 	for i, book := range books {
// 		if book.BookId == bookId {
// 			books = append(books[:i], books[i+1:]...)
// 			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
// }

// package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	loadBooksFromJSON("books.json")

// 	router := gin.Default()
// 	router.GET("/books", getBooks)
// 	router.POST("/books", postBooks)
// 	router.GET("/books/:bookId", getBookByBookId)
// 	router.PUT("/books/:bookId", updateBookByBookId)
// 	router.DELETE("/books/:bookId", deleteBookByBookId)

// 	//Search book by keyword route
// 	router.GET("/books/search", searchBooks)
// 	router.Run("localhost:8080")
// }

// func loadBooksFromJSON(filename string) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	bytes, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := json.Unmarshal(bytes, &books); err != nil {
// 		panic(err)
// 	}
// }

// // Book represents data about a book.
// type Book struct {
// 	BookId          string  `json:"bookId"`
// 	AuthorId        string  `json:"authorId"`
// 	PublisherId     string  `json:"publisherId"`
// 	Title           string  `json:"title"`
// 	PublicationDate string  `json:"publicationDate"`
// 	Isbn            string  `json:"isbn"`
// 	Pages           int     `json:"pages"`
// 	Genre           string  `json:"genre"`
// 	Description     string  `json:"description"`
// 	Price           float64 `json:"price"`
// 	Quantity        int     `json:"quantity"`
// }

// // books slice to seed book data.
// var books []Book

// // getBooks
// func getBooks(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, books)
// }

// // postBooks adds a book
// func postBooks(c *gin.Context) {
// 	var newBook Book

// 	// Bind received JSON to newBook.
// 	if err := c.BindJSON(&newBook); err != nil {
// 		return
// 	}

// 	books = append(books, newBook)
// 	c.IndentedJSON(http.StatusCreated, newBook)
// }

// // getBookByBookId function
// func getBookByBookId(c *gin.Context) {
// 	bookId := c.Param("bookId")

// 	for _, book := range books {
// 		if book.BookId == bookId {
// 			c.IndentedJSON(http.StatusOK, book)
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
// }

// // updateBookByBookId function
// func updateBookByBookId(c *gin.Context) {
// 	bookId := c.Param("bookId")
// 	var updatedBook Book

// 	if err := c.BindJSON(&updatedBook); err != nil {
// 		return
// 	}

// 	for i, book := range books {
// 		if book.BookId == bookId {
// 			books[i] = updatedBook
// 			c.IndentedJSON(http.StatusOK, updatedBook)
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
// }

// // deleteBookByBookId deletes the book that matches the provided bookId.
// func deleteBookByBookId(c *gin.Context) {
// 	bookId := c.Param("bookId")

// 	for i, book := range books {
// 		if book.BookId == bookId {
// 			books = append(books[:i], books[i+1:]...)
// 			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
// }

// func searchBooks(c *gin.Context) {
// 	query := c.Query("q")
// 	if query == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Query parameter 'q' is required"})
// 		return
// 	}

// 	results := performSearch(query)
// 	c.IndentedJSON(http.StatusOK, results)
// }

// func performSearch(query string) []Book {
// 	query = strings.ToLower(query)
// 	results := make([]Book, 0)

// 	numWorkers := 4
// 	bookChannel := make(chan Book, len(books))
// 	resultChannel := make(chan Book, len(books))

// 	// Start worker goroutines
// 	for i := 0; i < numWorkers; i++ {
// 		go searchWorker(query, bookChannel, resultChannel)
// 	}

// 	// Send books to the bookChannel
// 	go func() {
// 		for _, book := range books {
// 			bookChannel <- book
// 		}
// 		close(bookChannel)
// 	}()

// 	// Collect results from resultChannel
// 	for i := 0; i < len(books); i++ {
// 		result, ok := <-resultChannel
// 		if ok {
// 			results = append(results, result)
// 		}
// 	}

// 	return results
// }

// func searchWorker(query string, books <-chan Book, results chan<- Book) {
// 	for book := range books {
// 		if strings.Contains(strings.ToLower(book.Title), query) || strings.Contains(strings.ToLower(book.Description), query) {
// 			results <- book
// 		}
// 	}
// }

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	loadBooksFromJSON("books.json")

	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.GET("/books/:bookId", getBookByBookId)
	router.PUT("/books/:bookId", updateBookByBookId)
	router.DELETE("/books/:bookId", deleteBookByBookId)

	// Search book by keyword route
	router.GET("/books/search", searchBooks)
	router.Run("localhost:8080")
}

func loadBooksFromJSON(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(bytes, &books); err != nil {
		panic(err)
	}
}

// Book represents data about a book.
type Book struct {
	BookId          string  `json:"bookId"`
	AuthorId        string  `json:"authorId"`
	PublisherId     string  `json:"publisherId"`
	Title           string  `json:"title"`
	PublicationDate string  `json:"publicationDate"`
	Isbn            string  `json:"isbn"`
	Pages           int     `json:"pages"`
	Genre           string  `json:"genre"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	Quantity        int     `json:"quantity"`
}

var books []Book

// Retrive All the Books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// Add new Book Details
func postBooks(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// Get book by BookID
func getBookByBookId(c *gin.Context) {
	bookId := c.Param("bookId")

	for _, book := range books {
		if book.BookId == bookId {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// update book by BookID
func updateBookByBookId(c *gin.Context) {
	bookId := c.Param("bookId")
	var updatedBook Book

	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}

	for i, book := range books {
		if book.BookId == bookId {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// Delete book by BookID
func deleteBookByBookId(c *gin.Context) {
	bookId := c.Param("bookId")

	for i, book := range books {
		if book.BookId == bookId {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func searchBooks(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Query parameter 'q' is required"})
		return
	}

	results := performSearch(query)
	c.IndentedJSON(http.StatusOK, results)
}

func performSearch(query string) []Book {
	query = strings.ToLower(query)
	results := make([]Book, 0)

	numWorkers := 4
	bookChannel := make(chan Book, len(books))
	resultChannel := make(chan Book)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go searchWorker(query, bookChannel, resultChannel, &wg)
	}

	// Send books to the bookChannel
	go func() {
		for _, book := range books {
			bookChannel <- book
		}
		close(bookChannel)
	}()

	// Close resultChannel once all workers are done
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Collect results from resultChannel
	for result := range resultChannel {
		results = append(results, result)
	}

	return results
}

func searchWorker(query string, books <-chan Book, results chan<- Book, wg *sync.WaitGroup) {
	defer wg.Done()
	for book := range books {
		if strings.Contains(strings.ToLower(book.Title), query) || strings.Contains(strings.ToLower(book.Description), query) {
			results <- book
		}
	}
}
