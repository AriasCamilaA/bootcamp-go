package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var path_csv = "./tickets.csv"

// Ticket represents a flight ticket with various details
type Ticket struct {
	id          int    // Ticket ID
	name        string // Passenger name
	email       string // Passenger email
	destination string // Destination city
	country     string // Destination country
	flight      string // Flight number
	time        string // Flight time
	price       int64  // Ticket price
}

func GetData() ([][]string, error) {
	file, err := os.Open(path_csv) // Open the CSV file
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err) // Return error if file cannot be opened
	}
	defer file.Close() // Ensure the file is closed after function execution

	reader := csv.NewReader(file)    // Create a new CSV reader
	records, err := reader.ReadAll() // Read all records from the CSV file
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err) // Return error if file cannot be read
	}

	return records, nil
}

// GetTotalTickets calculates how many people travel to a specific destination
func GetTotalTickets(destination string) (int, error) {

	records, err := GetData()
	if err != nil {
		return 0, fmt.Errorf("could not get data: %w", err) // Return error if file cannot be read
	}

	count := 0
	for _, record := range records { // Iterate through each record
		if record[3] == destination { // Check if the destination matches
			count++ // Increment count if destination matches
		}
	}

	return count, nil // Return the total count of tickets for the destination
}

// GetMornings calculates how many people travel in the morning
func GetMornings(time string) (int, error) {

	// madrugada (0 → 6), mañana (7 → 12), tarde (13 → 19) y noche (20 → 00):
	records, err := GetData()
	if err != nil {
		return 0, fmt.Errorf("could not get data: %w", err) // Return error if file cannot be read
	}

	count := 0
	for _, record := range records { // Iterate through each record
		_record := record[4][:len(record[4])-3]
		hour, err := strconv.Atoi(_record)
		if err != nil {
			return 0, fmt.Errorf("could not parse hour: %w", err)
		}
		switch time {
		case "madrugada":
			if hour >= 0 && hour < 6 { // Check if the time is in the madrugada range
				count++ // Increment count if time matches
			}
		case "mañana":
			if hour >= 7 && hour < 12 { // Check if the time is in the mañana range
				count++ // Increment count if time matches
			}
		case "tarde":
			if hour >= 13 && hour < 19 { // Check if the time is in the tarde range
				count++ // Increment count if time matches
			}
		case "noche":
			if hour >= 20 || hour == 0 { // Check if the time is in the noche range
				count++ // Increment count if time matches
			}
		}
	}

	return count, nil
}

// AverageDestination calculates the average price for a specific destination
func AverageDestination(destination string, total int) (int, error) {
	records, err := GetData()
	if err != nil {
		return 0, fmt.Errorf("could not get data: %w", err) // Return error if file cannot be read
	}

	totalTickets := len(records) - 1 // Subtract 1 for the header row
	if totalTickets == 0 {
		return 0, fmt.Errorf("no tickets found")
	}

	destinationCount := 0
	for _, record := range records[1:] { // Skip the header row
		if record[3] == destination { // Check if the destination matches
			destinationCount++ // Increment count if destination matches
		}
	}

	percentage := (destinationCount / totalTickets) * 100 // Calculate the percentage
	return percentage, nil                                // Return the percentage
}
