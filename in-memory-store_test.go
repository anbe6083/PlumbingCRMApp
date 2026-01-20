package main

import (
	"net/http/httptest"
	"testing"
)

func TestInMemoryStore(t *testing.T) {
	store := &InMemoryLocationStore{
		locations: map[int]Location{},
	}
	server := NewLocationServer(store)
	locations := []Location{
		{Name: "Nancy", Id: 1},
		{Name: "Steph", Id: 2},
		{Name: "Louis", Id: 3},
	}
	for _, location := range locations {
		server.ServeHTTP(httptest.NewRecorder(), NewPostLocationRequest(location))
	}

	t.Run("Should get the correct location", func(t *testing.T) {
		expected := Location{Name: "Nancy", Id: 1}
		actual := server.store.GetLocation(1)
		assertLocationMap(t, expected, actual)
		assertLocationName(t, expected.Name, actual.Name)
	})
}
