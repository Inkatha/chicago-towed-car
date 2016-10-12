package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// CarDetails contains the details of a towed car
type CarDetails []struct {
	color          string `json:"color"`
	make           string `json:"make"`
	state          string `json:"state"`
	towDate        string `json:"tow_date"`
	towPhoneNumber string `json:"tow_facility_phone"`
	towedToAddress string `json:"towed_to_address"`
}

func (details *CarDetails) Color() string {
	return details.color
}

func (details *CarDetails) Make() string {
	return details.make
}

func (details *CarDetails) State() string {
	return details.state
}

func (details *CarDetails) TowDate() string {
	return details.towDate
}

func (details *CarDetails) TowPhoneNumber() string {
	return details.towPhoneNumber
}

func (details *CarDetails) TowedToAddress() string {
	return details.towedToAddress
}

func (details *CarDetails) SetColor(value string) {
	details.color = value
}

func (details *CarDetails) SetMake(value string) {
	details.make = value
}

func (details *CarDetails) SetState(value string) {
	details.state = value
}

func (details *CarDetails) SetTowDate(value string) {
	details.TowDate = value
}

func (details *CarDetails) SetTowPhoneNumber(value string) {
	details.towPhoneNumber = value
}

func (details *CarDetails) SetTowedToAddress(value string) {
	details.towedToAddress = value
}

// GetTowedCarData calls the Chicago API and returns a list
// of cars that have been towed in the last 7 days.
func GetTowedCarData() ([]CarDetails, error) {
	url := fmt.Sprintf("https://data.cityofchicago.org/resource/rp42-fxjt.json")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
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
		return
	}

	// Callers should close resp.Body
	// When done reading from it
	// Defer the closing of the Body
	defer resp.Body.Close()

	// Fill the record with the data from the json
	var record CarDetails

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
		return record, errors.New("An error occured while decoding the response.")
	}

	for _, element := range record {
		fmt.Println("Color = ", element.Color)
		fmt.Println("Make   = ", element.Make)
		fmt.Println("State  = ", element.State)
		fmt.Println("TowDate   = ", element.TowDate)
		fmt.Println("TowPhoneNumber  = ", element.TowPhoneNumber)
		fmt.Println("TowedToAddress  = ", element.TowedToAddress)
	}

	return record, nil
}
