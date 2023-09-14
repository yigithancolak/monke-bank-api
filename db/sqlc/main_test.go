package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yigithancolak/monke-bank-api/util"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbpool, err := pgxpool.New(context.Background(), config.PostgresURL)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
		os.Exit(1)
	}

	testStore = NewStore(dbpool)
	os.Exit(m.Run())
}
