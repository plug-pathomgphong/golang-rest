package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/plug-pathomgphong/golang-rest/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find event id."})
		return
	}

	// save event id with user id
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not register event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Register has been succesfully."})

}
func cancelRegistration(context *gin.Context) {
	// userId := context.GetInt64("userId")

	// eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
	// 	return
	// }
}
