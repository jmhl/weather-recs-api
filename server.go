package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"./forecast"
)

const (
	LATITUDE = "latitude"
	LONGITUDE = "longitude"
	PORT = ":3000"
)

func logAndPrintErr(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

func response(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte("Hello world"))
}

func WeatherHandler(rw http.ResponseWriter, request *http.Request) {
	u := request.URL
	query, err := url.ParseQuery(u.RawQuery)
	logAndPrintErr(err)

	lat, err := strconv.ParseFloat(query.Get(LATITUDE), 64)
	logAndPrintErr(err)
	long, err := strconv.ParseFloat(query.Get(LONGITUDE), 64)
	logAndPrintErr(err)

	f , _ := forecast.GetForecast(lat, long)
	formattedJSON, err := json.MarshalIndent(f, "", "\t")
	logAndPrintErr(err)

	rw.Write(formattedJSON)
}

func main() {
	http.HandleFunc("/", response)
	http.HandleFunc("/weather/", WeatherHandler)

	log.Println("Listening on port", PORT)

	http.ListenAndServe(PORT, nil)
}
