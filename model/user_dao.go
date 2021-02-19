package model

// GetUser Retieve user from database
func GetUser(userID int64) (User, error) {
	return User{
		ID:        uint64(userID),
		Firstname: "James",
		Lastname:  "Rooney",
		Email:     "jamesvrooney@hotmail.com",
	}, nil
}
