package repo

import (
	"gorm.io/gorm"

	"backend/data"
	"backend/data/models"
)

type WorkerRepo interface {
	GetWorkerByUserID(userID uint) (models.Worker, error)
	CreateWorker(worker models.Worker) (uint, error)
}

type WorkerRepoInstance struct {
	db *gorm.DB
}

func NewWorkerRepo() WorkerRepoInstance {
	return WorkerRepoInstance{db: data.DB}
}

func (w WorkerRepoInstance) CreateWorker(worker models.Worker) (uint, error) {
	query := w.db.Create(&worker)
	if err := query.Error; err != nil {
		return 0, err
	}

	return worker.ID, nil
}

func (w WorkerRepoInstance) GetWorkerByUserID(userID uint) (models.Worker, error) {
	worker := models.Worker{}
	if err := w.db.Model(&worker).Where("user_id = ?", userID).Error; err != nil {
		return models.Worker{}, err
	}
	return worker, nil
}
