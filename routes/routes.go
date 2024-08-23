package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/plug-pathomgphong/golang-rest/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authentication := server.Group("/")
	authentication.Use(middlewares.Authentication)
	authentication.POST("/events", createEvent)
	authentication.PUT("/events/:id", updateEvent)
	authentication.DELETE("/events/:id", deleteEvent)
	authentication.POST("/events/:id/register", registerForEvent)
	authentication.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
