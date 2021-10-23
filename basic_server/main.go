package main

// This demonstrates how a REST API server in Go might look like.

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	// The mux library is a request router which maps a URL to a handler function
	router := mux.NewRouter()
	router.StrictSlash(true)

	addRoutes(router)

	server := http.Server{
		// The server listens on this address
		Addr:    ":8000",
		Handler: router,
		// Timeout durations
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	// Don't forget to defer the Close
	defer server.Close()

	fmt.Println("Go to http://localhost:8000/ to see the output from the server!")

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func addRoutes(router *mux.Router) {
	backend := NewBackend()

	type route struct {
		path    string
		handler func(http.ResponseWriter, *http.Request)
	}

	// Define paths and their handler functions
	routes := [...]route{
		{"/", root},
		{"/vendors/", backend.vendors},
		{"/vendors/{id}", backend.vendor},
	}

	// For each route, register the path and its handler. We set them to only respond to GET
	// requests.
	for _, route := range routes {
		router.HandleFunc(route.path, route.handler).Methods(http.MethodGet)
	}
}

// Backend stores any server state such as databases.
type Backend struct {
	// In reality, you would store a database connection here
	Vendors []Vendor
}

func NewBackend() *Backend {
	vendors := make([]Vendor, 1)
	vendors[0] = Vendor{
		ID:   "cycledogs",
		Name: "Cycle Dogs",
	}
	return &Backend{vendors}
}

func root(w http.ResponseWriter, _ *http.Request) {
	// The simplest way to write a response is to call Write() on the ResponseWriter.
	w.Write([]byte("You've called the root route!\n"))
}

// Vendor is a hypothetical record type.
type Vendor struct {
	ID   string
	Name string
}

// writeJSON adds the correct Content-Type header to the response and writes the payload in JSON
// encoding. Notice how payload is the empty interface type. All types implement an empty interface.
func writeJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")

	// Since ResponseWriter is a Writer, we can pass the ResponseWriter to NewEncoder.
	// We check for errors in encoding. If an error occurred, we write the error into the response.
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
}

// vendors is the handler for /vendors/. It returns an array of all Vendors.
func (b *Backend) vendors(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, &b.Vendors)
}

// vendor is the handler for /vendors/{id}. It returns a vendor if given id matches a vendor, or
// otherwise returns StatusNotFound (404).
func (b *Backend) vendor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, vendor := range b.Vendors {
		if vendor.ID == id {
			writeJSON(w, &vendor)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

/*
https://gobyexample.com/http-servers
https://github.com/gorilla/mux
https://gobyexample.com/json
*/
