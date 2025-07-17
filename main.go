package main

import (
	"log"

	"github.com/Niranjini-Kathiravan/go-rest-api-v2/db"
	"github.com/Niranjini-Kathiravan/go-rest-api-v2/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	log.Println("Server starting on http://localhost:8080")
	server.Run(":8080")
}
