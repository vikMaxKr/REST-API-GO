package main

import (
	"fmt"
	"net/http"
	"rest-api-go/models"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 23

	models.Save(event)

	context.JSON(http.StatusCreated, gin.H{"event created success": event})
	fmt.Println("Event created:", event)
}
