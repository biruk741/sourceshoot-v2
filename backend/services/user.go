package services

import (
	"errors"
	"fmt"

	"backend/data/models"
	"backend/data/repo"
	types "backend/services/types"
)

type UserService interface {
	CreateFirebaseUser(firebaseID string) (uint, error)
	GetAppropriateLoggedInUser(firebaseID string) (*types.LoggedInUserResponse, error)
	GetUserByFirebaseID(id string) (*models.User, error)
	CreateUserFromAnswers(firebaseID string, answers []types.Answer) (interface{}, interface{})
}

type UserServiceInstance struct {
	repo.UserRepo

	// WorkerService is a service that provides methods to interact with the worker table in the database
	WorkerService WorkerService
}

func NewUserService(userRepo repo.UserRepo, workerService WorkerService) *UserServiceInstance {
	s := UserServiceInstance{userRepo, workerService}
	return &s
}

func (s *UserServiceInstance) CreateFirebaseUser(firebaseID string) (uint, error) {
	if firebaseID == "" {
		return 0, errors.New("invalid firebase id")
	}

	return s.CreateFirebaseUserInDB(firebaseID)
}

func (s *UserServiceInstance) GetAppropriateLoggedInUser(firebaseID string) (*types.LoggedInUserResponse, error) {
	user, err := s.UserRepo.GetUserByFirebaseIDinDB(firebaseID)
	if err != nil {
		return nil, fmt.Errorf("GetUserByFirebaseIDinDB: %w", err)
	} else if user == nil {
		return nil, fmt.Errorf("GetUserByFirebaseIDinDB: %w", err)
	}

	res := types.LoggedInUserResponse{UserType: user.UserType}

	if user.UserType == "" {
		return nil, errors.New("invalid user type")
	}

	switch user.UserType {
	case models.USER_TYPE_WORKER:
		// Handle worker user type
		worker, err := s.WorkerService.GetWorkerByUserID(user.Model.ID)
		if err != nil {
			return nil, fmt.Errorf("GetAppropriateLoggedInUser -> GetWorkerByUserIDinDB: %w", err)
		}
		res.Worker = &worker
	case models.USER_TYPE_BUSINESS:
		// Handle worker user type
		service := businessService.NewBusinessService()
		business, err := service.GetBusinessByUserID(user.Model.ID)
		if err != nil {
			return nil, fmt.Errorf("GetAppropriateLoggedInUser -> GetBusinessByUserIDinDB: %w", err)
		}
		res.Business = &business
	}

	return &res, nil
}

func (s *UserServiceInstance) GetUserByFirebaseID(id string) (*models.User, error) {
	user, err := s.UserRepo.GetUserByFirebaseIDinDB(id)
	if err != nil {
		return nil, fmt.Errorf("GetUserByFirebaseIDinDB: %w", err)
	} else if user == nil {
		return nil, fmt.Errorf("GetUserByFirebaseIDinDB: %w", err)
	}

	return user, nil
}

// This function updates the last sign in date for the user
func (s *UserServiceInstance) UpdateUserLastLogin(firebaseID string) error {
	if firebaseID == "" {
		return errors.New("invalid firebase id")
	}
	return s.UserRepo.UpdateUserLastLoginInDB(firebaseID)
}

func (s *UserServiceInstance) CreateUserFromAnswers(firebaseID string, answers []types.Answer) (interface{}, interface{}) {
	user, err := convertAnswersToUser(answers)
	if err != nil {
		return nil, err
	}
	user.FirebaseID = firebaseID
	userID, err := s.CreateUserInDB(user)
	if err != nil {
		return nil, err
	}
	switch user.UserType {
	case models.USER_TYPE_WORKER:
		worker := models.Worker{
			UserID: userID,
		}
		workerID, err := WorkerService.NewWorkerService().CreateWorker(worker)
		if err != nil {
			return nil, err
		}
		return workerID, nil
	case models.USER_TYPE_BUSINESS:
		business := models.Business{
			UserID: userID,
		}
		businessID, err := businessService.NewBusinessService().CreateBusiness(business)
		if err != nil {
			return nil, err
		}
		return businessID, nil
	default:
		return nil, errors.New("invalid user type")
	}
}

func convertAnswersToUser(answers []types.Answer) (models.User, error) {
	var user models.User
	for _, answer := range answers {
		switch answer.Key {
		case "email":
			user.Email = answer.Answer
		case "phone_number":
			user.PhoneNumber = answer.Answer
		case "user_type":
			user.UserType = convertUserTypeToEnum(answer.Answer)
		}
	}
	return user, nil // todo return error if user is not valid
}

func convertUserTypeToEnum(answer string) models.UserType {
	var enum models.UserType
	switch answer {
	case "worker":
		enum = models.USER_TYPE_WORKER
	case "private_party":
		enum = models.USER_TYPE_PRIVATE_PARTY
	case "business":
		enum = models.USER_TYPE_BUSINESS
	default:
		enum = models.USER_TYPE_WORKER
	}
	return enum
}
