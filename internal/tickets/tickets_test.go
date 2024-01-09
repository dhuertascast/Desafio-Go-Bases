package tickets_test

import (
	"testing"
	//"github.com/dhuertascast"
)

func TestAirline_FindTotalDestinations_Positive(t *testing.T) {
	airline := tickets.Airline{
		Tickets24Hours: []tickets.Ticket{
			{Destination: "Brazil"},
			{Destination: "USA"},
			{Destination: "France"},
			{Destination: "Brazil"},
		},
	}

	expected := 3
	result, err := airline.FindTotalDestinations()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}

func TestAirline_FindTotalDestinations_Empty(t *testing.T) {
	airline := tickets.Airline{
		Tickets24Hours: []tickets.Ticket{},
	}

	expectedErr := tickets.ErrNoTickets
	result, err := airline.FindTotalDestinations()

	if err == nil || err != expectedErr {
		t.Errorf("Expected error: %v, Got: %v", expectedErr, err)
	}

	if result != 0 {
		t.Errorf("Expected: 0, Got: %v", result)
	}
}

func TestAirline_GetTotalTickets_Positive(t *testing.T) {
	airline := tickets.Airline{
		Tickets24Hours: []tickets.Ticket{
			{Destination: "Brazil"},
			{Destination: "USA"},
			{Destination: "Brazil"},
			{Destination: "USA"},
			{Destination: "France"},
		},
	}

	expected := 2
	result, err := airline.GetTotalTickets("USA")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}

func TestAirline_GetTotalTickets_Zero(t *testing.T) {
	airline := tickets.Airline{
		Tickets24Hours: []tickets.Ticket{
			{Destination: "Brazil"},
			{Destination: "France"},
		},
	}

	expected := 0
	result, err := airline.GetTotalTickets("USA")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}
