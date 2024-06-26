package models

import (
	"time"

	"gorm.io/gorm"

	"backend/data"
)

type Business struct {
	gorm.Model
	UserID         uint       `gorm:"unique" json:"userID"`
	Email          string     `gorm:"type:varchar(255)" json:"email"`
	PhoneNumber    string     `json:"phoneNumber"`
	LastLogin      *time.Time `json:"lastLogin,omitempty"`
	UserType       UserType   `json:"userType"`
	ProfilePicture *string    `json:"profilePicture,omitempty"`
}

func (u Business) RunMigration() error {
	db := data.DB
	return db.AutoMigrate(&User{})
}
