package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infinityethback/pkg/models"
	"github.com/infinityethback/pkg/utility"
	"net/http"
	"os"
	"strconv"
)

func RIndex(c *gin.Context) {

	result, err := models.GetAllReferrals()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
		"message": "Users not found",
		})

		return
	}

	c.JSON(http.StatusOK, result)
}

func Refer(c *gin.Context) {

	token := c.Param("token")
	referral, err := models.GetReferral(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token not found",
		})
        return
	} else if utility.Timer(referral.TokenCreatedAt) >= float64(24) {
		models.UpdateToken(token)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token expired",
		})
        return
	}

	user, err := models.GetUserById(referral.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "User not found",
		})
		return
	}

	env := os.Getenv("REFERRAL_POINT")
	earnings, _ := strconv.Atoi(env)
	user.Earnings += uint64(earnings)
	user.ReferralCount += 1

	ok := models.UpdateReferralCount(user.ID, user.Earnings, user.ReferralCount)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H {
            "message": "Error updating user record",
        })
	}

	response := fmt.Sprintf("Referral point has been added to %s", user.WalletAddress)
	c.JSON(http.StatusBadRequest, gin.H{
		"message": response,
	})
    return

}