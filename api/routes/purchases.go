package routes

import (
	"api/database"
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePurchase(c *gin.Context) {
	var request struct {
		Token string `json:"token"`
		models.PurchaseRequest
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind purchase info! " + err.Error()})
		return
	}
	// re-write ClientPhone after binding request to avoid set it on request
	// avoiding to create many purchases without the client need
	request.ClientPhone = getTokenPhone(request.Token)
	if request.ClientPhone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token not found!"})
		return
	}
	response, err := database.CreatePurchase(&request.PurchaseRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't insert your purchase into database! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
	// delete token from availables after using it to avoid many purchases
	deleteToken(request.Token)
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
	items, err := database.GetAllPurchases(offset, limit, c.Query("search"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't get purhase items! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func GetPurchase(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id!"})
		return
	}
	item, err := database.GetPurchase(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't find item! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func SetPurchaseStage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id!"})
		return
	}
	var purchase struct {
		Stage string `json:"stage"`
	}
	if err = c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind purchase item!" + err.Error()})
		return
	}
	err = database.SetPurchaseStage(id, purchase.Stage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update stage! " + err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
