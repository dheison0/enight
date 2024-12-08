package routes

import (
	"log"
	"net/http"
	"server/database"
	"server/tokens"

	"github.com/gin-gonic/gin"
)

func GetTokenUser(c *gin.Context) {
	tokenID := c.Param("id")
	user := tokens.GetUser(tokenID)
	client, err := database.GetClient(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get client!"})
		log.Printf("Can't get client from database! Client: %s Error: %s", user, err.Error())
		return
	}
	c.JSON(http.StatusOK, client)
}
