package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/google/uuid"
)

var Airline tickets.Airline

func main() {
	ticketsList, err := tickets.OpenCSV("/Users/dhuertascast/Documents/desafio-go-bases/tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	Airline = tickets.Airline{
		ID:             uuid.New().String(),
		Name:           "LATAM Airlines Brazil",
		Cnpj:           "02.012.862/0001-60",
		Tickets24Hours: ticketsList,
	}

	// Requirement 1
	totalBrazil, err := Airline.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}
	totalUSA, err := Airline.GetTotalTickets("United States")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total tickets for Brazil:", totalBrazil)
	fmt.Println("Total tickets for the United States:", totalUSA)

	// Requirement 2
	ticketsMidnight, err := Airline.GetMornings("0")
	if err != nil {
		fmt.Println(err)
	}
	ticketsMorning, err := Airline.GetMornings("morning")
	if err != nil {
		fmt.Println(err)
	}
	ticketsAfternoon, err := Airline.GetMornings("14:30")
	if err != nil {
		fmt.Println(err)
	}
	ticketsNight, err := Airline.GetMornings("3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Number of people traveling in the midnight:", ticketsMidnight)
	fmt.Println("Number of people traveling in the morning:", ticketsMorning)
	fmt.Println("Number of people traveling in the afternoon:", ticketsAfternoon)
	fmt.Println("Number of people traveling at night:", ticketsNight)
	totalTickets24h := ticketsMidnight + ticketsMorning + ticketsAfternoon + ticketsNight
	fmt.Println("Total tickets in the last 24 hours:", totalTickets24h)

	// Requirement 3
	averageTicketsPerDestination, err := Airline.AverageDestination(len(Airline.Tickets24Hours))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Average tickets per destination in a day: %v\n", averageTicketsPerDestination)
}
