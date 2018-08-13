package pkg

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Get Scenes",
		"GET",
		"/scenes",
		getSceneGeoJSON,
	},
}
