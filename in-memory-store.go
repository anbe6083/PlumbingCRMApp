package main

type InMemoryLocationStore struct {
}

func (i *InMemoryLocationStore) GetLocation(id int) string {
	return "location name"
}
