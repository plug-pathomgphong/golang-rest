package main

import (
	"github.com/gin-gonic/gin"
	"github.com/plug-pathomgphong/golang-rest/db"
	"github.com/plug-pathomgphong/golang-rest/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
