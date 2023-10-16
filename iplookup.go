package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Org      string `json:"org"`
}

func main() {
	response, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		fmt.Println("Error fetching IP information:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("IP Address:", ipInfo.IP)
	fmt.Println("City:", ipInfo.City)
	fmt.Println("Region:", ipInfo.Region)
	fmt.Println("Country:", ipInfo.Country)
	fmt.Println("Location:", ipInfo.Location)
	fmt.Println("Organization:", ipInfo.Org)
}
