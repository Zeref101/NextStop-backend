package main

import (
	"net/http"
	"time"

	"github.com/Zeref101/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func greetings(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{
        "Hello": "visitor",
    })
    return
}

func main() {
	r := gin.Default()

	if err := r.SetTrustedProxies(nil); err!=nil{
		panic(err)
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET","POST","PUT","DELETE"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	routes.PlacesInfoRoutes(r)

	r.POST("/", greetings)

	r.Run(":7000")
}