package main

type Customer struct {
	id                int
	createdOn         string
	billingCustomer   string
	locations         []Location
	billingCustomerId int
	address           Address
}
