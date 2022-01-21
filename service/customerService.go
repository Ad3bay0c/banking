package service

import (
	"github.com/Ad3bay0c/banking/domain"
	"github.com/Ad3bay0c/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomerByID(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService  {
	return DefaultCustomerService{repo: repository}
}

func (d DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)  {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return d.repo.FindAll(status)
}

func (d DefaultCustomerService) GetCustomerByID(id string) (*domain.Customer, *errs.AppError)  {
	return d.repo.ByID(id)
}