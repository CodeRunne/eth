package models

import (
	"errors"
	"gorm.io/gorm"
	db "github.com/infinityethback/pkg/database"
	"github.com/infinityethback/pkg/utility"
	"time"
)

type User struct {
	*gorm.Model
	WalletAddress string `json:"wallet_address" gorm:"unique"`
	Earnings	  uint64 `json:"earnings" gorm:"default: 0"`
	ReferralCount uint `json:"referral_count" gorm:"default: 0"`
	QuoteCount 	  uint `json:"quote_count" gorm:"default: 0"`
}

func CreateUser(user *User) (*User, error) {

	db.GetDB().Create(user)
	if user.ID <= 0 {
		return &User{}, errors.New("Account creation was unsuccesfull")
	}

	err := CreateReferral(&Referral{
		UserId: user.ID,
		Token: utility.GenerateToken(20),
		TokenCreatedAt: time.Now(),
	})

	if err != nil {
        return &User{}, err
    }

	return user, nil

}

func GetAllUsers() ([]*User, error) {
	var users []*User
	result := db.GetDB().Table("users").Find(&users)
	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

func GetUser(address string) (map[string]interface{}, error) {
	user := map[string]interface{}{}
    result := db.GetDB().Raw("SELECT users.ID, users.wallet_address, users.earnings, users.referral_count, users.quote_count, referrals.token FROM users INNER JOIN referrals ON users.ID = referrals.user_id WHERE wallet_address= ?", address).Scan(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func Leaderboard(offset, limit uint) ([]map[string]interface{}, error){

	users := []map[string]interface{}{}
    result := db.GetDB().Raw("SELECT wallet_address, earnings FROM users ORDER BY earnings DESC LIMIT ?, ?", offset, limit).Scan(&users)
	if result.Error != nil {
		return users, result.Error 
	}

	return users, nil
}

func GetUserById(id uint) (*User, error) {

	var user User
	result := db.GetDB().Table("users").Where("ID = ?", id).Scan(&user)
	if result.Error != nil {
		return &User{}, result.Error
	}

	return &user, nil
}

func UpdateQuoteCount(address string, earning uint64, value uint) bool {

	result := db.GetDB().Table("users").Where("wallet_address = ?", address).Updates(&User{
		Earnings: earning,
		QuoteCount: value,
	})
	if result.Error != nil {
		return false
	}

	return true

}

func UpdateReferralCount(id uint, earning uint64, value uint) bool {

	result := db.GetDB().Table("users").Where("ID = ?", id).Updates(&User{
		Earnings: earning,
		ReferralCount: value,
	})
	if result.Error != nil {
		return false
	}

	return true

}