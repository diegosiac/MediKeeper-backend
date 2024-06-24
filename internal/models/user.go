package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"`
	Email     string `gorm:"not null;unique_index" json:"email"`
	Password  string `gorm:"not null" json:"password"`
}
