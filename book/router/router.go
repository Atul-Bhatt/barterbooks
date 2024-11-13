package router

import (
	"net/http"
	"strconv"

	"book/middleware"
	"book/model"
	"book/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var repo *repository.BookRepository

func SetupRouter(db *sqlx.DB) *gin.Engine {
	r := gin.New()
	r.Use(middleware.LoggingMiddleware())

	log.SetFormatter(&log.JSONFormatter{})

	repo = repository.NewBookRepository(db)

	r.GET("/health_check", healthCheck)
	r.GET("/books", getBooks)
	r.POST("/books/", createBook)
	r.GET("/books/:id", getBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)
	return r
}

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "hello from bookbarter!")
}

// get all books
func getBooks(c *gin.Context) {
	books, err := repo.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Message": "Error getting books."})
	} else {
		c.JSON(http.StatusOK, books)
	}
}

// create a book
func createBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	validationError := validate.Struct(book)
	if validationError != nil {
		// get the correlation id of the request
		correlationID, _ := c.Get("correlationID")
		log.WithFields(log.Fields{
			"correlationID": correlationID,
			"Method":        c.Request.Method,
			"Path":          c.Request.URL.Path,
			"Status":        c.Writer.Status,
		}).Error(validationError.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
		return
	}

	repo.Create(book)
	c.JSON(http.StatusCreated, book)
}

// get book by ISBN
func getBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := repo.GetBook(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Message": err.Error()})
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func updateBook(c *gin.Context) {
	var book model.Book
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := repo.GetBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	if err = c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = repo.UpdateBook(book, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := repo.GetBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	repo.DeleteBook(id)
	c.JSON(http.StatusNoContent, gin.H{})
}
