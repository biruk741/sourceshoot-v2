package models

import (
	"gorm.io/gorm"

	"backend/data"
)

type Worker struct {
	gorm.Model
	UserID           uint
	User             User
	FirstName        string
	LastName         string
	Address          Address `gorm:"embedded"`
	Email            string
	PhoneNumber      string
	Description      string
	CumulativeRating float64
	Industries       []Industry `gorm:"many2many:worker_industries;"`
	WorkerShifts     []WorkerShift
}

func (w Worker) RunMigration() error {
	db := data.DB
	return db.AutoMigrate(&Worker{})
}
