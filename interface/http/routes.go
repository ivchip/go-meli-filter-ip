package http

import "net/http"

// Route define a structure to create routes
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routesPrivate = Routes{
	Route{
		"GET",
		"/getIP",
		GetByIP,
	},
}

var routesPublic = Routes{
	Route{
		Method:      "GET",
		Pattern:     "/ipBlocking",
		HandlerFunc: GetAll,
	},
	Route{
		Method:      "GET",
		Pattern:     "/ipBlocking/{ip}",
		HandlerFunc: GetByIp,
	},
	Route{
		Method:      "POST",
		Pattern:     "/ipBlocking",
		HandlerFunc: Create,
	},
	Route{
		Method:      "PUT",
		Pattern:     "/ipBlocking/{ip}",
		HandlerFunc: Update,
	},
}
