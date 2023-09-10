package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
	"github.com/yigithancolak/monke-bank-api/util"
)

type createUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
}

type userResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}

func createUserResponse(user db.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		ID:       uuid.New(),
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		// errCode := db.ErrorCode(err)

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	userResponse := createUserResponse(user)

	ctx.JSON(http.StatusOK, userResponse)
}
