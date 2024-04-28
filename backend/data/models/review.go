package models

import (
	"gorm.io/gorm"

	"backend/data"
)

type Review struct {
	gorm.Model
	Score    float32 `gorm:"type:decimal(2,1);" json:"score"` // Score from 0 to 5
	Comment  string  `json:"comment"`
	WorkerID uint    `json:"workerId"`
	// Worker         *Worker   `gorm:"foreignKey:WorkerID" json:"worker,omitempty"`                   // Omitted if Worker is an empty struct
	BusinessID *uint `gorm:"index;column:business_id" json:"businessId,omitempty"` // Nullable foreign key to a Business
	// Business       *Business `gorm:"foreignKey:BusinessID" json:"business,omitempty"`               // Omitted if Business is nil
	PrivatePartyID *uint `gorm:"index;column:private_party_id" json:"privatePartyId,omitempty"` // Nullable foreign key to a PrivateParty
}

func (s Review) RunMigration() error {
	var db = data.DB
	return db.AutoMigrate(&Review{})
}
