package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/plug-pathomgphong/golang-rest/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	// token := context.Request.Header.Get("Authorization")
	// if token == "" {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authenticate."})
	// 	return
	// }

	// userId, err := utils.VerifyToken(token)
	// if err != nil {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authenticate."})
	// 	return
	// }

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requst data.", "error": err})
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requst data."})
		return
	}

	updateEvent.Id = eventId

	err = updateEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event delete successfully."})
}
