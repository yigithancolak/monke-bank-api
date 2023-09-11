package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yigithancolak/monke-bank-api/token"
)

const (
	authorizationHeaderKey  = "authorization"
	bearerTypeAuthorization = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

var (
	errEmptyAuthorizationHeader   = errors.New("empty authorization header in request")
	errInvalidAuthorizationFormat = errors.New("invalid authorization header format")
	errWrongAuthorizationType     = errors.New("not supported authorization type")
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errEmptyAuthorizationHeader))
			return
		}
		headerParts := strings.Fields(authorizationHeader)
		if len(headerParts) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errInvalidAuthorizationFormat))
			return
		}

		authorizationType := strings.ToLower(headerParts[0])
		if authorizationType != bearerTypeAuthorization {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errWrongAuthorizationType))
			return
		}

		accessToken := headerParts[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()

	}
}
