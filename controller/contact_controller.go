package controller

import (
	"encoding/json"
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

	data, err := service.GetContact(r.Context())

	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func PutContantsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var cont dto.Contact
	err := json.NewDecoder(r.Body).Decode(&cont)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.PutContact(r.Context(), cont, id)

	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func DeleteContantsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	err := service.DeleteContact(r.Context(), id)

	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetByIdContantsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	data, err := service.GetByIdContantsHandler(r.Context(), id)
	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
