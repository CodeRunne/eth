package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infinityethback/pkg/models"
	"net/http"
	"strconv"
)

type Param struct {
	Limit string `form:"limit"`
	Offset string `form:"offset"`
}

func UIndex(c *gin.Context) {

	result, err := models.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Users not found",
		})
		return 
	}

	c.JSON(http.StatusOK, result)
}

func UShow(c *gin.Context) {
	address := c.Param("address")
    result, err := models.GetUser(address)
	fmt.Println(err)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "User not found",
        })

		return
    }
	
	c.JSON(http.StatusOK, result)
}

func UStore(c *gin.Context) {

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	
	result, err := models.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exists",
		})

		return
	}
	
	c.JSON(http.StatusCreated, result)
}


func ULeaderboard(c *gin.Context) {

	limit, offset := getLeaderboardParams(c)
	
	result, err := models.Leaderboard(offset, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Users not found",
		})
		return
	}
	
	c.JSON(http.StatusOK, result)


}

func getLeaderboardParams(c *gin.Context) (uint, uint) {
	var param Param
	var limit, offset uint
	err := c.Bind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding parameters",
		})
		return uint(0), uint(0)
	}

	if (param == Param{}) {
		limit, offset = 10, 0
		return limit, offset
	}

	l, _ := strconv.Atoi(param.Limit)
	o, _ := strconv.Atoi(param.Offset)
	
	limit = uint(l)
	offset = uint(o)

	return limit, offset
}