package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infinityethback/pkg/models"
	"github.com/infinityethback/pkg/utility"
	"net/http"
)

func QStore(c *gin.Context) {

	var quote models.Quote
	c.BindJSON(&quote)

	lQuote, err := models.GetLatestQuote()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	} 
	
	if utility.Timer(lQuote.CreatedAt) <= float64(24) {
		response := fmt.Sprintf("Next quote is due in the next %.f hours", float64(24) - utility.Timer(lQuote.CreatedAt))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": response,
		})
	} else {
		result, err := models.CreateQuote(&quote)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error creating quote",
			})
			return
		}

		c.JSON(http.StatusCreated, result)
		return
	}

}

func QShow(c *gin.Context) {

	result, err := models.GetLatestQuote()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, result)

}