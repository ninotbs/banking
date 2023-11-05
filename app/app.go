package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func getCustomer(writter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprint(writter, vars["customer_id"])
}

func createCustomer(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writter, "Post request received")
}
