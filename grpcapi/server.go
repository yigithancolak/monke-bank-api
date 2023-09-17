package grpcapi

import (
	"fmt"

	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
	"github.com/yigithancolak/monke-bank-api/pb"
	"github.com/yigithancolak/monke-bank-api/token"
	"github.com/yigithancolak/monke-bank-api/util"
)

type Server struct {
	pb.UnimplementedMonkeBankServer
	store      db.Store
	config     util.Config
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)

	if err != nil {
		return nil, fmt.Errorf("err while creating tokenMaker: %v", err)
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
