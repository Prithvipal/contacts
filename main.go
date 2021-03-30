package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var m = make(map[int]contact)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/contants", getContantsHandler).Methods("GET")
	r.HandleFunc("/api/v1/contants", postContantsHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func getContantsHandler(w http.ResponseWriter, r *http.Request) {
	contList := make([]contact, 0)
	for _, cont := range m {
		contList = append(contList, cont)
	}
	data, _ := json.Marshal(contList)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func postContantsHandler(w http.ResponseWriter, r *http.Request) {
	var cont contact
	err := json.NewDecoder(r.Body).Decode(&cont)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := len(m) + 1
	cont.Id = id
	m[id] = cont
}

type contact struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	PhoneNumber []int    `json:"phone_number"`
	Owners      []string `json:"owners"`
}
