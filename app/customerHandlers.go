package app

import (
	"encoding/json"
	"github.com/Ad3bay0c/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}
func (c *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customers, _ := c.service.GetAllCustomers()
	json.NewEncoder(w).Encode(customers)
}

func (c *CustomerHandlers) getCustomerByID(w http.ResponseWriter, req *http.Request)  {
	params := mux.Vars(req)
	customer, err := c.service.GetCustomerByID(params["customer_id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	json.NewEncoder(w).Encode(customer)
}
