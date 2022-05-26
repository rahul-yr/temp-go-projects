package main

import (
	"net/http"

	"github.com/rahul-yr/temp-go-projects/employee"
)

type RegisterRoutes struct {
	RoutePath   string
	Method      string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

var AvailableRoutes = map[string]RegisterRoutes{
	"/employee/create": {
		RoutePath:   "/employee/create",
		Method:      "POST",
		HandlerFunc: employee.Create,
	},
	"/employee/update": {
		RoutePath:   "/employee/update",
		Method:      "PUT",
		HandlerFunc: employee.Update,
	},
	"/employee/delete": {
		RoutePath:   "/employee/delete",
		Method:      "DELETE",
		HandlerFunc: employee.Delete,
	},
	"/employee/read": {
		RoutePath:   "/employee/read",
		Method:      "GET",
		HandlerFunc: employee.Read,
	},
}
