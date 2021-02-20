package services

import (
	"mvc-james/model"
	"mvc-james/utils"
)

type userService struct {
}

// UserService Defines the UserService
var UserService userService

// GetUser Return a user
func (u *userService) GetUser(userID int64) (*model.User, *utils.ApplicationError) {
	return model.GetUser(userID)
}
