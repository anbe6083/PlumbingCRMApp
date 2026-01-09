package main

import (
	"fmt"
	"net/http"
	"strings"
)

type LocationStore interface {
	GetLocation(name string) string
}

type LocationServer struct {
	store LocationStore
}

func (ls *LocationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/location/")
	switch r.Method {
	case http.MethodGet:
		locationId := ls.store.GetLocation(name)
		if locationId == "" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			ls.processGetRequest(w, name)
		}
	case http.MethodPost:
		ls.processPostRequest(w)
	}
}

func (ls *LocationServer) processGetRequest(w http.ResponseWriter, name string) {
	fmt.Fprint(w, ls.store.GetLocation(name))
}

func (ls *LocationServer) processPostRequest(w http.ResponseWriter) {

	w.WriteHeader(http.StatusAccepted)
}
