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

func (i *InMemoryLocationStore) GetLocations() []Location {
	var locations []Location

	for _, val := range i.locations {
		locations = append(locations, val)
	}

	return locations
}
