package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLocation(t *testing.T) {
	store := &StubLocationStore{
		locations: map[string]Location{
			"Mary":  {name: "10"},
			"Janet": {name: "20"},
		},
	}
	server := LocationServer{store: store}

	t.Run("It should return 10 for Mary", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/location/Mary", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		expected := "10"
		actual := response.Body.String()
		assertResponse(t, expected, actual)
		assertStatusCode(t, http.StatusOK, response.Result().StatusCode)
	})

	t.Run("Should return a value and status ok for Janet", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/location/Janet", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		expected := "20"
		actual := response.Body.String()
		assertResponse(t, expected, actual)
		assertStatusCode(t, http.StatusOK, response.Result().StatusCode)
	})
	t.Run("Should return a 404 for GET request where user doesnt exist", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/location/Mark", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		expected := http.StatusNotFound
		actual := response.Result().StatusCode
		assertStatusCode(t, expected, actual)
	})

	t.Run("Should return status accepted on POSt request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/location/Mark", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		expected := http.StatusAccepted
		actual := response.Result().StatusCode
		assertStatusCode(t, expected, actual)

	})
}

func assertStatusCode(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Got %v, expected %v", actual, expected)
	}
}

func assertResponse(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Got %s, expected %s", actual, expected)
	}
}

func assertCustomerServerResponse(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Got %s, expected %s", actual, expected)
	}
}

type StubLocationStore struct {
	locations map[string]Location
}

func (s *StubLocationStore) GetLocation(name string) string {
	return s.locations[name].name
}
