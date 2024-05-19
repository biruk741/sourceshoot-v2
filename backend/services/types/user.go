package serviceTypes

import (
	"time"

	"backend/data"
	"backend/data/models"
)

type UserType = models.UserType

type User struct {
	ID             uint64       `json:"id"`
	FirebaseID     string       `json:"firebase_id"`
	Email          string       `json:"email"`
	PhoneNumber    string       `json:"phone_number"`
	LastLogin      time.Time    `json:"last_login"`
	UserType       UserType     `json:"user_type"`
	ProfilePicture string       `json:"profile_picture"`
	Address        data.Address `json:"address"`
}

type Answer struct {
	Key    string `json:"key"`
	Answer string `json:"answer"`
}

type CreateUserRequest struct {
	answers []Answer
}

// LoggedInUserResponse This struct is used to return the appropriate user given the user type
// For example, if the user type is worker, then the worker field will be populated
// and the other fields will be nil
type LoggedInUserResponse struct {
	UserType UserType         `json:"userType"`
	Business *models.Business `json:"business,omitempty"`
	Worker   *models.Worker   `json:"worker,omitempty"`
}
