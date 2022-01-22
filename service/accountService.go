package service

import (
	"github.com/Ad3bay0c/banking/domain/account"
	"github.com/Ad3bay0c/banking/errs"
)

type AccountService interface {
	NewAccount() *errs.AppError
}

type DefaultAccountService struct {
	repo account.Repository
}

