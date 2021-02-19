package services

import "mvc-james/model"

// GetUser Return a user
func GetUser(userID int64) (*model.User, error) {
	return model.GetUser(userID)
}
