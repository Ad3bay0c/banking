package domain

import "github.com/Ad3bay0c/banking/errs"

type Customer struct {
	ID      string
	Name    string
	City    string
	Zipcode string
	Dob     string
	Status  string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ByID(id string) (*Customer, *errs.AppError)
}