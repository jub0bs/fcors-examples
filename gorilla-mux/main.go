package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jub0bs/fcors"
)

func main() {
	// create a router
	router := mux.NewRouter()

	// configure the CORS middleware
	cors, err := fcors.AllowAccess(
		fcors.FromOrigins("https://example.com"),
		fcors.WithMethods(
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		),
		fcors.WithRequestHeaders("Authorization"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// apply the CORS middleware to the router
	handler := cors(router)

	// register the hello handler for the /hello route
	router.HandleFunc("/hello", hello)

	// start the server on port 8080
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"hello":"world"}`)
}
