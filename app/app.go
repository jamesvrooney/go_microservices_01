package app

import (
	"mvc-james/controllers"
	"net/http"
)

// StartApp Starts the application
func StartApp() {

	http.HandleFunc("/user", controllers.GetUser)

	http.ListenAndServe(":8080", nil)
}
