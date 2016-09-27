package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type Api struct {
	router *mux.Router
}

func NewApi() {
	api := &Api{
		router: NewRouter(),
	}

	api.Listen()
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.method).
			Path(route.path).
			Name(route.name).
			Handler(route.handler)
	}

	return router
}

func (a *Api) Listen() {
	log.Info("Listening on 8080")

	http.ListenAndServe(":8080", a.router)
}
