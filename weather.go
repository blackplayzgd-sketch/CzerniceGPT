package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeocodingResponse struct {
	Results []Location `json:"results"`
}

type WeatherResponse struct {
	Place string

	Current struct {
		Temperature         float64 `json:"temperature_2m"`
		ApparentTemperature float64 `json:"apparent_temperature"`
		WindSpeed           float64 `json:"wind_speed_10m"`
		WindDirection       float64 `json:"wind_direction_10m"`
		CloudCover          float64 `json:"cloud_cover"`
		RelativeHumidity    float64 `json:"relative_humidity_2m"`
	} `json:"current"`

	Daily struct {
		Time                   []string  `json:"time"`
		PrecipitationSum       []float64 `json:"precipitation_sum"`
		TemperatureMin         []float64 `json:"temperature_2m_min"`
		TemperatureMax         []float64 `json:"temperature_2m_max"`
		ApparentTemperatureMin []float64 `json:"apparent_temperature_min"`
		ApparentTemperatureMax []float64 `json:"apparent_temperature_max"`
		WeatherCode            []int     `json:"weather_code"`
	}
}

const (
	Temperature         string = "temperature_2m"
	ApparentTemperature string = "apparent_temperature"
	WindSpeed           string = "wind_speed_10m"
	WindDirection       string = "wind_direction_10m"
	CloudCover          string = "cloud_cover"
	RelativeHumidity    string = "relative_humidity_2m"

	Precipitation          string = "precipitation_sum"
	TemperatureMin         string = "temperature_2m_min"
	TemperatureMax         string = "temperature_2m_max"
	ApparentTemperatureMin string = "apparent_temperature_min"
	ApparentTemperatureMax string = "apparent_temperature_max"
	WeatherCode            string = "weather_code"
)

var (
	weatherStatsCurrentString string
	weatherStatsDailyString   string
)

func handleWeather(place string) {
	params := url.Values{}
	location := getLocationData(place)

	params.Add("latitude", strconv.FormatFloat(location.Latitude, 'f', -1, 64))
	params.Add("longitude", strconv.FormatFloat(location.Longitude, 'f', -1, 64))
	params.Add("models", "dwd_icon_seamless")
	params.Add("current", weatherStatsCurrentString)
	params.Add("daily", weatherStatsDailyString)
	params.Add("timezone", "Europe/Berlin")
	params.Add("forecast_days", "7")

	requestUrl := "https://api.open-meteo.com/v1/forecast?" + params.Encode()
	//fmt.Println(requestUrl)

	rq, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		panic(err)
	}

	rq.Header.Set("User-Agent", "OknoxGPT/1.0")

	client := &http.Client{}

	resp, err := client.Do(rq)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var weatherResponse WeatherResponse
	weatherResponse.Place = location.Name

	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		panic(err)
	}

	getWeatherOutput(weatherResponse)
}

func loadWeather() {
	weatherStatsCurrentString += Temperature + ","
	weatherStatsCurrentString += ApparentTemperature + ","
	weatherStatsCurrentString += WindSpeed + ","
	weatherStatsCurrentString += WindDirection + ","
	weatherStatsCurrentString += CloudCover + ","
	weatherStatsCurrentString += RelativeHumidity

	weatherStatsDailyString += Precipitation + ","
	weatherStatsDailyString += TemperatureMin + ","
	weatherStatsDailyString += TemperatureMax + ","
	weatherStatsDailyString += ApparentTemperatureMin + ","
	weatherStatsDailyString += ApparentTemperatureMax + ","
	weatherStatsDailyString += WeatherCode

}

func getLocationData(place string) Location {
	params := url.Values{}
	params.Add("name", place)
	params.Add("count", "1")
	params.Add("language", "en")
	params.Add("format", "json")
	resp, err := http.Get("https://geocoding-api.open-meteo.com/v1/search?" + params.Encode())
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var geoResponse GeocodingResponse
	//fmt.Println("https://geocoding-api.open-meteo.com/v1/search?name=" + place + "&count=1&language=en&format=json")
	err = json.Unmarshal(body, &geoResponse)
	if err != nil {
		panic(err)
	}

	location := geoResponse.Results[0]

	return location
}

func getWeatherOutput(weatherResponse WeatherResponse) {
	fmt.Println(toASCIIArtText("Weather in " + weatherResponse.Place + ":"))
	fPrintln("Current:")
	fPrintlnf("Temperature: " + fToStr(weatherResponse.Current.Temperature) + "°C")
	fPrintlnf("Apparent Temperature: " + fToStr(weatherResponse.Current.ApparentTemperature) + "°C")
	fPrintlnf("Wind: " + fToStr(weatherResponse.Current.WindSpeed) + "km/h coming from the " + getWindDirectionFromAngle(weatherResponse.Current.WindDirection))
	fPrintlnf("Cloud Cover: " + fToStr(weatherResponse.Current.CloudCover) + "%")
	fPrintlnf("Humidity: " + fToStr(weatherResponse.Current.RelativeHumidity) + "%")

	getForecastOutput(weatherResponse)
}

func getForecastOutput(weatherResponse WeatherResponse) {
	daily := weatherResponse.Daily

	fPrintln("\nForecast for the following week: ")
	for i := range len(daily.Time) {
		fPrintlnf(getWeatherDateString(daily.Time[i]))
		fPrintlnf("Total Precipitation: " + fToStr(daily.PrecipitationSum[i]) + "mm")
		fPrintlnf("Minimum / Maximum Temperature: " + fToStr(daily.TemperatureMin[i]) + "°C / " + fToStr(daily.TemperatureMax[i]) + "°C")
		fPrintlnf("Minimum / Maximum Apparent Temperature: " + fToStr(daily.ApparentTemperatureMin[i]) + "°C / " + fToStr(daily.ApparentTemperatureMax[i]) + "°C")
		fPrintlnf(parseWeatherCode(daily.WeatherCode[i]))
	}
}

func getWeatherDateString(date string) string {
	timeForm, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	niceForm := timeForm.Format("2 Jan")

	dayOfTheWeek := timeForm.Weekday().String()

	return "\n" + dayOfTheWeek + ", " + niceForm + ": "
}

func parseWeatherCode(code int) string {
	switch code {
	case 0:
		return "Clear Sky"
	case 1:
		return "Mainly Clear Sky"
	case 2:
		return "Partly Cloudy"
	case 3:
		return "Overcast"
	case 45:
		return "Foggy"
	case 48:
		return "Depositing Rime Fog"
	case 51:
		return "Light Drizzle"
	case 53:
		return "Moderate Drizzle"
	case 55:
		return "Dense Drizzle"
	case 56:
		return "Light Freezing Drizzle"
	case 57:
		return "Heavy Freezing Drizzle"
	case 61:
		return "Slight Rain"
	case 63:
		return "Moderate Rain"
	case 65:
		return "Heavy Rain"
	case 66:
		return "Light Freezing Rain"
	case 67:
		return "Heavy Freezing Rain"
	case 71:
		return "Slight Snowfall"
	case 73:
		return "Moderate Snowfall"
	case 75:
		return "Heavy Snowfall"
	case 77:
		return "Snow Grains"
	case 80:
		return "Slight Rain Showers"
	case 81:
		return "Moderate Rain Showers"
	case 82:
		return "Violent Rain Showers"
	case 85:
		return "Slight Snow Showers"
	case 86:
		return "Heavy Snow Showers"
	case 95:
		return "Thunderstorm"
	case 96:
		return "Thunderstorm w/ Slight Hail"
	case 99:
		return "Thunderstorm w/ Heavy Hail"
	default:
		return "Invalid weather code"
	}
}
