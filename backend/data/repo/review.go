package repo

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	"backend/data"
	"backend/data/models"
)

type ReviewRepo interface {
	CreateReview(review models.Review) error
	GetWorkerReviews(workerID uint, limit int, sortBy string) ([]models.Review, error)
	CalculateWorkerRating(workerID uint) (float32, error)
	UpdateWorkerRating(workerID uint) error
}

type ReviewRepoInstance struct {
	db *gorm.DB
}

func NewReviewRepo() ReviewRepoInstance {
	return ReviewRepoInstance{db: data.DB}
}

// CreateReview creates a new review in the database
func (r ReviewRepoInstance) CreateReview(review models.Review) error {
	if review.Score < 0 || review.Score > 5 {
		return fmt.Errorf("score must be between 0 and 5")
	}

	// Logic to check if only one of BusinessID or PrivatePartyID is provided can be added here.

	err := r.db.Create(review).Error
	if err != nil {
		return err
	}

	// Update the worker's rating
	return r.UpdateWorkerRating(review.WorkerID)
}

// GetWorkerReviewsFromDB retrieves all reviews for a given worker
func (r ReviewRepoInstance) GetWorkerReviews(workerID uint, limit int, sortBy string) ([]models.Review, error) {
	var reviews []models.Review
	db := data.DB

	if len(strings.Split(sortBy, " ")) < 2 {
		sortBy = sortBy + " desc"
	}

	tx := db.Preload("Worker").
		Preload("Business", "id IS NOT NULL").
		Preload("Business.User", "id IS NOT NULL").
		Where("worker_id = ?", workerID).
		Order(sortBy). // Add sort by functionality
		Limit(limit)   // Add limit

	if err := tx.Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

// CalculateWorkerRating calculates the average rating for a given worker
func (r ReviewRepoInstance) CalculateWorkerRating(workerID uint) (float32, error) {
	var averageRating *float32
	err := r.db.Model(&models.Review{}).Where("worker_id = ?", workerID).
		Select("AVG(score) as average_rating").Scan(&averageRating).Error
	if err != nil || averageRating == nil {
		return 0, err
	}

	return *averageRating, nil
}

// UpdateWorkerRating updates the worker's rating after a new review is added
func (r ReviewRepoInstance) UpdateWorkerRating(workerID uint) error {
	averageRating, err := r.CalculateWorkerRating(workerID)
	if err != nil {
		return err
	}

	// Update the worker's rating, considering potential NULL which means no rating yet
	return r.db.Model(&Worker{}).Where("id = ?", workerID).UpdateColumn("rating", gorm.Expr("COALESCE(?, rating)", averageRating)).Error
}
