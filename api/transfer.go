package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
	"github.com/yigithancolak/monke-bank-api/token"
)

const (
	errNotMatchedCurrencies = "currencies of accounts are not matched, choose two accounts have same currencies"
	errNotOwner             = "account which will make the transfer is not belong to the user"
)

var (
	ErrCurrencyNotMatched = errors.New(errNotMatchedCurrencies)
	ErrNotAccountOwner    = errors.New(errNotOwner)
)

type createTransferRequest struct {
	FromAccountID uuid.UUID `json:"from_account_id" binding:"required"`
	ToAccountID   uuid.UUID `json:"to_account_id" binding:"required"`
	Amount        int32     `json:"amount" binding:"required,gte=1"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccount, err := server.checkCurrencies(ctx, req.FromAccountID, req.ToAccountID)

	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		if err == ErrCurrencyNotMatched {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if fromAccount.Owner != authPayload.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(ErrNotAccountOwner))
		return
	}

	transferResult, err := server.store.TransferTx(ctx, db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transferResult)

}

func (server *Server) checkCurrencies(ctx context.Context, fromAccountID uuid.UUID, toAccountID uuid.UUID) (fromAccount db.Account, err error) {
	fromAccount, err = server.store.GetAccountById(ctx, fromAccountID)
	if err != nil {
		return
	}

	toAccount, err := server.store.GetAccountById(ctx, toAccountID)
	if err != nil {
		return
	}

	if fromAccount.CurrencyCode != toAccount.CurrencyCode {
		err = ErrCurrencyNotMatched
		return
	}

	return
}
