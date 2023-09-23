package models

import (
	"errors"
	"gorm.io/gorm"
	db "github.com/infinityethback/pkg/database"
)

type Quote struct {
	*gorm.Model
	Text string `json:"text"`
}

func CreateQuote(quote *Quote) (*Quote, error) {
	
	err := quote.ValidateQuote()
	if err != nil {
		return quote, err
	}
	
	result := db.GetDB().Create(&quote)
	if result.Error != nil {
		return quote, result.Error
	}

	return quote, nil
}

func GetLatestQuote() (*Quote, error) {

	var quote *Quote
    result := db.GetDB().Table("quotes").Last(&quote)
	if result.Error != nil {
		return quote, result.Error
	}

	return quote, nil

}


func GetQuote(id uint) (*Quote, error) {

	var quote *Quote
    result := db.GetDB().Table("quotes").Where("ID = ?", id).First(&quote)
	if result.Error != nil {
		return quote, result.Error
	}

	return quote, nil

}

func (q *Quote) ValidateQuote() error {
	if q.Text == "" {
        return errors.New("Text is required")
    } 

	return nil
}