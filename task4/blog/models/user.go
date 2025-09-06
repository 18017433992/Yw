package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique" json:"Username"`
	Password     string `json:"Password"`
	Email        *string
	TokenVersion uint `gorm:"default:1"`
}
