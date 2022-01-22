package account

import (
	"github.com/Ad3bay0c/banking/dto"
	"github.com/Ad3bay0c/banking/errs"
)

type Account struct {
	ID          string `db:"account_id"`
	CustomerID  string	`db:"customer_id, omitempty"`
	OpeningDate string	`db:"opening_date, omitempty"`
	AccountType string	`db:"account_type, omitempty"`
	Amount      float64	`db:"amount, omitempty"`
	Status      string	`db:"status, omitempty"`
}

type Repository interface {
	Save(account Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(id string) (*Account, *errs.AppError)
}

func (acc Account) ToNewAccountResponseDto() dto.AccountResponse {
	return dto.AccountResponse{
		AccountID: acc.ID,
	}
}

func (acc *Account) CanWithdraw(amount float64) bool {
	return acc.Amount >= amount
}


