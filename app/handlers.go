package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Create(w http.ResponseWriter, req *http.Request) {

}

func Greet(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	fmt.Fprintf(w, "%v", params["id"])
}
