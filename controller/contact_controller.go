package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Prithvipal/phone-dir/dto"
	"github.com/Prithvipal/phone-dir/service"
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
