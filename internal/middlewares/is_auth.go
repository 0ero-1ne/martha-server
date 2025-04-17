package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/0ero-1ne/martha-server/internal/utils"
)

func IsAuth(jwtManager utils.JWTManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		splittedHeader := strings.Split(authHeader, " ")

		if len(splittedHeader) == 2 {
			authToken := splittedHeader[1]

			if jwtManager.VerifyToken(authToken) {
				userId, err := jwtManager.ExtractIdFromToken(authToken)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, "Invalid access token")
					ctx.Abort()
					return
				}

				ctx.Set("user_id", userId)
				ctx.Next()
				return
			}

			ctx.JSON(http.StatusUnauthorized, "Invalid access token")
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusUnauthorized, "Invalid authorization header")
		ctx.Abort()
	}
}
