package service

import "github.com/Ad3bay0c/banking/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerByID(id string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error)  {
	return d.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService  {
	return DefaultCustomerService{repo: repository}
}

func (d DefaultCustomerService) GetCustomerByID(id string) (*domain.Customer, error)  {
	return d.repo.ByID(id)
}