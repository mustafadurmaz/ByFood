package handlers

import (
	"log"
	"net/http"
	"strconv"

	"byfood-task/models"
	"byfood-task/storage"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	storage storage.Storage
}

func NewBookHandler(s storage.Storage) *BookHandler {
	return &BookHandler{storage: s}
}

// @Summary Create a new book
// @Description Add a new book to the database
// @Tags books
// @Accept  json
// @Produce  json
// @Param   book  body   models.Book  true  "Book to add"
// @Success 201 {object} models.Book
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		log.Printf("ERROR: Invalid input for CreateBook: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.storage.CreateBook(&book); err != nil {
		log.Printf("ERROR: Failed to create book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	log.Printf("INFO: Book created successfully with ID: %d", book.ID)
	c.JSON(http.StatusCreated, book)
}

// @Summary Get all books
// @Description Retrieve all books from the database
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Failure 500 {object} models.ErrorResponse
// @Router /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.storage.GetBooks()
	if err != nil {
		log.Printf("ERROR: Failed to get books: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	log.Println("INFO: Retrieved all books successfully.")
	c.JSON(http.StatusOK, books)
}

// @Summary Get a book by ID
// @Description Retrieve a specific book using its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /books/{id} [get]
func (h *BookHandler) GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("ERROR: Invalid book ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := h.storage.GetBookByID(id)
	if err != nil {
		log.Printf("ERROR: Book with ID %d not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	log.Printf("INFO: Retrieved book with ID %d successfully.", id)
	c.JSON(http.StatusOK, book)
}

// @Summary Update a book
// @Description Update a book's information by ID
// @Tags books
// @Accept  json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Updated book"
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("ERROR: Invalid book ID for update: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		log.Printf("ERROR: Invalid input for UpdateBook: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.storage.UpdateBook(id, book); err != nil {
		log.Printf("ERROR: Failed to update book with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	log.Printf("INFO: Book with ID %d updated successfully.", id)
	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// @Summary Delete a book
// @Description Delete a book by ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("ERROR: Invalid book ID for delete: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := h.storage.DeleteBook(id); err != nil {
		log.Printf("ERROR: Failed to delete book with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	log.Printf("INFO: Book with ID %d deleted successfully.", id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}