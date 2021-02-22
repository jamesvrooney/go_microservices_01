package model

import (
	"fmt"
	"log"
	"mvc-james/utils"
	"net/http"
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

var (
	users = map[int64]*User{
		123: {123, "James", "Rooney", "jamesvrooney@hotmail.com"},
		234: {234, "Jack", "Rooney", "jackvrooney@hotmail.com"},
		345: {345, "Jamie", "Rooney", "jamievrooney@hotmail.com"},
	}

	// UserDao Returns ahandle to the userDao struct
	UserDao userDaoInterface
)

// GetUser Retieve user from database
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("We're accessing the database.")

	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    "The user with this user_id is not found",
		StatusCode: http.StatusNotFound,
		Code:       fmt.Sprintf("This user %v cannot be found", userID),
	}
}
