package routes

import (
	"api/database"
	"api/extra"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var tokens = map[string]string{}
var tokenLifetime = time.Minute * 60

func deleteToken(id string) {
	delete(tokens, id)
}

func deleteTokenAfterLifetime(id string) {
	<-time.After(tokenLifetime)
	deleteToken(id)
}

func getTokenPhone(token string) string {
	t, ok := tokens[token]
	if !ok {
		return ""
	}
	return t
}

func CreateToken(c *gin.Context) {
	var clientData struct {
		Phone string `json:"phone"`
		Token string `json:"token"`
	}
	if err := c.ShouldBindJSON(&clientData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind request JSON! " + err.Error()})
		return
	}
	clientData.Token = extra.RandomString(8)
	tokens[clientData.Token] = clientData.Phone
	c.JSON(http.StatusOK, clientData)
	go deleteTokenAfterLifetime(clientData.Token)
}

func GetTokenUser(c *gin.Context) {
	phone, ok := tokens[c.Param("id")]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "token not found!"})
		return
	}
	user, err := database.GetClient(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
