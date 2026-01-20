package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestLocation(t *testing.T) {

	t.Run("It should return 10 for Mary", func(t *testing.T) {
		store := StubLocationStore{
			locations: map[int]Location{
				1: {Name: "10", Id: 1},
				2: {Name: "20", Id: 2},
			},
		}
		server := NewLocationServer(&store)
		request, _ := http.NewRequest(http.MethodGet, "/location/1", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		expected := "10"
		actual := response.Body.String()
		assertResponse(t, expected, actual)
		assertStatusCode(t, http.StatusOK, response.Result().StatusCode)
	})

	t.Run("Should return a value and status ok for Janet", func(t *testing.T) {
		store := StubLocationStore{
			locations: map[int]Location{
				1: {Name: "10", Id: 1},
				2: {Name: "20", Id: 2},
			},
		}
		server := NewLocationServer(&store)
		request, _ := http.NewRequest(http.MethodGet, "/location/2", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		expected := "20"
		actual := response.Body.String()
		assertResponse(t, expected, actual)
		assertStatusCode(t, http.StatusOK, response.Result().StatusCode)
	})
	t.Run("Should return a 404 for GET request where user doesnt exist", func(t *testing.T) {
		store := StubLocationStore{
			locations: map[int]Location{
				1: {Name: "10", Id: 1},
				2: {Name: "20", Id: 2},
			},
		}
		server := NewLocationServer(&store)
		request, _ := http.NewRequest(http.MethodGet, "/location/Mark", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		expected := http.StatusNotFound
		actual := response.Result().StatusCode
		assertStatusCode(t, expected, actual)
	})

	t.Run("Should return status accepted on POSt request", func(t *testing.T) {
		store := StubLocationStore{
			locations: map[int]Location{
				1: {Name: "10", Id: 1},
				2: {Name: "20", Id: 2},
			},
		}
		server := NewLocationServer(&store)
		expected := Location{
			Name: "Lisa",
			Id:   4,
		}
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewPostLocationRequest(expected))
		assertStatusCode(t, http.StatusAccepted, response.Result().StatusCode)
		if len(store.locations) != 3 {
			t.Errorf("New location was not added. Got locations store of length %d, expected %d", len(store.locations), 3)
		}
		assertLocationMap(t, expected, store.locations[expected.Id])
	})

	t.Run("/locations should return 200", func(t *testing.T) {
		store := StubLocationStore{
			locations: map[int]Location{
				1: {Name: "10", Id: 1},
				2: {Name: "20", Id: 2},
			},
		}
		server := NewLocationServer(&store)
		wantedLocations := []Location{{Name: "10", Id: 1}, {Name: "20", Id: 2}}
		request, _ := http.NewRequest(http.MethodGet, "/locations", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		var got []Location

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Errorf("Problem parsing response from server %q into slice of Location, %v", response.Body, err)
		}

		assertStatusCode(t, http.StatusOK, response.Code)
		if !reflect.DeepEqual(got, wantedLocations) {
			t.Errorf("Got %v, expected %v", got, wantedLocations)
		}
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
	locations map[int]Location
}

func (s *StubLocationStore) GetLocation(id int) Location {
	return s.locations[id]
}

func (s *StubLocationStore) AddLocation(location Location) {
	s.locations[location.Id] = location
}

func (s *StubLocationStore) GetLocations() []Location {
	locations := []Location{}
	for _, values := range s.locations {
		locations = append(locations, values)
	}
	return locations
}
