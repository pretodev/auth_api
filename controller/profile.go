package controller

import (
	"github.com/gin-gonic/gin"
	"main/models"
)

func (c Controllers) GetProfile(ctx *gin.Context) {
	header := AuthHeader{}
	err := ctx.ShouldBindHeader(&header)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "parâmetros incorretos, verifique a api",
		})
		return
	}
	profile, err := db.GetProfile(header.decodeToken())
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "perfil não encontrado",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": gin.H{
			"profile": gin.H{
				"name":     profile.FullName,
				"email":    profile.Email,
				"whatsapp": profile.Whatsapp,
				"github":   profile.Github,
				"linkedin": profile.Linkedin,
			},
		},
	})
}

func (c Controllers) UpdateProfile(ctx *gin.Context) {
	header := AuthHeader{}
	err := ctx.ShouldBindHeader(&header)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "parâmetros incorretos, verifique a api",
		})
		return
	}
	profile := models.Profile{}
	err = ctx.ShouldBind(&profile)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "parâmetros incorretos, verifique a api",
		})
		return
	}
	db.SaveProfile(header.decodeToken(), profile)
	ctx.JSON(200, gin.H{
		"data": gin.H{
			"message": "perfil atualizado",
		},
	})
}
