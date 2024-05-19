package services

import (
	"backend/data/models"
	"backend/data/repo"
	serviceTypes "backend/services/types"
)

// IndustryService is an interface for managing industries.
type IndustryService interface {
	CreateIndustry(industry models.Industry) error
	GetAllIndustries() ([]models.Industry, error)
}

// LocationInstance implements the IndustryService interface.
type IndustryInstance struct {
	repo.IndustryRepo
}

// NewIndustryService creates a new IndustryService.
func NewIndustryService(industryRepo repo.IndustryRepo) *IndustryInstance {
	s := IndustryInstance{industryRepo}
	return &s
}

// CreateIndustry creates a new Industry in the database.
func (s *IndustryInstance) CreateIndustry(industry models.Industry) error {
	// Create the models record in the database
	_, err := s.IndustryRepo.CreateIndustry(industry)
	return err
}

// GetAllIndustries retrieves all industries from the database.
func (s *IndustryInstance) GetAllIndustries() ([]models.Industry, error) {
	return s.IndustryRepo.GetAllIndustries()
}

func ConvertGormIndustryToServiceIndustry(i models.Industry) serviceTypes.Industry {
	return serviceTypes.Industry{
		IndustryID:   i.ID,
		IndustryName: i.Name,
	}
}
