package routes

import (
	"net/http"

	"github.com/Zeref101/internal/services/auth"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine){
	userGroup := router.Group("/auth")
	{
		userGroup.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":"success",
			})
			
		})
		userGroup.POST("/sign_up", auth.Sign_up)
		userGroup.POST("/sign_in", auth.Sign_in)
	}
	
}