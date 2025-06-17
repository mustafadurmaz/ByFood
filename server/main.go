package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"byfood-task/handlers"
	"byfood-task/storage"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	log.Println("Database connection successful.")

	dbStorage := storage.NewPostgresStorage(dbpool)

	dbStorage.Migrate()

	bookHandler := handlers.NewBookHandler(dbStorage)

	urlHandler := handlers.NewURLHandler()

	router := gin.Default()

	api := router.Group("/books")
	{
		api.POST("/", bookHandler.CreateBook)
		api.GET("/", bookHandler.GetBooks)
		api.GET("/:id", bookHandler.GetBookByID)
		api.PUT("/:id", bookHandler.UpdateBook)
		api.DELETE("/:id", bookHandler.DeleteBook)
	}

	router.POST("/process-url", urlHandler.ProcessURL)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		MaxAge: 300,
	})
	handler := c.Handler(router)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
