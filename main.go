package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Event struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Descritpion string `json:"descritpion"`
}

type allEvents []Event

var events = allEvents{
	{
		Id:          1,
		Title:       "API",
		Descritpion: "Trocando Conhecimento com a turma",
	},
}

func main() {
	log.Println("Starting API")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/HandleCheck", HandleCheck).Methods("GET")
	router.HandleFunc("/events", GetAllEvents).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func HandleCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando HandleCheck!")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Aplicação em execução!\n")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aplicação em execução!\n")
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando o endpoint get all events")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
