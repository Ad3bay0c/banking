package app

import (
	"encoding/json"
	"github.com/Ad3bay0c/banking/logger"
	"github.com/Ad3bay0c/banking/service"
	"github.com/gorilla/mux"
	"net/http"
	"time"
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
		WriteResponse(w, err.Code, nil, nil, err.Message)
		return
	}

	WriteResponse(w, http.StatusOK, customers, "Successful", nil)
}

func (c *CustomerHandlers) getCustomerByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	customer, err := c.Service.GetCustomerByID(params["customer_id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		logger.Error(err.Message)
		WriteResponse(w, err.Code, nil, nil, err.Message)
		return
	}

	WriteResponse(w, http.StatusOK, customer, "Successful", nil)
}

func WriteResponse(w http.ResponseWriter, code int, data, success, errorMessage interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	message := map[string]interface{}{
		"error": errorMessage,
		"data": data,
		"message": success,
		"date": time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}
