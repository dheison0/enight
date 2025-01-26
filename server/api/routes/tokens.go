package routes

import (
	"log/slog"
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
		slog.Warn("Can't get client from database!", slog.String("user", user), slog.String("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, client)
}
