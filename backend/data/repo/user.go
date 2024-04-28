package repo

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"backend/data"
	"backend/data/models"
)

type UserRepo interface {
	CreateUserInDB(user models.User) (uint, error)
	CreateFirebaseUserInDB(firebaseID string) (uint, error)
	GetUserByFirebaseIDinDB(id string) (*models.User, error)
	UpdateUserLastLoginInDB(firebaseID string) error
}

type UserRepoInstance struct {
	db *gorm.DB
}

func NewUserRepo() UserRepo {
	return &UserRepoInstance{db: data.DB}
}

// CreateUserInDB This function checks if a user for some firebaseID exists and if so, fills in the rest of the data.
// Otherwise, it creates a new user with the given firebaseID.
func (r UserRepoInstance) CreateUserInDB(user models.User) (uint, error) {
	db := data.DB

	// Check if a user with the given firebaseID exists
	var existingUser models.User
	query := db.Where("firebase_id = ?", user.FirebaseID).First(&existingUser).Scan(&existingUser)
	if err := query.Error; err != nil {
		// If the error is not "record not found", return the error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, err
		}
		// Otherwise, create a new user
		if err := db.Create(&user).Error; err != nil {
			return 0, err
		}
		return user.Model.ID, nil
	}

	// If a user with the given firebaseID exists, update the user
	if err := db.Where("firebase_id = ?", user.FirebaseID).Updates(user).Error; err != nil {
		return 0, err
	}

	return user.Model.ID, nil
}

// This function creates a user only with the firebaseId column filled in to be used for authentication
func (r UserRepoInstance) CreateFirebaseUserInDB(firebaseID string) (uint, error) {
	db := data.DB
	user := models.User{FirebaseID: firebaseID}

	if err := db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.Model.ID, nil
}

func (r UserRepoInstance) GetUserByFirebaseIDinDB(id string) (*models.User, error) {
	var user models.User
	db := data.DB

	// Use the Where clause to filter records by firebase_id
	if err := db.Where("firebase_id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// This function updates the last sign in date for the user to the current date
func (r UserRepoInstance) UpdateUserLastLoginInDB(firebaseID string) error {
	db := data.DB
	now := time.Now()

	query := db.Model(&models.User{}).Where("firebase_id = ?", firebaseID).Update("last_login", now)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}
