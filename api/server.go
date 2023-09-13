package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
	"github.com/yigithancolak/monke-bank-api/token"
	"github.com/yigithancolak/monke-bank-api/util"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	config     util.Config
	tokenMaker token.Maker
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/auth/register", server.createUser)
	router.POST("/auth/login", server.loginUser)

	protectedRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	protectedRoutes.POST("/accounts", server.createAccount)

	protectedRoutes.POST("/transfers", server.createTransfer)

	server.router = router
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

	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
