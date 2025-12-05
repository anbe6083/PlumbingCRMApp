package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"
)

func TestGETCustomer(t *testing.T) {

	t.Run("Should get a customer balance for Mary", func(t *testing.T) {
		request := newGetCustomerBalanceRequest("Mary")
		response := httptest.NewRecorder()

		store := &StubCustomerStore{[]Customer{{Name: "Mary", Balance: 10000}, {Name: "Adam", Balance: 20000}}}
		server := NewCustomerServer(store)
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10000"
		assertCustomerBalance(t, got, want)
		assertStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("Should get a customer balance for Adam", func(t *testing.T) {
		request := newGetCustomerBalanceRequest("Adam")
		response := httptest.NewRecorder()

		store := &StubCustomerStore{[]Customer{{Name: "Mary", Balance: 10000}, {Name: "Adam", Balance: 20000}}}
		server := NewCustomerServer(store)
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "20000"
		assertCustomerBalance(t, got, want)
		assertStatusCode(t, response.Code, http.StatusOK)

	})

	t.Run("Should get a 404 for customers who don't exist", func(t *testing.T) {
		request := newGetCustomerBalanceRequest("Nancy")
		response := httptest.NewRecorder()

		store := &StubCustomerStore{[]Customer{{Id: 1, Name: "mary", Balance: 10000}, {Id: 2, Name: "adam", Balance: 20000}}}
		server := NewCustomerServer(store)
		server.ServeHTTP(response, request)
		assertStatusCode(t, response.Code, NotFoundStatus)
	})

	t.Run("Should return status ok when getting list of all customers", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
		response := httptest.NewRecorder()

		store := &StubCustomerStore{[]Customer{{Id: 1, Name: "mary", Balance: 10000}, {Id: 2, Name: "adam", Balance: 20000}}}
		server := NewCustomerServer(store)
		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)

	})

	t.Run("Should return a list of all customers", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
		response := httptest.NewRecorder()

		store := &StubCustomerStore{[]Customer{{Id: 2, Name: "adam", Balance: 20000}, {Id: 1, Name: "mary", Balance: 10000}}}
		server := NewCustomerServer(store)
		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		want := []Customer{{Id: 2, Name: "adam", Balance: 20000}, {Id: 1, Name: "mary", Balance: 10000}}

		var got []Customer
		err := json.NewDecoder(response.Body).Decode(&got)
		assertJsonError(t, err)
		AssertAllCustomerResponse(t, got, want)

	})
	t.Run("Should return a list of customers sorted by balance", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
		response := httptest.NewRecorder()

		store := &StubCustomerStore{[]Customer{{Id: 1, Name: "mary", Balance: 10000}, {Id: 2, Name: "adam", Balance: 20000}}}
		server := NewCustomerServer(store)
		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		want := []Customer{{Id: 2, Name: "adam", Balance: 20000}, {Id: 1, Name: "mary", Balance: 10000}}

		var got []Customer
		err := json.NewDecoder(response.Body).Decode(&got)
		assertJsonError(t, err)
		AssertAllCustomerResponse(t, got, want)
	})
}

func TestPOSTCustomer(t *testing.T) {
	t.Run("Should return accepted on POST", func(t *testing.T) {
		customerObj := Customer{Name: "Tom", Balance: 0, Id: 3}
		marshalledJson, err := json.Marshal(customerObj)
		assertJsonMarshalError(t, err, customerObj)

		request, _ := NewPOSTCustomerRequest(marshalledJson)
		response := httptest.NewRecorder()

		store := &StubCustomerStore{[]Customer{}}
		server := NewCustomerServer(store)

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, StatusAccepted)
	})
	t.Run("Should add a new customer", func(t *testing.T) {
		want := Customer{Name: "Tom", Balance: 0, Id: 3}
		marshalledJson, err := json.Marshal(want)
		assertJsonMarshalError(t, err, want)

		request, _ := NewPOSTCustomerRequest(marshalledJson)
		response := httptest.NewRecorder()

		store := &StubCustomerStore{[]Customer{}}
		server := NewCustomerServer(store)
		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, StatusAccepted)

		var got Customer
		err = json.NewDecoder(response.Body).Decode(&got)
		assertJsonError(t, err)
		assertCustomerResponse(t, got, want)
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
	customers []Customer
}

func (s *StubCustomerStore) GetCustomerBalance(name string) float64 {
	for _, c := range s.customers {
		if c.Name == name {
			return c.Balance
		}
	}
	return 0
}

func (s *StubCustomerStore) RecordNewCustomer(c Customer) {
	s.customers = append(s.customers, c)
}

func assertStatusCode(t testing.TB, got int, want int) {
	if got != want {
		t.Errorf("Wrong status code, got %v expected %v", got, want)
	}
}

func (s *StubCustomerStore) GetCustomers() Customers {
	sort.Slice(s.customers, func(i, j int) bool {
		return s.customers[i].Balance > s.customers[j].Balance
	})

	return s.customers
}

func assertCustomerResponse(t testing.TB, got, want Customer) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wrong customer returned: Got %v Expected: %v", got, want)
	}
}

func assertJsonError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Problem decoding response.body: %v", err)
	}
}

func AssertAllCustomerResponse(t testing.TB, got, want []Customer) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wrong customer array returned: Got %v, want %v", got, want)
	}
}

func assertJsonMarshalError(t testing.TB, err error, customer Customer) {
	if err != nil {
		t.Errorf("Problem marshalling json %v, %v", customer, err)
	}
}

func NewPOSTCustomerRequest(c []byte) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, "/customer/", bytes.NewReader(c))
	return request, err
}
