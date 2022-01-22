package account

import (
	"github.com/Ad3bay0c/banking/dto"
	"github.com/Ad3bay0c/banking/errs"
)

type Account struct {
	ID          string `db:"account_id"`
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type Repository interface {
	Save(account Account) (*Account, *errs.AppError)
}

func (acc Account) ToNewAccountResponseDto() dto.AccountResponse {
	return dto.AccountResponse{
		AccountID: acc.ID,
	}
}
