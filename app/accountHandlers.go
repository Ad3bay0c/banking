package app

import (
	"encoding/json"
	"github.com/Ad3bay0c/banking/dto"
	"github.com/Ad3bay0c/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	Service service.AccountService
}

func (a AccountHandler) newAccount(w http.ResponseWriter, req *http.Request) {
	var request dto.AccountRequest

	vars := mux.Vars(req)
	customerId := vars["customer_id"]

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, nil, nil, err.Error())
		return
	}
	request.CustomerID = customerId
	response, respErr := a.Service.NewAccount(request)
	if respErr != nil {
		WriteResponse(w, respErr.Code, nil, nil, respErr.Message)
		return
	}

	WriteResponse(w, http.StatusOK, response, "Successful", nil)
}
