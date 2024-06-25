package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string    `gorm:"not null;unique_index" json:"username"`
	Email    string    `gorm:"not null;unique_index" json:"email"`
	Password string    `gorm:"not null" json:"password"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

type UserMedications struct {
	gorm.Model

	UserID       uint `json:"userId"`
	MedicationID uint `json:"medicationId"`
	Quantity     uint `json:"quantity"`
}
