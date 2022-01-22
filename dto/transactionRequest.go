package dto

import (
	"github.com/Ad3bay0c/banking/errs"
	"strings"
)

const (
	WITHDRAW = "withdrawal"
	DEPOSIT = "deposit"
)
type TransactionRequest struct {
	TransactionID        string `json:"id"`
	AccountID			 string `json:"account_id"`
	Amount				 float64 `json:"amount"`
	TransactionType		 string `json:"transaction_type"`
	TransactionDate      string `json:"transaction_date"`
}

func (req TransactionRequest) Validate() *errs.AppError {
	if strings.ToLower(req.TransactionType) != WITHDRAW && strings.ToLower(req.TransactionType) != DEPOSIT {
		return errs.NewValidationError("transaction type can either by withdrawal or deposit")
	}
	return nil
}

func (req TransactionRequest) IsTransactionWithdrawal() bool {
	return req.TransactionType == WITHDRAW
}