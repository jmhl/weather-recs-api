package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"./forecast"
)

const (
	LATITUDE = "latitude"
	LONGITUDE = "longitude"
)

func response(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte("Hello world"))
}

func WeatherHandler(rw http.ResponseWriter, request *http.Request) {
	u := request.URL
	query, _ := url.ParseQuery(u.RawQuery)

	lat, _ := strconv.ParseFloat(query[LATITUDE][0], 64)
	long, _ := strconv.ParseFloat(query[LONGITUDE][0], 64)

	f , _ := forecast.GetForecast(lat, long)
	formattedJSON, err := json.MarshalIndent(f, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	rw.Write(formattedJSON)
}

func main() {
	http.HandleFunc("/", response)
	http.HandleFunc("/weather", WeatherHandler)
	http.ListenAndServe(":3000", nil)
}
