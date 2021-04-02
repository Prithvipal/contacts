package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var m = make(map[int]contact)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/contants", getContantsHandler).Methods("GET")
	r.HandleFunc("/api/v1/contants", postContantsHandler).Methods("POST")
	r.HandleFunc("/api/v1/contants/{id}", deleteContantsHandler).Methods("DELETE")
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

func deleteContantsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	contantID, err := strconv.Atoi(id)
	if err != nil {
		errMsg := fmt.Sprintf("Could not convert value of id to integer: %v. Are passing non int value?", id)
		log.Println(errMsg, err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	_, ok := m[contantID]
	if !ok {
		errMsg := fmt.Sprintf("Could not found contant: %v. Is it already delete?", id)
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusGone)
		return
	}
	delete(m, contantID)

}

type contact struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	PhoneNumber []int    `json:"phone_number"`
	Owners      []string `json:"owners"`
}
