package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/yigithancolak/monke-bank-api/api"
	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
)

func main() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	store := db.New(dbpool)

	runGinServer(*store)
}

func runGinServer(store db.Queries) {
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start("localhost:8888")
	if err != nil {
		log.Fatal("cannot start server")
	}
}
