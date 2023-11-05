package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ninotbs/banking/domain"
	"github.com/ninotbs/banking/service"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)


	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
