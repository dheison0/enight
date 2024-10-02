package routes

import (
	"api/database"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClient(c *gin.Context) {
	var client models.ClientDatabase
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse client information! " + err.Error()})
		return
	}
	err := database.CreateClient(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create client! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

func GetAllClients(c *gin.Context) {
	clients, err := database.GetAllClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get clients! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func GetClient(c *gin.Context) {
	phone := c.Param("phone")
	client, err := database.GetClient(phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to find client! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

func DeleteClient(c *gin.Context) {
	phone := c.Param("phone")
	err := database.DeleteClient(&models.ClientDatabase{Phone: phone})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete client! " + err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
