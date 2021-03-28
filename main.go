package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/hello", helloHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	c := contact{
		Name:        "Prithvi",
		Address:     "Pune",
		PhoneNumber: []int{19292, 3040404},
		Owners:      []string{"Prithvipal", "Prithviraj"},
	}
	data, _ := json.Marshal(c)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	// w.Write([]byte("Hello Prithvi!!! How are you??"))
}

type contact struct {
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	PhoneNumber []int    `json:"phone_number"`
	Owners      []string `json:"owners"`
}
