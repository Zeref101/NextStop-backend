package services

import (
	"fmt"
	"net/http"

	"github.com/Zeref101/internal/types"
	"github.com/gin-gonic/gin"
)


func GetPlaces(c *gin.Context){
    var places_info []types.PlacesProps
    
    if err := c.ShouldBindJSON(&places_info); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    fmt.Println(places_info)
    c.JSON(http.StatusOK, gin.H{
        "places_info": places_info,
    })
}
func GreetingPlaces(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{
        "message":"/places",
    })
}
