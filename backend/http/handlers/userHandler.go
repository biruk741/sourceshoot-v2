package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	http2 "backend/http/auth"
	errors2 "backend/http/errors"
	"backend/services"
	serviceTypes "backend/services/types"
)

func (h Handler) HandleCreateUser(c *gin.Context) {
	r := http2.CreateUserRequest{}

	firebaseID, err := http2.GetFirebaseIDFromContext(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.BindJSON(&r)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		fmt.Print(err)
		return
	}

	userService := services.NewUserService()

	user, err := h.UserService.CreateUserFromAnswers(*firebaseID, r.Answers)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h Handler) HandleGetLoggedInUser(c *gin.Context) {
	firebaseID, err := http2.GetFirebaseIDFromContext(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	res, err := h.UserService.GetAppropriateLoggedInUser(*firebaseID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h Handler) HandleFirebaseSignIn(c *gin.Context) {
	user, err := http2.GetUserFromContext(c)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// err = h.UserService.(user.FirebaseID)
	// if err != nil {
	// 	c.AbortWithError(http.StatusInternalServerError, err)
	// 	return
	// }

	c.JSON(http.StatusOK, serviceTypes.User{
		FirebaseID:  user.FirebaseID,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		UserType:    user.UserType,
	})
}

// This function signs up the user if they don't exist in the database
func (h Handler) HandleFirebaseSignUp(c *gin.Context) {
	firebaseID, err := http2.GetFirebaseIDFromContext(c)
	if err != nil || firebaseID == nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := http2.GetUserFromContext(c)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	if !errors.Is(err, errors2.UserNotFoundError) || user != nil {
		c.AbortWithError(http.StatusConflict, err)
		return
	}

	userID, err := h.UserService.CreateFirebaseUser(*firebaseID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": userID})
}
