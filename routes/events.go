package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Niranjini-Kathiravan/go-rest-api-v2/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {
	log.Println("POST /events called")

	var event models.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		log.Println("Failed to bind JSON:", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse data.",
			"error":   err.Error(),
		})
		return
	}

	event.ID = 1
	event.UserID = 1

	err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later"})
		return
	}

	log.Println("Event saved:", event)

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		return
	}

	var updatedEvent models.Event
	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event data.", "error": err.Error()})
		return
	}

	updatedEvent.ID = eventId
	updatedEvent.UserID = 1 // temporarily hardcoded

	if err := updatedEvent.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not send back the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}
