package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func NewPostLocationRequest(location Location) *http.Request {
	body, _ := json.Marshal(location)
	request, _ := http.NewRequest(http.MethodPost, "/location/", bytes.NewBuffer(body))
	return request
}

func assertLocationMap(t *testing.T, expected, actual Location) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Wrong map entry. Got %q expected %q", actual, expected)
	}
}

func assertLocationName(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Wrong location returned: expected: %s, Got:%s", expected, actual)
	}
}
