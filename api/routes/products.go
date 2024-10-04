package routes

import (
	"api/database"
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	product := models.ProductResponse{}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind product model JSON! " + err.Error()})
		return
	}
	if err := database.CreateProduct(&product.Product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create product! " + err.Error()})
		return
	}
	for idx := range product.Sizes {
		// write everthing directly to struct for future use on response data
		product.Sizes[idx].ProductID = product.ID
		if err := database.AddProductSize(&product.Sizes[idx]); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert product size! " + err.Error()})
			// use defer to reduce the delay of closing connection after sending response
			defer database.DeleteProduct(product.Product.ID)
			return
		}
	}
	c.JSON(http.StatusOK, product)
}

func AddProductSize(c *gin.Context) {
	rawID := c.Param("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id!"})
		return
	}
	size := models.ProductSize{ProductID: id}
	if err = c.ShouldBindJSON(&size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind JSON to product size body! " + err.Error()})
		return
	}
	if err = database.AddProductSize(&size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to add product size! " + err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func GetAllProducts(c *gin.Context) {
	products, err := database.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get products! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	rawID := c.Param("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id!"})
		return
	}
	product, err := database.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't find product! " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	rawID := c.Param("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id!"})
		return
	}
	err = database.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete product! " + err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteProductSize(c *gin.Context) {
	rawProductID := c.Param("id")
	rawSizeID := c.Param("sid")
	var pid, sid int
	var err error
	if pid, err = strconv.Atoi(rawProductID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id!"})
		return
	} else if sid, err = strconv.Atoi(rawSizeID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid size id!"})
		return
	}
	err = database.DeleteProductSize(&models.ProductSize{ID: sid, ProductID: pid})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete size! " + err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
