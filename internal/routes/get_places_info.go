package routes

import (
	"github.com/Zeref101/internal/services"
	"github.com/gin-gonic/gin"
)

func PlacesInfoRoutes(router *gin.Engine){
	placesInfo := router.Group("/placesInfo")
	{
		placesInfo.GET("/",services.GreetingPlaces)
		placesInfo.POST("/place", services.GetPlaces)
	}
}