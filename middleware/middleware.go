package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/techschool/simplebank/token"
	"strings"

	"net/http"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		authorizationType := strings.ToLower(fields[0])
		// compare auth types react accordingly
		if authorizationType != authorizationTypeBearer {
			err := errors.New("authorization not supported by server")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		//save authorization payload to context
		ctx.Set(authorizationPayloadKey, payload)
		// forward request to next handler
		ctx.Next()
	}
}
