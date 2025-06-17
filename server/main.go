package main

import (
	"context"
	"log"
	"os"

	"byfood-task/handlers"
	"byfood-task/storage"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
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

	router := gin.Default()

	api := router.Group("/books")
	{
		api.POST("/", bookHandler.CreateBook)
		api.GET("/", bookHandler.GetBooks)
		api.GET("/:id", bookHandler.GetBookByID)
		api.PUT("/:id", bookHandler.UpdateBook)
		api.DELETE("/:id", bookHandler.DeleteBook)
	}

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
