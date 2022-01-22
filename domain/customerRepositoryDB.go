package domain

import (
	"database/sql"
	"fmt"
	"github.com/Ad3bay0c/banking/errs"
	"github.com/Ad3bay0c/banking/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	TABLE = "customers"
)

type CustomerRepositoryDb struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(DbClient *sqlx.DB) *CustomerRepositoryDb {
	return &CustomerRepositoryDb{
		db: DbClient,
	}
}

func (c *CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var customers []Customer
	var query string
	if status == "1" {
		query = fmt.Sprintf("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM %s WHERE status = $1", TABLE)
	} else if status == "0" {
		query = fmt.Sprintf("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM %s WHERE status = $1", TABLE)
	} else {
		status = "0"
		query = fmt.Sprintf("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM %s WHERE status >= $1", TABLE)
	}
	//rows, err := c.db.Query(query, status)
	err := c.db.Select(&customers, query, status)
	if err != nil {
		log.Print(err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return customers, nil
}

func (c *CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	var customer Customer
	query := fmt.Sprintf("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM %s WHERE customer_id = $1", TABLE)
	err := c.db.Get(&customer, query, id)
	if err != nil {
		logger.Error("Error: " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			return nil, errs.NewUnexpectedError("unexpected database errs")
		}
	}
	return &customer, nil
}
