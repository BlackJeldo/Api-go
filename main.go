package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Types
type ticket struct {
	ID      			int    `json:"ID"`
	User    			string `json:"User"`
	Fecha_creacion 		string `json:"Fecha_creacion"`
	Fecha_actualizacion string `json:"Fecha_actualizacion"`
	Status    			bool   `json:"Status"`
}

type allTickets []ticket

// Persistence
var tickets = allTickets{
	{
		ID:      1,
		User:    "Task One",
		Fecha_creacion: "20/10/2021",
		Fecha_actualizacion: "25/10/2021",
		Status: false,
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my GO API!")
}

func createTicket(w http.ResponseWriter, r *http.Request) {
	var newTicket ticket
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a valid ticket")
	}

	json.Unmarshal(reqBody, &newTicket)

	newTicket.ID = tickets[len(tickets) - 1].ID + 1
	tickets = append(tickets, newTicket)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTicket)
}

func updateTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketId, err := strconv.Atoi(vars["id"])
	var updateTicket ticket

	if err !=nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err !=nil {
		fmt.Fprintf(w, "Please enter valid data")
		return
	}
	
	json.Unmarshal(reqBody, &updateTicket)

	for i, ticket := range tickets {
		if ticketId == ticket.ID {
			updateTicket.ID = ticketId
			tickets[i] =updateTicket

			fmt.Fprintf(w, "The ticket with ID %v has been updated succesfully", ticketId)
		}
	} 
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketId, err := strconv.Atoi(vars["id"])

	if err !=nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for _, ticket := range tickets {
		if ticketId == ticket.ID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ticket)
		}
	} 
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketId, err := strconv.Atoi(vars["id"])

	if err !=nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for i, ticket := range tickets {
		if ticketId == ticket.ID {
			tickets = append(tickets[:i], tickets[i + 1:]...)
			fmt.Fprintf(w, "The ticket with ID %v has benn remove succesfully", ticketId)
		}
	} 
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/ticket", getTasks).Methods("GET")
	router.HandleFunc("/ticket", createTicket).Methods("POST")
	router.HandleFunc("/ticket/{id}", getTask).Methods("GET")
	router.HandleFunc("/ticket/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/ticket/{id}", updateTicket).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
