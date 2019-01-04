package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type TestContext struct {
	dataStore DataStore
	router    *mux.Router
}

func newTestContext() *TestContext {
	serverId := "test"
	dataStore := NewInMemoryDataStore()
	return &TestContext{dataStore: NewInMemoryDataStore(), router: createRouter(serverId, dataStore)}
}

func testRequest(c *TestContext, req *http.Request, expectedBody string, t *testing.T) {
	rr := httptest.NewRecorder()
	c.router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedBody)
	}
}

func TestHello(t *testing.T) {
	c := newTestContext()
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	testRequest(c, req, "Hello World", t)
}

func TestRoot(t *testing.T) {
	c := newTestContext()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	testRequest(c, req, "Id: test, timestamp: 0", t)
}

func TestEcho(t *testing.T) {
	c := newTestContext()
	req, err := http.NewRequest("POST", "/echo", bytes.NewBuffer([]byte(`data`)))
	if err != nil {
		t.Fatal(err)
	}

	testRequest(c, req, `{"body":"data"}`, t)
}

func TestPutGetRequest(t *testing.T) {
	c := newTestContext()
	putRequest1, err := http.NewRequest("POST", "/key", bytes.NewBuffer([]byte(`{"hello": "bello"}`)))
	if err != nil {
		t.Fatal(err)
	}
	putRequest2, err := http.NewRequest("POST", "/key", bytes.NewBuffer([]byte(`{"hello": "csa"}`)))
	if err != nil {
		t.Fatal(err)
	}
	getRequest, err := http.NewRequest("GET", "/key/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	testRequest(c, putRequest1, `{"hello": "bello"}`, t)
	testRequest(c, getRequest, `{"hello": "bello"}`, t)
	testRequest(c, putRequest2, `{"hello": "csa"}`, t)
	testRequest(c, getRequest, `{"hello": "csa"}`, t)
}
