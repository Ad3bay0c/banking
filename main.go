package main

import (
	"github.com/Ad3bay0c/banking/app"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile | log.Llongfile | log.Ldate | log.Ltime)
	if err := godotenv.Load(); err != nil {
		log.Printf("Error reading .env file\n")
	}
	app.Start()
}
