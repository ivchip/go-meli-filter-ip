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
		Method:  "GET",
		Pattern: "/",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome anonymous"))
		},
	},
}
