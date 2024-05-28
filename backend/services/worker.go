package services

import (
	"backend/data/models"
	"backend/data/repo"
)

type WorkerService interface {
	GetWorkerByUserID(userID uint) (models.Worker, error)
}

type WorkerInstance struct {
	repo.WorkerRepo
}

func NewWorkerService(workerRepo repo.WorkerRepo) *WorkerInstance {
	s := WorkerInstance{workerRepo}
	return &s
}

func (w *WorkerInstance) GetWorkerByUserID(userID uint) (models.Worker, error) {
	return w.WorkerRepo.GetWorkerByUserID(userID)
}
