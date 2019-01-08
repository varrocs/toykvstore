package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type EchoResponse struct {
	Body string `json:"body"`
}

func createRootHandler(serverId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[ ] root handler for server %v\n", serverId)
		fmt.Fprintf(w, "Id: %v, timestamp: %d", serverId, 0)
	}
}

func createGetHandler(dataStore DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["key"]
		fmt.Printf("[ ] get called for id: '%v'\n", key)
		value, err := dataStore.Get(key)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, `{"%s": "%s"}`, key, value)
	}
}

func createPutHandler(dataStore DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		valueBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var data map[string]string
		err = json.Unmarshal(valueBytes, &data)
		if err != nil {
			panic(err)
		}
		for key, value := range data {
			fmt.Printf("[ ] put called for id: %v, value %v\n", key, value)
			err = dataStore.Put(key, value)
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, `{"%s": "%s"}`, key, value)
			break
		}
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[ ] hello handler")
	fmt.Fprintf(w, "Hello World")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[ ] echo handler")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	responseStruct := EchoResponse{Body: string(reqBody)}
	responseJson, err := json.Marshal(responseStruct)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(responseJson)
}

func createRouter(serverId string, dataStore DataStore) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/echo", echoHandler)
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/key/{key}", createGetHandler(dataStore)).Methods("GET")
	r.HandleFunc("/key", createPutHandler(dataStore)).Methods("POST")
	r.HandleFunc("/", createRootHandler(serverId))
	return r
}

func main() {
	serverId := "12"
	redisAddress := GetRedisAddress()
	dataStore, err := NewRedisDataStore(GetRedisAddress())
	if err == nil {
		fmt.Println("Starting with redis store: ", redisAddress)
	} else {
		fmt.Println("Failed to start with redis on :", redisAddress, ", error:", err, "Starting in memory DB")
		dataStore = NewInMemoryDataStore()
	}

	r := createRouter(serverId, dataStore)

	listenAddress := ":8080"
	fmt.Printf("Starting %v on address %v\n", serverId, listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, r))
}
