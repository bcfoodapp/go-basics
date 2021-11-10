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

	fmt.Println("Open http://localhost:8000/ in your web browser to see the output from the server!")

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func addRoutes(router *mux.Router) {
	api := NewAPI()

	type route struct {
		path    string
		handler func(http.ResponseWriter, *http.Request)
	}

	// Define paths and their handler functions
	// Path is relative to the root path.
	// So if you want to access /vendors/, the URL is http://localhost:8000/vendors/
	routes := [...]route{
		{"/", root},
		{"/vendors/", api.vendors},
		// {id} matches any string. You can get the string that was in {id} with mux.Vars().
		{"/vendors/{id}", api.vendor},
	}

	// For each route, register the path and its handler. We set them to only respond to GET
	// requests.
	for _, route := range routes {
		router.HandleFunc(route.path, route.handler).Methods(http.MethodGet)
	}
}

// API stores any server state such as databases.
type API struct {
	// In reality, you would store a database connection here
	Vendors []Vendor
}

func NewAPI() *API {
	vendors := make([]Vendor, 1)
	vendors[0] = Vendor{
		ID:   "cycledogs",
		Name: "Cycle Dogs",
	}
	return &API{vendors}
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
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		// Encoding should not fail unless there is a bug
		log.Panic(err)
	}
}

// vendors is the handler for /vendors/. It returns an array of all Vendors.
func (b *API) vendors(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, &b.Vendors)
}

// vendor is the handler for /vendors/{id}. It returns a vendor if given id matches a vendor, or
// otherwise returns StatusNotFound (404).
func (b *API) vendor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"] // Retrieves id value from path
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
