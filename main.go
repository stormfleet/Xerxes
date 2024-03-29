package main

import (
    "io/ioutil"
    "fmt"
    "net/http"
    "strings"
    "encoding/json"
)

type ShodanResp struct {
	RegionCode  interface{}   `json:"region_code"`
	IP          string        `json:"ip"`
	AreaCode    interface{}   `json:"area_code"`
	CountryName string        `json:"country_name"`
	Hostnames   []interface{} `json:"hostnames"`
	PostalCode  interface{}   `json:"postal_code"`
	DmaCode     interface{}   `json:"dma_code"`
	CountryCode string        `json:"country_code"`
	Data        []struct {
		Product    string        `json:"product,omitempty"`
		Os         interface{}   `json:"os"`
		Timestamp  string        `json:"timestamp"`
		Isp        string        `json:"isp"`
		Asn        string        `json:"asn"`
		Banner     string        `json:"banner"`
		Hostnames  []interface{} `json:"hostnames"`
		Devicetype string        `json:"devicetype,omitempty"`
		Location   struct {
			City         interface{} `json:"city"`
			RegionCode   interface{} `json:"region_code"`
			AreaCode     interface{} `json:"area_code"`
			Longitude    float64     `json:"longitude"`
			CountryCode3 string      `json:"country_code3"`
			CountryName  string      `json:"country_name"`
			PostalCode   interface{} `json:"postal_code"`
			DmaCode      interface{} `json:"dma_code"`
			CountryCode  string      `json:"country_code"`
			Latitude     float64     `json:"latitude"`
		} `json:"location"`
		IP      string        `json:"ip"`
		Domains []interface{} `json:"domains"`
		Org     string        `json:"org"`
		Port    int           `json:"port"`
		Opts    struct {
		} `json:"opts"`
	} `json:"data"`
	City         interface{} `json:"city"`
	Longitude    float64     `json:"longitude"`
	CountryCode3 string      `json:"country_code3"`
	Latitude     float64     `json:"latitude"`
	Os           interface{} `json:"os"`
	Ports        []int       `json:"ports"`
}

func apikey() (string) {
	key, err := ioutil.ReadFile(".api_key")
	if err != nil {
		fmt.Print("No API Key found!")
	}
	return string(key)
}

func query() {
	key := apikey()
	ip := "121.42.98.22"
	url := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", ip, key)
	read_url := strings.TrimRight(url, "\r\n")
	resp, err := http.Get(read_url)
	if err != nil {
	     panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("Response Status: ", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var answ ShodanResp
	json.Unmarshal(body, &answ)
	fmt.Print(answ.Ports)
	//fmt.Print(string(body))
}

func main() {
	apikey()
	query()
}
