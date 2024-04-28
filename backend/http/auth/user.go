package auth

import (
	"errors"

	"github.com/gin-gonic/gin"

	"backend/services"
	serviceTypes "backend/services/types"
)

type CreateUserRequest struct {
	Answers []serviceTypes.Answer `json:"answers"`
}

type CreateFirebaseUserRequest struct {
	Answers    []serviceTypes.Answer `json:"answers"`
	FirebaseID string                `json:"firebaseID"`
}

func GetUserFromContext(c *gin.Context) (*serviceTypes.User, error) {
	firebaseID, ok := c.Get("firebaseID")
	if !ok || firebaseID == nil || firebaseID == "" {
		return nil, errors.New("user not found in context")
	}

	idParsed, ok := firebaseID.(string)
	if !ok {
		return nil, errors.New("invalid user id in context")
	}

	// todo: use dependency injection
	userService := services.NewUserService()

	return userService.GetUserByFirebaseID(idParsed)
}
func GetFirebaseIDFromContext(c *gin.Context) (*string, error) {
	firebaseID, ok := c.Get("firebaseID")
	if !ok || firebaseID == nil || firebaseID == "" {
		return nil, errors.New("user not found in context")
	}

	idParsed, ok := firebaseID.(string)
	if !ok {
		return nil, errors.New("invalid user id in context")
	}

	return &idParsed, nil
}
