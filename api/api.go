package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type Api struct {
}

func NewApi() {
	api := Api{}

	api.Listen()
}

func (a *Api) Listen() {
	log.Info("Listening on 8080")

	http.ListenAndServe(":8080", NewRouter())
}
