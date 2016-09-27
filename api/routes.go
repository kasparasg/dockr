package api

import "net/http"

type Route struct {
	name    string
	method  string
	path    string
	handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"CreateBuild", "POST", "/builds", CreateBuild},
	Route{"CreateComposeBuild", "POST", "/compose-builds", CreateComposeBuild},
}
