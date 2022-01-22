package main

import (
	"github.com/Ad3bay0c/banking/app"
	"github.com/Ad3bay0c/banking/domain"
	"github.com/Ad3bay0c/banking/service"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile | log.Llongfile | log.Ldate | log.Ltime)
	if err := godotenv.Load(); err != nil {
		log.Printf("Error reading .env file\n")
	}
	db := app.GetDBClient()
	customerRepository := domain.NewCustomerRepositoryDB(db)
	//accountRepository := account.NewRepositoryDB(db)
	ch := &app.CustomerHandlers{Service: service.NewCustomerService(customerRepository)}
	app.Start(ch)
}
