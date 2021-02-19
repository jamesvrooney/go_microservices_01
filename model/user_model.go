package model

// User Encapsulates a user object
type User struct {
	ID        uint64 `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
}
