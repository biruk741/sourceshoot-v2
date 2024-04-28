package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"backend/data/models"
	"backend/data/repo"
	types "backend/services/types"
)

type UserService interface {
	CreateFirebaseUser(firebaseID string) (uint, error)
	GetAppropriateLoggedInUser(firebaseID string) (*types.LoggedInUserResponse, error)
	GetUserByFirebaseID(id string) (*models.User, error)
}

type UserServiceInstance struct {
	repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) *UserServiceInstance {
	s := UserServiceInstance{userRepo}
	return &s
}

func (s *UserServiceInstance) CreateFirebaseUser(firebaseID string) (uint, error) {
	if firebaseID == "" {
		return 0, errors.New("invalid firebase id")
	}

	return s.CreateFirebaseUserInDB(firebaseID)
}

func (s *UserServiceInstance) GetAppropriateLoggedInUser(firebaseID string) (*types.LoggedInUserResponse, error) {
	user, err := repo.GetUserByFirebaseIDinDB(firebaseID)
	if err != nil {
		return nil, fmt.Errorf("GetUserByFirebaseIDinDB: %w", err)
	} else if user == nil {
		return nil, fmt.Errorf("GetUserByFirebaseIDinDB: %w", err)
	}

	res := types.LoggedInUserResponse{UserType: user.UserType}

	if user.UserType == "" {
		return nil, errors.New("invalid user type")
	}

	// switch user.UserType {
	// case models.USER_TYPE_WORKER:
	// 	// Handle worker user type
	// 	service := WorkerService.NewWorkerService()
	// 	worker, err := service.GetWorkerByUserID(user.Model.ID)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("GetAppropriateLoggedInUser -> GetWorkerByUserIDinDB: %w", err)
	// 	}
	// 	res.Worker = &worker
	// case models.USER_TYPE_BUSINESS:
	// 	// Handle worker user type
	// 	service := businessService.NewBusinessService()
	// 	business, err := service.GetBusinessByUserID(user.Model.ID)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("GetAppropriateLoggedInUser -> GetBusinessByUserIDinDB: %w", err)
	// 	}
	// 	res.Business = &business
	// case models.USER_TYPE_PRIVATE_PARTY:
	// 	// Handle worker user type
	// 	service := privatePartyService.NewPrivatePartyService()
	// 	privateParty, err := service.GetPrivatePartyByUserID(user.Model.ID)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("GetAppropriateLoggedInUser -> GetPrivatePartyByUserIDinDB: %w", err)
	// 	}
	// 	res.PrivateParty = &privateParty
	// }

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

func stringToSliceOfInts(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	var ints []int
	for _, part := range parts {
		i, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
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

func validateAnswers(answers []types.Answer) bool {
	for _, answer := range answers {
		if answer.Key == "" || answer.Answer == "" {
			return false
		}
	}
	return true
}
