package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func getOnePerson(w http.ResponseWriter, r *http.Request) {
}
func createPerson(w http.ResponseWriter, r *http.Request) {
}
func deletePerson(w http.ResponseWriter, r *http.Request) {

}
func main() {
	r := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "john", Lastname: "doo", Address: &Address{City: "Dublin", State: "CA"}})
	people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "DOO"})
	r.HandleFunc("/people", getPeople).Methods("GET")
	r.HandleFunc("/people/{id}", getOnePerson).Methods("GET")
	r.HandleFunc("/people/{id}", createPerson).Methods("POST")
	r.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r))
}
