package models

import (
	"errors"
	"gorm.io/gorm"
	db "github.com/infinityethback/pkg/database"
)

type Social struct {
	*gorm.Model
	QuoteId uint `json:"quote_id"`
	WalletAddress string `json:"wallet_address"`
	QuoteCap uint `json:"quote_cap" gorm:"default:6"`
	TwitterDailyCount uint `json:"twitter_daily_count" gorm:"default:0"`
	WhatsappDailyCount uint `json:"whatsapp_daily_count" gorm:"default:0"`
}

func CreateSocial(social *Social) (*Social, error) {

	result := db.GetDB().Create(&social)
	if result.Error != nil {
		return social, result.Error
	}

	return social, nil
}

func GetSocial(quote_id uint, address string) (*Social, error){

	var social *Social
	result := db.GetDB().Select("quote_id, wallet_address, quote_cap, twitter_daily_count, whatsapp_daily_count, created_at").Where("quote_id =? AND wallet_address =?", quote_id, address).Find(&social)

	if result.Error != nil {
        return &Social{}, result.Error
    } else if result.RowsAffected < 1 {
		return &Social{}, errors.New("Column not found")
	}

	return social, nil

}

func UpdateSocial(app string, id uint, address string) bool {
	
	var column string
	if app == "twitter" {
		column = "twitter_daily_count"
	} else if app == "whatsapp" {
		column = "whatsapp_daily_count"
	}

	result := db.GetDB().Table("socials").Where("quote_id = ?", id).Where("wallet_address = ?", address).Update(column, 1)
	if result.Error != nil {
		return false
	}
	return true
}