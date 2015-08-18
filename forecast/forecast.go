package forecast

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "strconv"
    "../secrets"
)

const FORECAST string = "https://api.forecast.io/forecast/"

type Forecast struct {
    Latitude  Float64
    Longitude Float64
    Timezone
    Offset
    Currently
    Minutely
    Hourly
    Daily
    Alerts
    Flags
    APICalls
    Code
}

func formatUrl(longitude string, latitude string) string {
    return FORECAST + secrets.APIKEY() + "/" + latitude + "," + longitude
}

func formatJSON(blob []byte) (*ForecastResponse, error) {
    var data ForecastResponse
    err := json.Unmarshal(blob, &data)
    if err != nil {
	return nil, err
    }
    return &data, nil
}

// INTERFACE
func GetForecast() (*ForecastResponse, error) {
    latitude := strconv.FormatFloat(-122.431369, 'f', -1, 64)
    longitude := strconv.FormatFloat(37.764467, 'f', -1, 64)
    url := formatUrl(latitude, longitude)
    res, err := http.Get(url)
    if err != nil {
	return nil, err
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
	return nil, err
    }
    return formatJSON(body)
}
