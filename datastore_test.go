package main

import "testing"

func createDataStore() DataStore {
	return NewInMemoryDataStore()
}

func TestPutGet(t *testing.T) {
	ds := createDataStore()

	ds.Put("hello", "bello")

	result, err := ds.Get("hello")
	if result != "bello" || err != nil {
		t.Errorf("result: '%v', expected: '%v'", result, "bello")
	}
}

func TestEmpty(t *testing.T) {
	ds := createDataStore()

	result, err := ds.Get("hello")
	if result != "" || err != nil {
		t.Errorf("result: '%v', expected: '%v'", result, "")
	}
}

func TestAddGetModify(t *testing.T) {
	ds := createDataStore()

	ds.Put("hello", "bello")
	ds.Put("hello", "csa")

	result, err := ds.Get("hello")
	if result != "csa" || err != nil {
		t.Errorf("result: '%v', expected: '%v'", result, "csa")
	}
}
