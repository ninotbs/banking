package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"github.com/ninotbs/banking/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}


func (ch *CustomerHandlers) getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	//customers := []Customer{
	//	{Name: "John", City: "New York", Zipcode: "123"},
	//}

	customers, _ := ch.service.GetAllCustomers()

	if request.Header.Get("Content-Type") == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(writer).Encode(customers)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(customers)
	}
}
