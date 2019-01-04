package main

import "testing"

func createDataStore() DataStore {
	return NewInMemoryDataStore()
}

func TestPutGet(t *testing.T) {
	ds := createDataStore()

	ds.Put("hello", "bello")

	result := ds.Get("hello")
	if result != "bello" {
		t.Errorf("result: '%v', expected: '%v'", result, "bello")
	}
}

func TestEmpty(t *testing.T) {
	ds := createDataStore()

	result := ds.Get("hello")
	if result != "" {
		t.Errorf("result: '%v', expected: '%v'", result, "")
	}
}

func TestAddGetModify(t *testing.T) {

	ds := createDataStore()

	ds.Put("hello", "bello")
	ds.Put("hello", "csa")

	result := ds.Get("hello")
	if result != "csa" {
		t.Errorf("result: '%v', expected: '%v'", result, "csa")
	}
}
