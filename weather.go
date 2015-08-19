package main

import (
    "encoding/json"
    "fmt"
    "os"
    "./forecast"
)

func main() {
    forecast, err := forecast.GetForecast()
    if err != nil {
	fmt.Println("error:", err)
    }

    formattedJSON, err := json.MarshalIndent(forecast, "", "\t")

    if err != nil {
	fmt.Println("error:", err)
    }

    os.Stdout.Write(formattedJSON)
}
