package model

import (
	"fmt"
)

var (
	users = map[int64]*User{
		123: {123, "James", "Rooney", "jamesvrooney@hotmail.com"},
		234: {234, "Jack", "Rooney", "jackvrooney@hotmail.com"},
		345: {345, "Jamie", "Rooney", "jamievrooney@hotmail.com"},
	}
)

// GetUser Retieve user from database
func GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, fmt.Errorf("This user %v cannot be found", userID)
}
