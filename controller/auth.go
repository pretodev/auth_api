package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthHeader struct {
	Authorization string `header:"Authorization" binding:"required"`
}

func (a *AuthHeader) decodeToken() string {
	data, _ := base64.StdEncoding.DecodeString(a.Authorization)
	return fmt.Sprintf("%s", data)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := AuthHeader{}
		err := ctx.ShouldBindHeader(&header)

		if err == nil {
			if db.ExistsId(header.decodeToken()) {
				fmt.Println(header.decodeToken())
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Usuário não autorizado.",
		})
	}
}
