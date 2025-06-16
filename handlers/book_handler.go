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

// GET /books
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

// GET /books/:id
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

// PUT /books/:id
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

// DELETE /books/:id
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