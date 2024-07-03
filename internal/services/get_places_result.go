package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlacesProps struct {
	Title string `json:"title"`
	ImgURL string `json:"img_url"`
	Description string `json:"description"`
}

func GetPlaces(c *gin.Context){
    var places_info []PlacesProps
    
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
