package middlewares

import (
	"net/http"

	"github.com/Niranjini-Kathiravan/go-rest-api-v2/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized."})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not Authorized"})
		return
	}

	//log.Println("POST /events called")

	context.Set("userId", userId)

	context.Next()

}
