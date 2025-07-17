package routes

import (
	"log"
	"net/http"

	"github.com/Niranjini-Kathiravan/go-rest-api-v2/models"
	"github.com/Niranjini-Kathiravan/go-rest-api-v2/utils"
	"github.com/gin-gonic/gin"
)

// POST /signup
func signup(context *gin.Context) {
	var user models.User

	// Bind and validate incoming JSON
	if err := context.ShouldBindJSON(&user); err != nil {
		log.Println("Failed to bind JSON during signup:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user data.",
			"error":   err.Error(),
		})
		return
	}

	// Save user to the database
	if err := user.Save(); err != nil {
		log.Println("Error saving user:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user.",
		})
		return
	}

	// Respond with success
	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user, // You may remove this if exposing user data is not safe
	})
}

// POST /login
func login(context *gin.Context) {
	var user models.User

	// Bind incoming JSON
	if err := context.ShouldBindJSON(&user); err != nil {
		log.Println("Failed to bind JSON during login:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid login data.",
			"error":   err.Error(),
		})
		return
	}

	// Validate credentials
	if err := user.ValidateCredentials(); err != nil {
		log.Println("Authentication failed:", err)
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password.",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authenticate user.",
		})
		return
	}

	// Success
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful", "token": token})
}
