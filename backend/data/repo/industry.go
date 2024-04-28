package repo

import (
	"gorm.io/gorm"

	"backend/data"
	"backend/data/models"
)

type IndustryRepo interface {
	CreateIndustry(industry models.Industry) (uint, error)
	GetAllIndustries() ([]models.Industry, error)
}

type IndustryRepoInstance struct {
	db *gorm.DB
}

func NewIndustryRepo() IndustryRepoInstance {
	return IndustryRepoInstance{db: data.DB}
}

func (r IndustryRepoInstance) CreateIndustry(industry models.Industry) (uint, error) {
	db := data.DB
	query := db.Create(&industry)
	if err := query.Error; err != nil {
		return 0, err
	}

	return industry.IndustryID, nil
}

// GetAllIndustries retrieves all industries from the database.
func (r IndustryRepoInstance) GetAllIndustries() ([]models.Industry, error) {
	var industries []models.Industry

	// Use Find to retrieve all records from the "industries" table
	if err := r.db.Find(&industries).Error; err != nil {
		return nil, err
	}

	return industries, nil
}
