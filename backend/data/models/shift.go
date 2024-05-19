package models

import (
	"time"

	"gorm.io/gorm"

	"backend/data"
)

type Shift struct {
	gorm.Model
	EmployerID   uint
	Title        string
	Description  string
	PaymentRate  PaymentRate `gorm:"embedded"`
	Address      Address     `gorm:"embedded"`
	Status       string
	IsRecurring  bool
	StartDate    time.Time
	EndDate      time.Time
	WorkerShifts []WorkerShift
}

type WorkerShift struct {
	gorm.Model
	EmployerID uint
	ShiftID    uint
	Shift      Shift
	WorkerID   uint
	Worker     Worker
	Status     string
}

type ShiftQuestionnaire struct {
	ID            uint `gorm:"primary_key"`
	ShiftID       uint
	WorkerID      uint
	QuestionsJSON string
}

type ShiftApplication struct {
	ID              uint `gorm:"primary_key"`
	ShiftID         uint
	WorkerID        uint
	QuestionnaireID uint
	ResponseJSON    string
	Status          string
}

func (w ShiftApplication) RunMigration() error {
	db := data.DB
	return db.AutoMigrate(&ShiftApplication{})
}

func (w ShiftQuestionnaire) RunMigration() error {
	db := data.DB
	return db.AutoMigrate(&ShiftQuestionnaire{})
}

func (w Shift) RunMigration() error {
	db := data.DB
	return db.AutoMigrate(&Shift{})
}

func (w WorkerShift) RunMigration() error {
	db := data.DB
	return db.AutoMigrate(&WorkerShift{})
}
