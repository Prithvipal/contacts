package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Prithvipal/phone-dir/dto"
	"github.com/Prithvipal/phone-dir/service"
	"github.com/gorilla/mux"
)

func PostContantsHandler(w http.ResponseWriter, r *http.Request) {
	var cont dto.Contact
	err := json.NewDecoder(r.Body).Decode(&cont)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.SaveContact(r.Context(), cont)

	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetContantsHandler(w http.ResponseWriter, r *http.Request) {

	searchParam := r.URL.Query().Get("search")
	fmt.Println(searchParam)
	data, err := service.GetContact(r.Context(), searchParam)

	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, data)
}

func PutContantsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var cont dto.Contact
	err := json.NewDecoder(r.Body).Decode(&cont)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.PutContact(r.Context(), id, cont)

	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}
func DeleteContantsHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	err := service.DeleteContact(r.Context(), id)

	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func GetByIdContantsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, err := service.GetByIdContantsHandler(r.Context(), id)
	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, data)
}

func writeJSON(w http.ResponseWriter, records interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(records)
	if err != nil {
		log.Println("Error while getting TODO List", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
}
