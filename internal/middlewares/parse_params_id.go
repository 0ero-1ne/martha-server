package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseParamsId(params []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, param := range params {
			id, err := strconv.ParseUint(ctx.Param(param), 10, 64)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid '%s' param value", param)})
				return
			}

			ctx.Set(param, uint(id))
		}

		ctx.Next()
	}
}
