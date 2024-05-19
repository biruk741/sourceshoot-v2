package models

import (
	"gorm.io/gorm"

	"backend/data"
)

type Skill struct {
	gorm.Model
	SkillName   string
	Description string
	IndustryID  uint     `gorm:"index"`
	Industry    Industry `gorm:"foreignKey:ID;references:IndustryID"`
}

func (s Skill) RunMigration() error {
	db := data.DB
	return db.AutoMigrate(&Skill{})
}
