package services

import (
	"mvc-james/model"
	"mvc-james/utils"
)

func init() {
	UserService = &userService{}
}

type userServiceInterface interface {
	GetUser(int64) (*model.User, *utils.ApplicationError)
}

type userService struct{}

// UserService Defines the UserService
var UserService userServiceInterface

// GetUser Return a user
func (u *userService) GetUser(userID int64) (*model.User, *utils.ApplicationError) {
	return model.UserDao.GetUser(userID)
}
