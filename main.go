package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
	"github.com/replit/database-go"
)

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

type EditingUser struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
  Password string `json:"password" binding:"required"`
	Ballance int    `json:"ballance" binding:"required"`
}

type Auth struct {
  Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthHeader struct {
	Authorization string `header:"Authorization" binding:"required"`
}

func (a *AuthHeader) decodeToken() string {
  data, _ := base64.StdEncoding.DecodeString(a.Authorization)
  return fmt.Sprintf("%s", data)
}

func exists(email string) bool {
	_, err := database.Get(email)
	return err == nil
}

func save(user EditingUser) {
	rand.Seed(time.Now().UnixNano())
	ballance := rand.Intn(500000) + 100
	data := fmt.Sprintf("%s,%s,%s,%d", user.Name, user.Email, user.Password, ballance)
	database.Set(user.Email, data)
}

func getFromEmail(email string) User {
	data, _ := database.Get(email)
	split := strings.Split(data, ",")
  fmt.Println(split)
	ballance, _ := strconv.Atoi(split[3])
	return User{
		split[0],
		split[1],
    split[2],
		ballance,
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := AuthHeader{}
		err := ctx.ShouldBindHeader(&header)
    
    fmt.Println(err)
		if err == nil {
			if exists(header.decodeToken()) {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "Usuário não autorizado.",
		})
	}
}

func main() {
	server := gin.Default()
  server.Use(cors.Default())
	server.GET("/users/infos", AuthMiddleware(), func(ctx *gin.Context) {
		header := AuthHeader{}
    ctx.ShouldBindHeader(&header)
    user := getFromEmail(header.decodeToken())
		ctx.JSON(200, gin.H{
			"data": gin.H{
				"user": gin.H{
					"name":  user.Name,
					"email": user.Name,
				},
				"bill": gin.H{
					"ballance": user.Ballance,
				},
			},
		})
	})
  
	server.POST("/users", func(ctx *gin.Context) {
		register := EditingUser{}
		err := ctx.ShouldBind(&register)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "parâmetros incorretos, verifique a api",
			})
			return
		}
		if exists(register.Email) {
			ctx.JSON(400, gin.H{
				"message": "Usuário já cadastrado",
			})
			return
		}
		save(register)
		ctx.JSON(200, gin.H{
			"data": gin.H{
				"name":  register.Name,
				"email": register.Email,
			},
		})
	})
  
	server.POST("/login", func(ctx *gin.Context) {
    auth := Auth{}
		err := ctx.ShouldBind(&auth)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "parâmetros incorretos, verifique a api",
			})
			return
		}
    user := getFromEmail(auth.Email)
    if(user.Password != auth.Password) {
      ctx.JSON(400, gin.H{
				"message": "dados incorretos, verifique as credenciais",
			})
			return
    }
    token := base64.StdEncoding.EncodeToString([]byte(user.Email))
		ctx.JSON(200, gin.H{
			"data": gin.H{
				"token": token,
			},
		})
	})
	server.Run()
}
