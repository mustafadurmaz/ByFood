package storage

import (
	"context"
	"log"

	"byfood-task/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	CreateBook(book *models.Book) error
	GetBooks() ([]models.Book, error)
	GetBookByID(id int) (models.Book, error)
	UpdateBook(id int, book models.Book) error
	DeleteBook(id int) error
	Migrate()
}

type PostgresStorage struct {
	db *pgxpool.Pool
}

func NewPostgresStorage(dbpool *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{db: dbpool}
}

func (s *PostgresStorage) Migrate() {
	query := `
    CREATE TABLE IF NOT EXISTS books (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        author VARCHAR(255) NOT NULL,
        year INT,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
    );`
	_, err := s.db.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Unable to create books table: %v\n", err)
	}
	log.Println("Books table migrated successfully.")
}

func (s *PostgresStorage) CreateBook(book *models.Book) error {
	query := `INSERT INTO books (title, author, year) VALUES ($1, $2, $3) RETURNING id, created_at`
	return s.db.QueryRow(context.Background(), query, book.Title, book.Author, book.Year).Scan(&book.ID, &book.CreatedAt)
}

func (s *PostgresStorage) GetBooks() ([]models.Book, error) {
	query := `SELECT id, title, author, year, created_at FROM books`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.CreatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *PostgresStorage) GetBookByID(id int) (models.Book, error) {
	query := `SELECT id, title, author, year, created_at FROM books WHERE id = $1`
	var book models.Book
	err := s.db.QueryRow(context.Background(), query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.CreatedAt)
	return book, err
}

func (s *PostgresStorage) UpdateBook(id int, book models.Book) error {
	query := `UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4`
	_, err := s.db.Exec(context.Background(), query, book.Title, book.Author, book.Year, id)
	return err
}

func (s *PostgresStorage) DeleteBook(id int) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := s.db.Exec(context.Background(), query, id)
	return err
}