package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Ticket struct {
	ID          int     `json:"ID"`
	Name        string  `json:"Name"`
	Email       string  `json:"Email"`
	Destination string  `json:"Destination"`
	FlightTime  string  `json:"FlightTime"`
	Price       float64 `json:"Price"`
}

type Airline struct {
	ID             string   `json:"ID"`
	Name           string   `json:"Name"`
	Cnpj           string   `json:"Cnpj"`
	Tickets24Hours []Ticket `json:"Tickets24Hours"`
}

var ErrNoTickets = errors.New("The ticket list is empty")

func (a Airline) FindTotalDestinations() (int, error) {
	var auxSlice []string
	if len(a.Tickets24Hours) > 0 {
		for _, ticket := range a.Tickets24Hours {
			if !Contains(auxSlice, ticket.Destination) {
				auxSlice = append(auxSlice, ticket.Destination)
			}
		}
	} else {
		return 0, ErrNoTickets
	}
	fmt.Errorf("%v", auxSlice)
	return len(auxSlice), nil
}

var ErrTimeFormat = errors.New(`Please insert a valid time format.
	Accepted values for:
		Dawn: "dawn", 0, 00:00 to 06:00
		Morning: "morning", 1, 07:00 to 12:00
		Afternoon: "afternoon", 2, 13:00 to 19:00
		Night: "night", 3, 20:00 to 23:00
	Example: 14:00 or afternoon or 2`)

// GetTotalTickets searches for the total tickets to a specific country in a day
func (a Airline) GetTotalTickets(destination string) (int, error) {
	count := 0
	for _, ticket := range a.Tickets24Hours {
		if strings.ToUpper(ticket.Destination) == strings.ToUpper(destination) {
			count++
		}
	}
	return count, nil
}

// GetMornings searches for the total tickets at a specific time in a day
func (a Airline) GetMornings(time string) (int, error) {
	var shift string

	// Identify format and validate the received parameter: "time"
	if time != Dawn.String() && time != Morning.String() && time != Afternoon.String() && time != Night.String() && time != Dawn.Num() && time != Morning.Num() && time != Afternoon.Num() && time != Night.Num() {
		if strings.Contains(time, ":") {
			format := regexp.MustCompile(`\d\d:\d\d`)
			match := format.MatchString(time)
			if !match {
				return 0, ErrTimeFormat
			}
			parts := strings.SplitN(time, ":", 2)
			var hours, minutes int
			var err error
			hours, err = strconv.Atoi(parts[0])
			if err != nil {
				return 0, err
			}
			minutes, err = strconv.Atoi(parts[1])
			if err != nil {
				return 0, err
			}

			if hours < 0 || hours > 23 || minutes > 59 {
				return 0, ErrTimeFormat
			}
			if hours >= 0 && hours <= 6 {
				shift = "dawn"
			} else if hours >= 7 && hours <= 12 {
				shift = "morning"
			} else if hours >= 13 && hours <= 19 {
				shift = "afternoon"
			} else {
				shift = "night"
			}
		} else {
			num, err := strconv.ParseInt(time, 10, 0)
			if err != nil {
				return 0, err
			}
			if num > 3 {
				return 0, ErrTimeFormat
			}
			shift = time
		}
	} else {
		shift = time
	}

	count := 0
	for _, v := range a.Tickets24Hours {
		partsAux := strings.SplitN(v.FlightTime, ":", 2)
		hoursAux, err := strconv.Atoi(partsAux[0])
		if err != nil {
			return 0, err
		}
		if (hoursAux >= 0 && hoursAux <= 6) && (shift == "dawn" || time == Dawn.Num()) {
			count++
		} else if (hoursAux >= 7 && hoursAux <= 12) && (shift == "morning" || time == Morning.Num()) {
			count++
		} else if (hoursAux >= 13 && hoursAux <= 19) && (shift == "afternoon" || time == Afternoon.Num()) {
			count++
		} else if (hoursAux >= 20 && hoursAux <= 23) && (shift == "night" || time == Night.Num()) {
			count++
		}
	}

	return count, nil
}

// AverageDestination calculates the average of tickets per destination in a day
func (a Airline) AverageDestination(totalTickets int) (float64, error) {
	// total trips/total countries
	totalCountries, err := a.FindTotalDestinations()
	if err != nil {
		return 0, err
	}
	average := float64(totalTickets) / float64(totalCountries)
	return average, nil
}

type shift uint

var shifts = [...]string{"dawn", "morning", "afternoon", "night"}

const (
	Dawn shift = iota
	Morning
	Afternoon
	Night
)

func (s shift) String() string {
	return shifts[s]
}

func (s shift) Num() string {
	switch s {
	case Dawn:
		return "0"
	case Morning:
		return "1"
	case Afternoon:
		return "2"
	case Night:
		return "3"
	}
	return ""
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func OpenCSV(path string) ([]Ticket, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	var ticketsList []Ticket
	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Errorf("error: %w", err)
			break
		}
		id, err := strconv.Atoi(record[0])
		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return nil, err
		}
		ticketsList = append(ticketsList, Ticket{
			ID:          id,
			Name:        record[1],
			Email:       record[2],
			Destination: record[3],
			FlightTime:  record[4],
			Price:       price,
		})
	}
	return ticketsList, nil
}
