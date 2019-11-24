package main

import (
    "io/ioutil"
    "fmt"
    "net/http"
)

func apikey() (string) {
	key, err := ioutil.ReadFile(".api_key")
	if err != nil {
		panic(err)
	}
	return string(key)
}

func query() {
	key := apikey()
	
	url := fmt.Sprintf("https://api.shodan.io/shodan/host/121.42.98.22?key=%s", key)
	fmt.Print(url)
	resp, err := http.Get(url)
	if err != nil {
	     panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(body))
}

func main() {
	apikey()
	query()
}
