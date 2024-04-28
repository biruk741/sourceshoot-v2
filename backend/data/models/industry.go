package models

import (
	"gorm.io/gorm"

	"backend/data"
)

type Industry struct {
	gorm.Model
	Name       string
	Businesses []Business `gorm:"many2many:employer_industries;"`
	Workers    []Worker   `gorm:"many2many:worker_industries;"`
}

func (i Industry) RunMigration() error {
	db := data.DB
	return db.AutoMigrate(&Industry{})
}
