package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	router := gin.Default()
	router.GET("/accounts", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "hello world"})

	})

	router.Run("localhost:8888")
}
