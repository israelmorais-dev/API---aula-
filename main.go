package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

// @title           Swagger API-Event
// @version         1.0
// @description     Documentação da API de eventos.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API-Event Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://api-events-aula.herokuapp.com/

func main() {
	log.Println("Starting API")
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/HandleCheck", HandleCheck).Methods("GET")
	router.HandleFunc("/events", GetAllEvents).Methods("GET")

	http.ListenAndServe(":"+port, router)
}

func HandleCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando Handle-Check!")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Aplicação em execução!\n")

}

// ShowAllEvents godoc
// @Summary      Show all events
// @Description  List all events
// @Tags         events
// @Accept       json
// @Produce      json
// @Success      200  {object}  Event
// @Router       /events [get]

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aplicação em execução!\n")
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando o endpoint get all events")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
