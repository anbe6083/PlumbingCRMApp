package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETCustomer(t *testing.T) {

	t.Run("Should get a customer balance for Mary", func(t *testing.T) {
		request := newGetCustomerBalanceRequest("Mary")
		response := httptest.NewRecorder()

		store := &StubCustomerStore{map[string]int{"Mary": 10000, "Adam": 20000}}
		server := &CustomerServer{store: store}
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10000"
		assertCustomerBalance(t, got, want)
		assertStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("Should get a customer balance for Adam", func(t *testing.T) {
		request := newGetCustomerBalanceRequest("Adam")
		response := httptest.NewRecorder()

		store := &StubCustomerStore{map[string]int{"Mary": 10000, "Adam": 20000}}
		server := &CustomerServer{store: store}
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "20000"
		assertCustomerBalance(t, got, want)
		assertStatusCode(t, response.Code, http.StatusOK)

	})

	t.Run("Should get a 404 for customers who don't exist", func(t *testing.T) {
		request := newGetCustomerBalanceRequest("Nancy")
		response := httptest.NewRecorder()

		store := &StubCustomerStore{map[string]int{"Mary": 10000, "Adam": 20000}}
		server := &CustomerServer{store: store}
		server.ServeHTTP(response, request)
		assertStatusCode(t, response.Code, NotFoundStatus)
	})

}

func assertCustomerBalance(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, Wanted %q", got, want)
	}

}

func newGetCustomerBalanceRequest(c string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/customer/%s", c), nil)
	return request
}

type StubCustomerStore struct {
	balances map[string]int
}

func (s *StubCustomerStore) GetCustomerBalance(name string) int {
	var balance = s.balances[name]
	return balance
}

func assertStatusCode(t testing.TB, got int, want int) {
	if got != want {
		t.Errorf("Wrong status code, got %v expected %v", got, want)
	}
}
