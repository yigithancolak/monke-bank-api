package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		ID:           uuid.New(),
		Owner:        uuid.New(), //will change
		CurrencyCode: req.Currency,
		Balance:      0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		// errCode := db.ErrorCode(err)

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}
