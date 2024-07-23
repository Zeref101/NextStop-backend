package auth

import (
	"net/http"

	"github.com/Zeref101/internal/db"
	"github.com/Zeref101/internal/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Sign_up(c *gin.Context){

	var details types.User

	if err := c.BindJSON(&details); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dbInstance, err := db.ConnectToDB()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	defer db.CloseDB(dbInstance)

	var user types.User

	result := dbInstance.Where("email = ?", details.Email).First(&user)

	if result.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(details.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}
	details.Password = string(hashedPassword)

	result = dbInstance.Create(&details)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Sign_in(ctx *gin.Context){

	var user types.Sign_in_props
	var fetchedUser types.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	dbInsance, err := db.ConnectToDB()
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	defer db.CloseDB(dbInsance)

	res := dbInsance.Select([]string{"username, email"}).Where("email = ?", user.Email).First(&fetchedUser)

	if res.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte (fetchedUser.Password), []byte (user.Password)); err !=nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":"Invalid email or password",
		})
	}


	ctx.JSON(http.StatusOK, gin.H{
		"message":"Sign in successful",
	})
}