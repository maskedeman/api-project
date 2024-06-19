package models

import (
	"time"

	"gorm.io/gorm"
)

type Timestamp struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoCreateTime:nano" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

/*
User represents a database model for storing user-related information, extending the Timestamp
model.
*/
type User struct {
	Timestamp

	Name            string           `json:"name"`
	Email           string           `json:"email" validate:"required,email" gorm:"unique"`
	Phone           uint             `json:"phone" validate:"required" gorm:"unique" min:"10" max:"10"`
	Password        string           `json:"password" validate:"required" min:"8" `
	TransactionLogs []TransactionLog `gorm:"foreignKey:UserID"`
}
