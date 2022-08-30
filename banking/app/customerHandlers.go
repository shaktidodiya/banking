package app

import (
	"Desktop/golang/banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// handler function
//func greetHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello World!")
//}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{
	//		Name:    "Shakti",
	//		City:    "Blr",
	//		ZipCode: "1234",
	//	},
	//	{
	//		Name:    "Shyam",
	//		City:    "Blr",
	//		ZipCode: "1234",
	//	},
	//}

	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

	//if r.Header.Get("Content-Type") == "application/xml" {
	//	w.Header().Add("Content-type", "application/xml")
	//	xml.NewEncoder(w).Encode(customers)
	//} else {
	//	w.Header().Add("Content-type", "application/json")
	//	json.NewEncoder(w).Encode(customers)
	//}

}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func (ch *CustomerHandlers) getByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status") //

	log.Println(status)
	customers, err := ch.service.GetAllByStatus(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil{
		panic(err)
	}
}
