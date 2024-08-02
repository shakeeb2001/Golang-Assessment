package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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

// Retrive All the Books with pagination
func getBooks(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid limit"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid offset"})
		return
	}

	start := offset
	end := offset + limit

	if start > len(books) {
		start = len(books)
	}
	if end > len(books) {
		end = len(books)
	}

	paginatedBooks := books[start:end]
	c.IndentedJSON(http.StatusOK, paginatedBooks)
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

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

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
