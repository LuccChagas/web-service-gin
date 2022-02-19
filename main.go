package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.100.255"}) // localhost

	models.ConnectDatabase()

	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumsByID)
	router.POST("/albums", controllers.PostAlbums)
	router.PATCH("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("albums/:id", controllers.DeleteAlbumByID)

	router.Run("localhost:8080")
}
