package routes

import (
	"fmt"
	"net/http"
	"rest-api-go/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		if err.Error() == "record not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		}
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event //data will come from payload
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"event created success": event})
	fmt.Println("Event created:", event)
}

func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = eventId

	err = event.Update()

	if err != nil {
		if err.Error() == "record not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"event updated success": event})
	fmt.Println("Event updated:", event)
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	err = models.DeleteEventByID(eventId)

	if err != nil {
		if err.Error() == "record not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
	fmt.Println("Event deleted with ID:", eventId)
}
