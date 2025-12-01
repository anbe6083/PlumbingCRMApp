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

		CustomerServer(response, request)

		got := response.Body.String()
		want := "10000"
		assertCustomerBalance(t, got, want)
	})

	t.Run("Should get a customer balance for Adam", func(t *testing.T) {
		request := newGetCustomerBalanceRequest("Adam")
		response := httptest.NewRecorder()

		CustomerServer(response, request)

		got := response.Body.String()
		want := "20000"
		assertCustomerBalance(t, got, want)
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
