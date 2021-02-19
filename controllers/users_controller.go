package controllers

import (
	"fmt"
	"mvc-james/services"
	"net/http"
	"strconv"
)

// GetUser Returns a user
func GetUser(writer http.ResponseWriter, request *http.Request) {
	userIDParam := request.URL.Query().Get("user_id")

	userId, err := strconv.ParseInt(userIDParam, 10, 64)

	if err != nil {
		return
	}

	user, err := services.GetUser(userId)

	if err != nil {
		return
	}

	fmt.Fprint(writer, user)
}
