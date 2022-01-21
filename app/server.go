package app

import (
	"context"
	"fmt"
	"github.com/Ad3bay0c/banking/domain"
	"github.com/Ad3bay0c/banking/service"
	"github.com/gorilla/mux"
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


func Start() {
	// initialize the router
	router := mux.NewRouter()

	ch := &CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	defineRoutes(router, ch)

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PORT == ":" {
		PORT = ":8080"
	}

	s := &http.Server{
		Handler: router,
		Addr: PORT,
	}
	go func() {
		log.Print("Server Started at http://localhost:8080")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf(err.Error())
		}
	}()
	wait := make(chan os.Signal)

	signal.Notify(wait, os.Interrupt)
	<-wait
	log.Printf("Server shutting down...")

	time.Sleep(2 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
        log.Printf(err.Error())
    }
	log.Printf("Server exits sucessfully")
}
