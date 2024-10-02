package routes

import (
	"api/database"
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// we need to do better error handling here

// here we need some informacion of the location, the name and distance, id will be ignored
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

// this endpoint can delete one location per time, it only needs the ID of the item
func DeleteLocation(c *gin.Context) {
	locationIDParam := c.Param("id")
	locationID, err := strconv.Atoi(locationIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id. it must be an integer!"})
		return
	}
	if err := database.DeleteLocation(&models.Location{ID: locationID}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete item! " + err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// all locations inserted in the database will be returned
func GetAllLocations(c *gin.Context) {
	locations, err := database.GetAllLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't find locations! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, locations)
}
