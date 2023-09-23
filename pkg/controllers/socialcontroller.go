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

func SStore(c *gin.Context) {

	data := map[string]interface{}{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error parsing data",
		})
		return
	}

	quote_id := uint(data["quote_id"].(float64))
	address := string(data["wallet_address"].(string))
	app := data["social"].(string)

	fetch, err := models.GetSocial(quote_id, address)
	if err != nil {

		user, ok := validateData(c, address, quote_id)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Wallet address / quote record not found",
			})
			return
		}

		var social = models.Social {
			QuoteId: quote_id,
            WalletAddress: address,
		}
		
		count := uint(user["quote_count"].(int64))
		earning := user["earnings"].(int64)

		if app == "twitter" {
			social.TwitterDailyCount = 1
		} else if app == "whatsapp" {
			social.WhatsappDailyCount = 1
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Social app not supported",
			})
			return
		}

		incrementUserCount(c, earning, count, address)

		_, err := models.CreateSocial(&social)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}

		response := fmt.Sprintf("Quote point has been added to %s", address)
		c.JSON(http.StatusCreated, gin.H{
			"message": response,
		})
		return
	} else {

		user, ok := validateData(c, address, quote_id)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Wallet address / quote record not found",
			})
			return
		}

		count := uint(user["quote_count"].(int64))
		earning := user["earnings"].(int64)
		response := fmt.Sprintf("Quote point has been added to %s", address)
		timerResponse := fmt.Sprintf("Next quote is due in the next %.f hours", float64(24) - utility.Timer(fetch.CreatedAt))
		
		if fetch.TwitterDailyCount == 1 && fetch.WhatsappDailyCount == 1 {
			if utility.Timer(fetch.CreatedAt) <= float64(24) {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": timerResponse,
				})
			}
			return
		}
		
		if fetch.TwitterDailyCount == 0 && app == "twitter" {
			ok := models.UpdateSocial("twitter", quote_id, address)
			if ok == false {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error updating database",
				})
			}

			incrementUserCount(c, earning, count, address)

			c.JSON(http.StatusOK, gin.H{
				"message": response,
			})
			return
			
		} else if fetch.WhatsappDailyCount == 0 && app == "whatsapp" {
			ok := models.UpdateSocial("whatsapp", quote_id, address)
			if ok == false {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error updating database",
				})
				return
			}

			incrementUserCount(c, earning, count, address)

			c.JSON(http.StatusOK, gin.H{
				"message": response,
			})
			return
		} else if (fetch.TwitterDailyCount == 1 && app == "twitter") || (fetch.WhatsappDailyCount == 1 && app == "whatsapp") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": timerResponse,
			})
		} else if app != "twitter" || app != "whatsapp" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Social app not supported",
			})
			return
		}
		return
	}
	

}

func validateData(c *gin.Context,address string, quote_id uint) (map[string]interface{}, bool) {
	user, _ := models.GetUser(address)
	quote, _ := models.GetQuote(quote_id)
	
	if len(user) == 0 || (*quote == models.Quote{}) {
		return user, false
	}

	return user, true
}

func incrementUserCount(c *gin.Context, earning int64, quote_count uint, address string) {
	
	env := os.Getenv("QUOTESHARE_POINT")
	earnings, _ := strconv.Atoi(env)
	earning += int64(earnings)

	count := quote_count + 1
	ok := models.UpdateQuoteCount(address, uint64(earning), count)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error updating user record",
		})
		return
	}
	return
}