package main

type InMemoryLocationStore struct {
	locations map[int]Location
}

func (i *InMemoryLocationStore) GetLocation(id int) Location {
	return i.locations[id]
}

func (i *InMemoryLocationStore) AddLocation(location Location) {
	i.locations[location.Id] = location
}
