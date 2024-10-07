package routes

import (
	"api/database"
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePurchase(c *gin.Context) {
	request := models.PurchaseRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind purchase info! " + err.Error()})
		return
	}
	response, err := database.CreatePurchase(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't insert your purchase into database! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetAllPurchases(c *gin.Context) {
	var err error
	offset := 0
	limit := 10
	if c.Query("offset") != "" {
		offset, err = strconv.Atoi(c.Query("offset"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset!"})
			return
		}
	}
	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit!"})
			return
		}
	}
	items, err := database.GetAllPurchases(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't get purhase items! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
