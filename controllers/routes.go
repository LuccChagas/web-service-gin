package controllers

import (
	"example/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)

	c.IndentedJSON(http.StatusOK, gin.H{"data": albums})
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var input models.CreateAlbumImput

	// Validade Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album := models.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  input.Price,
	}

	models.DB.Create(&album)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

// getAlbumByID locates the album whose ID value matches the ID
func GetAlbumsByID(c *gin.Context) {
	var album models.Album

	ID := c.Param("id")

	if err := models.DB.Where("id = ?", ID).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func UpdateAlbum(c *gin.Context) {
	var input models.UpdateAlbumInput
	var album models.Album

	ID := c.Param("id")

	// check if ID exists
	if err := models.DB.Where("id = ?", ID).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}

	// Bind album object
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&album).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func DeleteAlbumByID(c *gin.Context) {
	var album models.Album

	ID := c.Param("id")

	// check if ID exists
	if err := models.DB.Where("id = ?", ID).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}

	models.DB.Delete(&album)

	c.JSON(http.StatusOK, gin.H{"data": album})
}
