package controllers

import (
	"encoding/json"
	"mvc-james/services"
	"net/http"
	"strconv"
)

// GetUser Returns a user
func GetUser(writer http.ResponseWriter, request *http.Request) {
	userIDParam := request.URL.Query().Get("user_id")

	userID, err := strconv.ParseInt(userIDParam, 10, 64)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("user_id must be a number"))

		return
	}

	user, err := services.GetUser(userID)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(err.Error()))

		return
	}

	jsonValue, _ := json.Marshal(*user)

	writer.Write(jsonValue)
}
