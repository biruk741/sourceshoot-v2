package services

import (
	"backend/data/models"
	"backend/data/repo"
)

type BusinessService interface {
	GetAllBusinesses() ([]models.Business, error)
	GetBusinessByUserID(userID uint64) (*models.Business, error)
	CreateBusiness(business *models.Business) error
	UpdateBusiness(id int, business *models.Business) error
	DeleteBusiness(id int) error
}

type businessService struct {
	repo repo.BusinessRepository
}

func NewBusinessService(repo repo.BusinessRepository) BusinessService {
	return &businessService{repo: repo}
}

func (s *businessService) GetAllBusinesses() ([]models.Business, error) {
	// TODO: Implement
	return nil, nil
}

func (s *businessService) GetBusinessByUserID(userID uint64) (*models.Business, error) {
	return s.repo.FindBusinessByUserID(userID)
}

func (s *businessService) CreateBusiness(business *models.Business) error {
	// TODO: Implement
	return nil
}

func (s *businessService) UpdateBusiness(id int, business *models.Business) error {
	// TODO: Implement
	return nil
}

func (s *businessService) DeleteBusiness(id int) error {
	// TODO: Implement
	return nil
}
