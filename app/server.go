package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// defineRoutes defines the routes for the application
func defineRoutes(router *mux.Router, ch *CustomerHandlers) {
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomerByID).Methods(http.MethodGet)
}

func Start(ch *CustomerHandlers) {
	// initialize the router
	router := mux.NewRouter()

	defineRoutes(router, ch)

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PORT == ":" {
		PORT = ":8080"
	}

	s := &http.Server{
		Handler: router,
		Addr:    PORT,
	}
	go func() {
		log.Print("Server Started at http://localhost:8080")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf(err.Error())
		}
	}() // start the server

	wait := make(chan os.Signal) // create a channel to wait for a signal

	signal.Notify(wait, os.Interrupt) // register the channel to be notified on an interrupt (Ctrl+C)
	<-wait
	log.Printf("Server shutting down...")

	time.Sleep(2 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// shut down gracefully, but wait no longer than 3 seconds before halting
	if err := s.Shutdown(ctx); err != nil {
		log.Printf(err.Error())
	}
	log.Printf("Server exits sucessfully")
}

func GetDBClient() *sqlx.DB {
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
	return db
}
