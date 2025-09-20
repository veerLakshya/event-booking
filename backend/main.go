package main

import (
	"backend/db"
	"backend/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Event Booking API",
		})
	})

	routes.RegisterRoutes(server)

	server.Run(":4040") //localhost:4040

}
