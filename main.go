package main

import (
	//"context"

	"net/http"

	"github.com/Prithvipal/phone-dir/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/contants", controller.GetContantsHandler).Methods("GET")
	r.HandleFunc("/api/v1/contants", controller.PostContantsHandler).Methods("POST")
	r.HandleFunc("/api/v1/contants/{id}", controller.DeleteContantsHandler).Methods("DELETE")
	r.HandleFunc("/api/v1/contants/{id}", controller.GetByIdContantsHandler).Methods("GET")
	r.HandleFunc("/api/v1/contants/{id}", controller.PutContantsHandler).Methods("PUT")
	http.ListenAndServe(":8080", r)
}
