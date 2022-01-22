package app

import (
	"encoding/json"
	"github.com/Ad3bay0c/banking/logger"
	"github.com/Ad3bay0c/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandlers struct {
	Service service.CustomerService
}

func (c *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status := r.URL.Query().Get("status")
	customers, err := c.Service.GetAllCustomers(status)
	if err != nil {
		logger.Error("Error: " + err.Message)
		WriteResponse(w, err.Code, err.Message)
		return
	}
	WriteResponse(w, http.StatusOK, customers)
}

func (c *CustomerHandlers) getCustomerByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	customer, err := c.Service.GetCustomerByID(params["customer_id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		logger.Error(err.Message)
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
