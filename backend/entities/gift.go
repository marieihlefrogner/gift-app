package entities

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Gift struct {
	gorm.Model
	GiftID      uuid.UUID `gorm:"type:char(36);unique;not null;" json:"giftId"`
	Name        string    `json:"name"`
	DisplayText string    `json:"displayText"`
	Amount      int       `json:"amount"`
}

func (gift *Gift) BeforeCreate(tx *gorm.DB) (err error) {
	gift.GiftID = uuid.NewV4()
	return
}
