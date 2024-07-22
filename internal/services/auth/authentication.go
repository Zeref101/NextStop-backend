package auth

import (
	"fmt"
	"net/http"

	"github.com/Zeref101/internal/types"
	"github.com/gin-gonic/gin"
)

func Sign_up(c *gin.Context){
	var details types.UserSignUpProps
	fmt.Println("oiaefjo")
	if err := c.BindJSON(&details); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("details: ", details)
}