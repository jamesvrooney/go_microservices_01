package services

import (
	"mvc-james/model"
	"mvc-james/utils"
)

// GetUser Return a user
func GetUser(userID int64) (*model.User, *utils.ApplicationError) {
	return model.GetUser(userID)
}
