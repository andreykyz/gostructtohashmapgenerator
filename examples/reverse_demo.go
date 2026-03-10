package main

import (
	"fmt"
	"log"
)

func main() {
	// Example map that matches the Person struct
	m := map[string]any{
		"name": "Bob",
		"age":  25,
		"address": map[string]any{
			"city":    "London",
			"country": "UK",
		},
	}

	// Convert map to Person using generated function
	p, err := MapToPerson(m)
	if err != nil {
		log.Fatalf("MapToPerson failed: %v", err)
	}

	fmt.Printf("Successfully converted map to Person:\n")
	fmt.Printf("  Name: %s\n", p.Name)
	fmt.Printf("  Age: %d\n", p.Age)
	fmt.Printf("  Address.City: %s\n", p.Address.City)
	fmt.Printf("  Address.Country: %s\n", p.Address.Country)

	// Convert back to map using forward function
	m2 := PersonToMap(p)
	fmt.Printf("\nConverted back to map: %#v\n", m2)
}
