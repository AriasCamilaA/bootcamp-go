package main

import (
	"fmt"
	"log"

	"github.com/AriasCamilaMeli/go-bases/internal/tickets"
)

func main() {
	country := "Colombia"
	time := "tarde"

	total, err := tickets.GetTotalTickets(country)
	if err != nil {
		log.Fatalf("Error getting total tickets: %v", err)
	}
	fmt.Printf("Total tickets to %s: %d\n", country, total)

	morningCount, err := tickets.GetMornings(time)
	if err != nil {
		log.Fatalf("Error getting morning tickets: %v", err)
	}
	fmt.Printf("Total tickets in the %s: %d\n", time, morningCount)

	average, err := tickets.AverageDestination(country, total)
	if err != nil {
		log.Fatalf("Error getting average destination: %v", err)
	}
	fmt.Printf("Average percentage of tickets to %s: %d%%\n", country, average)
}
