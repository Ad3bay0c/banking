package app

import (
	"encoding/json"
	"github.com/Ad3bay0c/banking/service"
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

