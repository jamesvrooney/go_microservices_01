package controllers

import (
	"encoding/json"
	"mvc-james/services"
	"mvc-james/utils"
	"net/http"
	"strconv"
)

// GetUser Returns a user
func GetUser(writer http.ResponseWriter, request *http.Request) {
	userIDParam := request.URL.Query().Get("user_id")

	userID, err := strconv.ParseInt(userIDParam, 10, 64)

	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		json, _ := json.Marshal(appErr)

		writer.WriteHeader(appErr.StatusCode)
		writer.Write(json)

		return
	}

	user, appErr := services.UserService.GetUser(userID)

	if appErr != nil {
		json, _ := json.Marshal(appErr)

		writer.WriteHeader(appErr.StatusCode)
		writer.Write(json)

		return
	}

	jsonValue, _ := json.Marshal(*user)

	writer.Write(jsonValue)
}
