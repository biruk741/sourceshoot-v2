package repo

import (
	"gorm.io/gorm"

	"backend/data"
	"backend/data/models"
)

type BusinessRepository interface {
	GetAll() ([]models.Business, error)
	FindBusinessByUserID(userID uint64) (*models.Business, error)
	Create(business *models.Business) error
	Update(id int, business *models.Business) error
	Delete(id int) error
}

type businessRepository struct {
	*gorm.DB
}

func NewBusinessRepository(db *gorm.DB) BusinessRepository {
	return &businessRepository{DB: data.DB}
}

func (r *businessRepository) GetAll() ([]models.Business, error) {
	// TODO: Implement
	return nil, nil
}

func (r *businessRepository) FindBusinessByUserID(userID uint64) (*models.Business, error) {
	var business models.Business
	if err := r.DB.Where("user_id = ?", userID).First(&business).Error; err != nil {
		return nil, err
	}
	return &business, nil
}

func (r *businessRepository) Create(business *models.Business) error {
	// TODO: Implement
	return nil
}

func (r *businessRepository) Update(id int, business *models.Business) error {
	// TODO: Implement
	return nil
}

func (r *businessRepository) Delete(id int) error {
	// TODO: Implement
	return nil
}
