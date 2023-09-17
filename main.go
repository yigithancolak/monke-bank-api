package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yigithancolak/monke-bank-api/api"
	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
	"github.com/yigithancolak/monke-bank-api/grpcapi"
	"github.com/yigithancolak/monke-bank-api/pb"
	"github.com/yigithancolak/monke-bank-api/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbpool, err := pgxpool.New(context.Background(), config.PostgresURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	store := db.NewStore(dbpool)

	// runGinServer(config, store)
	runGrpcServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(":8888")
	if err != nil {
		log.Fatal("cannot start server")
	}
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := grpcapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMonkeBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("Starting grpc server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server")
	}

}
