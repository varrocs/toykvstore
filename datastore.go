package main

type DataStore interface {
	Put(id, value string) error
	Get(id string) (string, error)
}

type InMemoryDataStore struct {
	data map[string]string
}

func (self *InMemoryDataStore) Put(id, value string) error {
	self.data[id] = value
	return nil
}

func (self *InMemoryDataStore) Get(id string) (string, error) {
	result := self.data[id]
	return result, nil
}

func NewInMemoryDataStore() DataStore {
	return &InMemoryDataStore{data: make(map[string]string)}
}
