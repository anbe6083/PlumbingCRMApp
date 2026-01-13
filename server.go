package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type LocationStore interface {
	GetLocation(id int) string
	AddLocation(location Location)
}

type LocationServer struct {
	store LocationStore
}

func (ls *LocationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/location/"))
	switch r.Method {
	case http.MethodGet:
		ls.processGetRequest(w, id)
	case http.MethodPost:
		ls.processPostRequest(w, r)
	}
}

func (ls *LocationServer) processGetRequest(w http.ResponseWriter, id int) {
	locationId := ls.store.GetLocation(id)
	if locationId == "" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprint(w, ls.store.GetLocation(id))
	}
}

func (ls *LocationServer) processPostRequest(w http.ResponseWriter, r *http.Request) {
	var location Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ls.store.AddLocation(location)

	w.WriteHeader(http.StatusAccepted)
}
