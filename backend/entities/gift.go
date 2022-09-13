package entities

import "gorm.io/gorm"

type Gift struct {
	gorm.Model
	Name        string `json:"name"`
	DisplayText string `json:"displayText"`
	Amount      int    `json:"amount"`
}
