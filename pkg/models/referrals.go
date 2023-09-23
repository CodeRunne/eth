package models

import (
	"gorm.io/gorm"
	db "github.com/infinityethback/pkg/database"
	"github.com/infinityethback/pkg/utility"
	"time"
)

type Referral struct {
	*gorm.Model
	UserId uint `json:"user_id"`
	Token string `json:"token"`
	TokenCreatedAt time.Time `json:"token_created_at"`
}

func CreateReferral(referral *Referral) (error) {
	
	result := db.GetDB().Table("referrals").Create(&referral)
	if result.Error!= nil {
        return result.Error
    }

	return nil
}

func GetAllReferrals() ([]*Referral, error){

	var referrals []*Referral
	result := db.GetDB().Table("referrals").Find(&referrals)
	if result.Error != nil {
		return referrals, result.Error
	}

	return referrals, nil
}

func GetReferral(token string) (*Referral, error) {

	var referral Referral
	result := db.GetDB().Table("referrals").Where("token = ?", token).First(&referral)
	if result.Error != nil {
        return &Referral{}, result.Error
    }

	return &referral, nil

}

func UpdateToken(token string) bool {
	result := db.GetDB().Table("referrals").Where("token = ?", token).Updates(&Referral{Token: utility.GenerateToken(20), TokenCreatedAt: time.Now()})

	if result.Error != nil {
		return false
	}

	return true
}