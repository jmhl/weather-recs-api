package forecast

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
    "../secrets"
)

const FORECAST string = "https://api.forecast.io/forecast/"

type DataPoint struct {
    Time		       float64 `json:"time"`
    Summary		       string  `json:"summary"`
    Icon		       string  `json:"icon"`
    SunriseTime		       float64 `json:"nearestStormDistance"`
    SunsetTime		       float64 `json:"nearestStormDistance"`
    MoonPhase		       float64 `json:"nearestStormDistance"`
    NearestStormDistance       float64 `json:"nearestStormDistance"`
    NearestStormBearing	       float64 `json:"nearestStormDistance"`
    PrecipIntensity	       float64 `json:"precipIntensity"`
    PrecipIntensityMax	       float64 `json:"precipIntensityMax"`
    PrecipIntensityMaxTime     float64 `json:"precipIntensityMaxTime"`
    PrecipIntensityError       float64 `json:"precipIntensityError"`
    PrecipProbability	       float64 `json:"precipProbability"`
    PrecipType		       string  `json:"precipType"`
    PrecipAccumulation	       float64 `json:"precipAccumulation"`
    Temperature		       float64 `json:"temperature"`
    TemperatureMin	       float64 `json:"temperatureMin"`
    TemperatureMinTime	       float64 `json:"temperatureMinTime"`
    TemperatureMax	       float64 `json:"temperatureMax"`
    TemperatureMaxTime	       float64 `json:"temperatureMaxTime"`
    ApparentTemperature	       float64 `json:"apparentTemperature"`
    ApparentTemperatureMin     float64 `json:"apparentTemperatureMin"`
    ApparentTemperatureMinTime float64 `json:"apparentTemperatureMinTime"`
    ApparentTemperatureMax     float64 `json:"apparentTemperatureMax"`
    ApparentTemperatureMaxTime float64 `json:"apparentTemperatureMaxTime"`
    DewPoint		       float64 `json:"dewPoint"`
    WindSpeed		       float64 `json:"windSpeed"`
    WindBearing		       float64 `json:"windBearing"`
    CloudCover		       float64 `json:"cloudCover"`
    Humidity		       float64 `json:"humidity"`
    Pressure		       float64 `json:"pressure"`
    Visibility		       float64 `json:"visibility"`
    Ozone		       float64 `json:"ozone"`
}

type DataBlock struct {
    Summary string    `json:"summary"`
    Icon    string    `json:"icon"`
    Data    DataPoint `json:"data"`
}

type Alerts struct {
    Title	string  `json:"title"`
    Expires	float64 `json:"expires"`
    Description string  `json:"description"`
    Uri	        string  `json:"uri"`
}

type Flags struct {
    DarkSkyUnavailable string `json:"darksky-unavailable"`
    DarkSkyStations    []string `json:"darksky-stations"`
    DataPointStations  []string `json:"datapoint-stations"`
    ISDStations	       []string `json:"isd-stations"`
    LAMPStations       []string `json:"lamp-stations"`
    METARStations      []string `json:"metar-stations"`
    MetNoLicsense      string	`json:"metno-license"`
    Sources	       []string `json:"sources"`
    Units	       string	`json:"units"`
}

type Forecast struct {
    Latitude  float64	`json:"latitude"`
    Longitude float64	`json:"longitude"`
    Timezone  string	`json:"timezone"`
    Offset    float64	`json:"offset"`
    Currently DataPoint `json:"currently"`
    Minutely  DataBlock	`json:"minutely"`
    Hourly    DataBlock	`json:"minutely"`
    Daily     DataBlock	`json:"minutely"`
    Alerts    Alerts	`json:"alerts"`
    Flags     Flags	`json:"flags"`
}

func formatUrl(longitude string, latitude string) string {
    return FORECAST + secrets.APIKEY() + "/" + latitude + "," + longitude
}

func formatJSON(blob []byte) (*Forecast, error) {
    var data Forecast
    err := json.Unmarshal(blob, &data)
    if err != nil {
	fmt.Println("error:", err)
	return nil, err
    }
    return &data, nil
}

// INTERFACE
func GetForecast() (*Forecast, error) {
    latitude := strconv.FormatFloat(-122.431369, 'f', -1, 64)
    longitude := strconv.FormatFloat(37.764467, 'f', -1, 64)
    url := formatUrl(latitude, longitude)

    res, err := http.Get(url)
    if err != nil {
	fmt.Println("error:", err)
	return nil, err
    }

    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
	fmt.Println("error:", err)
	return nil, err
    }

    return formatJSON(body)
}
