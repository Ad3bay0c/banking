package dto

import (
	"github.com/Ad3bay0c/banking/errs"
	"strings"
)

type AccountRequest struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (acc AccountRequest) Validate() *errs.AppError {
	if acc.Amount < 5000 {
		return errs.NewValidationError("Amount should be greater than 5000")
	}
	if strings.ToLower(acc.AccountType) != "saving" && strings.ToLower(acc.AccountType) != "current" {
		return errs.NewValidationError("Account type should be either saving or current")
	}
	return nil
}