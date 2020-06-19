package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

type Info struct {
    	Name     		string     	`json:"name"`
    	Country_id    	string     	`json:"country_id"`
}

func httpExampleGetJsonDecodeStruct() {
	fmt.Println("--- running httpExampleGetJsonDecodeStruct ---")

	// get http example
	resp, err := http.Get("https://api.mercadolibre.com/sites/MLA")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var info Info

	// decode json to struct
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
}

func main() {
  httpExampleGetJsonDecodeStruct()
}
