package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type Api struct {
	Router *mux.Router
}

func NewApi() {
	api := &Api{
		Router: NewRouter(),
	}

	api.Listen()
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}

func (a *Api) Listen() {
	log.Info("Listening on 8080")

	http.ListenAndServe(":8080", a.Router)
}
