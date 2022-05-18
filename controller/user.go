package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"main/models"
)

func (c Controllers) CreateUser(ctx *gin.Context) {
	register := models.EditingUser{}
	err := ctx.ShouldBind(&register)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "parâmetros incorretos, verifique a api",
		})
		return
	}
	if db.ExistsUsername(register.Username) {
		ctx.JSON(409, gin.H{
			"message": "Usuário já cadastrado",
		})
		return
	}
	db.SaveUser(register)
	ctx.JSON(201, gin.H{
		"data": gin.H{
			"username": register.Username,
		},
	})
}

func (c Controllers) CreateToken(ctx *gin.Context) {
	credentials := models.Credentials{}
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "parâmetros incorretos, verifique a api",
		})
		return
	}
	id, err := db.CheckCredential(credentials)
	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "dados incorretos, verifique as credenciais",
		})
		return
	}
	token := base64.StdEncoding.EncodeToString([]byte(id))
	ctx.JSON(201, gin.H{
		"data": gin.H{
			"token": token,
		},
	})
}
