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
var routes = Routes{
	Route{
		"GET",
		"/getIP",
		GetByIP,
	},
}
