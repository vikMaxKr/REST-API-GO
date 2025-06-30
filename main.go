package main

import (
	"rest-api-go/db"
	"rest-api-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the database connection
	db.InitDB()

	server := gin.Default()
	routes.RegisterEvents(server)
	server.Run(":8080")
}
