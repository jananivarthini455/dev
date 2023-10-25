package main

import (
	"Laundary-Backend/model"
	repo "Laundary-Backend/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Database *gorm.DB

func main() {
	repo.SetUpRepo()

	router := gin.Default()

	router.GET("/health", handleHealth)
	router.POST("/signup", handleSignup)

	router.Run(":8080")
}

func handleHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "health is ok",
	})
}

func handleSignup(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		return
	}

	error := repo.Signup(user)
	if error != nil {
		c.JSON(400, gin.H{
			"error": error,
		})
		return
	}
	c.JSON(200, gin.H{
		"data": "signup for user: " + user.Username + " number: " + user.Username + " email: " + user.Email,
	})
}
