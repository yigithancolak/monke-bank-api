package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
	"github.com/yigithancolak/monke-bank-api/util"
)

type Server struct {
	store  db.Queries
	router *gin.Engine
	config util.Config
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/auth/register", server.createUser)

	router.POST("/accounts", server.createAccount)

	server.router = router
}

func NewServer(config util.Config, store db.Queries) (*Server, error) {

	server := &Server{
		store:  store,
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
