package main

import (
	//"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Prithvipal/phone-dir/controller"
	"github.com/Prithvipal/phone-dir/dto"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/contants", getContantsHandler).Methods("GET")
	r.HandleFunc("/api/v1/contants", controller.PostContantsHandler).Methods("POST")
	r.HandleFunc("/api/v1/contants/{id}", deleteContantsHandler).Methods("DELETE")
	r.HandleFunc("/api/v1/contants/{id}", getByIdContantsHandler).Methods("GET")
	r.HandleFunc("/api/v1/contants/{id}", putContantsHandler).Methods("PUT")
	http.ListenAndServe(":8080", r)
}

func getByIdContantsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//contantID, err := strconv.Atoi(id)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://prithvi:prithvi123@cluster0.rmlet.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = client.Connect(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(r.Context())

	contactDB := client.Database("contact")
	companyColl := contactDB.Collection("companies")

	var cont bson.M
	if err := companyColl.FindOne(r.Context(), bson.M{"_id": objID}).Decode(&cont); err != nil {
		log.Println(err)
	}

	data, _ := json.Marshal(cont)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func putContantsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var cont dto.Contact
	errr := json.NewDecoder(r.Body).Decode(&cont)
	if errr != nil {
		log.Println("Could not parse request payload", errr.Error())
		http.Error(w, errr.Error(), http.StatusBadRequest)
		return
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://prithvi:prithvi123@cluster0.rmlet.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = client.Connect(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(r.Context())

	contactDB := client.Database("contact")
	companyColl := contactDB.Collection("companies")
	update := bson.M{
		"$set": cont,
	}
	upsert := true
	opt := options.FindOneAndUpdateOptions{Upsert: &upsert}
	result := companyColl.FindOneAndUpdate(r.Context(), bson.M{"_id": objID}, update, &opt)
	if result.Err() != nil {
		log.Println(result.Err())
	}

}

func getContantsHandler(w http.ResponseWriter, r *http.Request) {
	//contList := make([]contact, 0)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://prithvi:prithvi123@cluster0.rmlet.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = client.Connect(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(r.Context())

	contactDB := client.Database("contact")
	companyColl := contactDB.Collection("companies")

	cursor, err := companyColl.Find(r.Context(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(r.Context())
	var cont []bson.M
	if err = cursor.All(r.Context(), &cont); err != nil {
		log.Fatal(err)
	}
	data, _ := json.Marshal(cont)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func deleteContantsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://prithvi:prithvi123@cluster0.rmlet.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = client.Connect(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(r.Context())

	contactDB := client.Database("contact")
	companyColl := contactDB.Collection("companies")

	result, err := companyColl.DeleteOne(r.Context(), bson.M{"_id": objID})
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

}
