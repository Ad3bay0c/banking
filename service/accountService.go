package service

import (
	"github.com/Ad3bay0c/banking/domain/account"
	"github.com/Ad3bay0c/banking/dto"
	"github.com/Ad3bay0c/banking/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.AccountRequest) (*dto.AccountResponse, *errs.AppError)
	SaveTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo account.Repository
}

func (d DefaultAccountService) NewAccount(req dto.AccountRequest) (*dto.AccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	acc := account.Account{
		ID:          "",
		CustomerID:  req.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAcc, err := d.repo.Save(acc)
	if err != nil {
		return nil, err
	}
	response := newAcc.ToNewAccountResponseDto()
	return &response, nil
}
func NewAccountService(repo account.Repository) DefaultAccountService {
	return DefaultAccountService{
		repo: repo,
	}
}
