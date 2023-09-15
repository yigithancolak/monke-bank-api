package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/yigithancolak/monke-bank-api/db/sqlc"
)

// type renewAccessTokenRequest struct {
// 	RefreshToken string `json:"refresh_token" binding:"required"`
// }

type renewAccesTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	// var req renewAccessTokenRequest

	// err := ctx.ShouldBindJSON(&req)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }

	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		cookieErr := errors.New("failed to retrive cookie")
		ctx.JSON(http.StatusBadRequest, errorResponse(cookieErr))
		return
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if session.IsBlocked {
		err := fmt.Errorf("incorrect session user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if session.UserID != refreshPayload.UserID {
		err := fmt.Errorf("mismatched session user id")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if session.RefreshToken != refreshToken {
		err := fmt.Errorf("mismatched session token")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(refreshPayload.ID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := renewAccesTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: time.Unix(accessPayload.ExpiresAt, 0),
	}

	ctx.JSON(http.StatusOK, resp)
}
