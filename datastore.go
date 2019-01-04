package main

type DataStore interface {
	Put(id, value string)
	Get(id string) string
}

type InMemoryDataStore struct {
	data map[string]string
}

func (self *InMemoryDataStore) Put(id, value string) {
	self.data[id] = value
}

func (self *InMemoryDataStore) Get(id string) string {
	result := self.data[id]
	return result
}

func NewInMemoryDataStore() DataStore {
	return &InMemoryDataStore{data: make(map[string]string)}
}
