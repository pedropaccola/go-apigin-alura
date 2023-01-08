package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}
