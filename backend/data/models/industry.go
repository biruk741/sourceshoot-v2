package models

import (
	"backend/data"
)

type Industry struct {
	IndustryID   uint `gorm:"primarykey"`
	IndustryName string
	Description  string
}

func (i Industry) RunMigration() error {
	db := data.DB
	return db.AutoMigrate(&Industry{})
}
