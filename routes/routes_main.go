package routes

import "github.com/kataras/iris"

// Initialize our routers
var Routers = Routes{
	// study MiniApp api
	studyRoutes,

	// test Routes
	test,
}

// Defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
	Name        string //
	Method      string // if Method=Party , Routes!=nil
	Pattern     string
	HandlerFunc iris.Handler
	Routes      Routes
}

// Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route
