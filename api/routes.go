package api

import "net/http"

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"CreateBuild", "POST", "/builds", CreateBuild},
}
