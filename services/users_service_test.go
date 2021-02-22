package services

import (
	"mvc-james/model"
	"mvc-james/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	getUserfunction func(int64) (*model.User, *utils.ApplicationError)
)

func init() {
	model.UserDao = &userDaoMock{}
}

type userDaoMock struct{}

func (u *userDaoMock) GetUser(userID int64) (*model.User, *utils.ApplicationError) {
	return getUserfunction(userID)
}

func TestGetUserNotFound(t *testing.T) {
	getUserfunction = func(int64) (*model.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "The user with this user_id is not found",
			Code:       "This user 0 cannot be found",
		}
	}

	user, err := UserService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "The user with this user_id is not found", err.Message)
	assert.EqualValues(t, "This user 0 cannot be found", err.Code)
}

// func TestGetUserNoErrorUserFound(t *testing.T) {
// 	user, err := UserService.GetUser(0)

// 	assert.Nil(t, err)
// 	assert.NotNil(t, user)
// 	// assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
// 	// assert.EqualValues(t, "The user with this user_id is not found", err.Message)
// 	// assert.EqualValues(t, "This user 0 cannot be found", err.Code)
// }
