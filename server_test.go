package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETCustomer(t *testing.T) {

	t.Run("Should get a customer balance for Mary", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/customer/Mary", nil)
		response := httptest.NewRecorder()

		CustomerServer(response, request)

		got := response.Body.String()
		want := "10000"
		assertCustomerBalance(t, got, want)
	})

	t.Run("Should get a customer balance for Adam", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/customer/Adam", nil)
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
