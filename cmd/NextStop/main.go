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

	//? ROUTES
	routes.PlacesInfoRoutes(r)
	routes.UserRoutes(r)

	r.GET("/", greetings)
	// r.GET("/test-db", func(ctx *gin.Context){
	// 	db.ConnectToDB()
	// 	defer db.CloseDB()

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "Database connection successfully tested",
	// 	})
	// })

	r.Run(":7000")
}