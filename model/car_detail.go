package model

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CarDetails contains the details of a towed car
type CarDetails []struct {
	Color          string `json:"color"`
	Make           string `json:"make"`
	State          string `json:"state"`
	TowDate        string `json:"tow_date"`
	TowPhoneNumber string `json:"tow_facility_phone"`
	TowedToAddress string `json:"towed_to_address"`
}

// GetTowedCarData calls the Chicago API and returns a list
// of cars that have been towed in the last 7 days.
func GetTowedCarData() CarDetails {
	url := fmt.Sprintf("https://data.cityofchicago.org/resource/rp42-fxjt.json")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	/// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A client is an HTTP client
	client := &http.Client{}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	// Callers should close resp.Body
	// When done reading from it
	// Defer the closing of the Body
	defer resp.Body.Close()

	// Fill the record with the data from the json
	var record CarDetails

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record
}
