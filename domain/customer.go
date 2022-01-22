package domain

import (
	"github.com/Ad3bay0c/banking/dto"
	"github.com/Ad3bay0c/banking/errs"
)

type Customer struct {
	ID      string `json:"customer_id" db:"customer_id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Dob     string `json:"date_of_birth" db:"date_of_birth"`
	Status  string `json:"status"`
}

type CustomerRepository interface {
	// FindAll status can be 1: active, 2: inactive
	FindAll(status string) ([]Customer, *errs.AppError)
	ByID(id string) (*Customer, *errs.AppError)
}

func (c Customer) ToDto() *dto.CustomerResponse {
	statusText := "active"
	if c.Status == "0" {
		statusText = "inactive"
	}
	return &dto.CustomerResponse{
		ID:      c.ID,
		Name:    c.Name,
		City:    c.City,
		Zipcode: c.Zipcode,
		Dob:     c.Dob,
		Status:  statusText,
	}
}
