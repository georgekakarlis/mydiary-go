package models

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	UserName string `gorm:"column:username;unique"`
	Email    string	`gorm:"column:email;unique"`
	PasswordHash string 
	Diaries  []Diary `gorm:"foreignKey:UserID"` 
}