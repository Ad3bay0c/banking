package domain

import (
	"database/sql"
	"fmt"
	"github.com/Ad3bay0c/banking/errs"
	"github.com/Ad3bay0c/banking/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

const (
	TABLE = "customers"
)
type CustomerRepositoryDb struct {
	db *sqlx.DB
}
func NewCustomerRepositoryDB() *CustomerRepositoryDb {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s  dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"))
	db, err := sqlx.Open("postgres", dns)


	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		panic(err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	log.Println("Database Connected")

	return &CustomerRepositoryDb{
		db: db,
	}
}
func (c *CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError)  {
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
		}else {
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
		logger.Error("Error: "+err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			return nil, errs.NewUnexpectedError("unexpected database errs")
		}
    }
	return &customer, nil
}