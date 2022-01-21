package app

import (
	"encoding/json"
	"github.com/Ad3bay0c/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CustomerHandlers struct {
	Service service.CustomerService
}

func (c *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customers, _ := c.Service.GetAllCustomers()
	json.NewEncoder(w).Encode(customers)
}

func (c *CustomerHandlers) getCustomerByID(w http.ResponseWriter, req *http.Request)  {
	params := mux.Vars(req)
	customer, err := c.Service.GetCustomerByID(params["customer_id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err.Message)
		WriteResponse(w, err.Code, err.Message)
		return
	}

	WriteResponse(w, http.StatusOK, customer)
}

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}