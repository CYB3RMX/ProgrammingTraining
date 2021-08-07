package main

import (
	"encoding/json" // Parsing JSON data
	"fmt"           // Simple input/output operations
	"io/ioutil"     // Parsing response data
	"net/http"      // Making HTTP requests
)

// Creating a struct that will parse incoming json data
type TargetIP struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	Countrycode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Latitude    float32 `json:"lat"`
	Longtitude  float32 `json:"lon"`
	Timezone    string  `json:"timezone"`
}

func main() {
	var ipaddr string // Target IP address to scan
	fmt.Printf("Enter target IP: ")
	fmt.Scanln(&ipaddr) // Input from user

	scanaddr := fmt.Sprintf("http://ip-api.com/json/%s", ipaddr) // Parsing URL

	// Making request and handle error
	query, err := http.Get(scanaddr)
	if err != nil {
		fmt.Println(err)
	}

	// Parse response and handle error
	querydata, err := ioutil.ReadAll(query.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Define struct and parse
	var targetip TargetIP
	errt := json.Unmarshal(querydata, &targetip)
	if errt != nil {
		fmt.Println(errt)
	}

	// Printing all
	fmt.Printf("\nStatus: %s || Country: %s\n", targetip.Status, targetip.Country)
	fmt.Printf("CountryCode: %s || Region: %s\n", targetip.Countrycode, targetip.Region)
	fmt.Printf("RegionName: %s || City: %s\n", targetip.RegionName, targetip.City)
	fmt.Printf("ISP: %s || Organization: %s || As: %s\n", targetip.Isp, targetip.Org, targetip.As)
	fmt.Printf("Latitude: %f || Longtitude: %f\n", targetip.Latitude, targetip.Longtitude)
	fmt.Printf("Timezone: %s\n", targetip.Timezone)
}
