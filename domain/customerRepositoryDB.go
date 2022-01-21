package domain

import (
	"database/sql"
	"fmt"
	"github.com/Ad3bay0c/banking/errs"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

const (
	TABLE = "customers"
)
type CustomerRepositoryDb struct {
	db *sql.DB
}
func NewCustomerRepositoryDB() *CustomerRepositoryDb {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s  dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", dns)


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
func (c *CustomerRepositoryDb) FindAll() ([]Customer, error)  {
	var customers []Customer
	rows, err := c.db.Query(fmt.Sprintf("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM %s", TABLE))

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.City, &customer.Zipcode, &customer.Dob, &customer.Status)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	var customer Customer
	row := c.db.QueryRow(fmt.Sprintf("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM %s WHERE customer_id = $1", TABLE), id)

	err := row.Scan(&customer.ID, &customer.Name, &customer.City, &customer.Zipcode, &customer.Dob, &customer.Status)
	if err != nil {
		log.Printf("Error: %v", err)
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			return nil, errs.NewUnexpectedError("unexpected database errs")
		}
    }
	return &customer, nil
}