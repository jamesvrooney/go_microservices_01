package model

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)

	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "We were expecting an error when the user id is 0.")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "The user with this user_id is not found", err.Message)
	assert.EqualValues(t, "This user 0 cannot be found", err.Code)
}

func TestGetUserUserFound(t *testing.T) {
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "James", user.Firstname)
	assert.EqualValues(t, "Rooney", user.Lastname)
	assert.EqualValues(t, "jamesvrooney@hotmail.com", user.Email)
}
