package main

type Customer struct {
	Id      int
	Balance float64
	Name    string
}

type Customers []Customer

func (c Customers) Find(name string) *Customer {
	for i, customer := range c {
		if customer.Name == name {
			return &c[i]
		}
	}
	return nil
}
