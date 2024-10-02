package routes

import (
	"api/database"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// we need to do better error handling here

func CreateLocation(c *gin.Context) {
	location := &models.Location{}
	if err := c.ShouldBindJSON(location); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "failed to bind request body to location model! " + err.Error()},
		)
		return
	}
	if err := database.CreateLocation(location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create location! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, location)
}

func GetAllLocations(c *gin.Context) {
	locations, err := database.GetAllLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't find locations! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, locations)
}
