package models

import "gorm.io/gorm"




type Diary struct {
	gorm.Model
	Content string
	UserID  uint
}