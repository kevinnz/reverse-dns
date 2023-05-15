package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type GeoIP struct {
	IP        string  `json:"ip"`
	City      string  `json:"city"`
	Region    string  `json:"region_name"`
	Country   string  `json:"country_name"`
	Zip       string  `json:"zip"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func GetLocation(ip string, accessKey string) (*GeoIP, error) {
	response, err := http.Get("http://api.ipstack.com/" + ip + "?access_key=" + accessKey)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	geoIP := new(GeoIP)
	err = json.Unmarshal(data, geoIP)
	if err != nil {
		return nil, err
	}

	return geoIP, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-address\n", os.Args[0])
		os.Exit(1)
	}

	accessKey, exists := os.LookupEnv("IPSTACK_ACCESS_KEY")
	if !exists {
		fmt.Println("IPSTACK_ACCESS_KEY environment variable not set")
		os.Exit(2)
	}

	location, err := GetLocation(os.Args[1], accessKey)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(3)
	}

	fmt.Println("IP: ", location.IP)
	fmt.Println("City: ", location.City)
	fmt.Println("Region: ", location.Region)
	fmt.Println("Country: ", location.Country)
	fmt.Println("Zip: ", location.Zip)
	fmt.Println("Latitude: ", location.Latitude)
	fmt.Println("Longitude: ", location.Longitude)
}
