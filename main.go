package main

import (
	"flag"
	"fmt"
	"github.com/niemanjosh18/weather-cli/weatherUtils"
	"net/url"
)

func main() {
	var cityFlag = flag.String("city", "", "Get the weather of this city.")
	var hoursFlag = flag.Int("hours", 12, "Number of hours to retrieve.")

	flag.Parse()
	// fmt.Println("Getting the weather for ", *cityFlag)

	// apiKey := os.Getenv("WEATHER_API_KEY")
	apiKey, err := weatherUtils.GetApiKey("../keys/weather-api-key.txt")
	if err != nil {
		fmt.Println("Failed to aquire API key", err)
		return
	}

	results, err := weatherUtils.GetGeocode(url.PathEscape(*cityFlag), apiKey)
	if err != nil {
		fmt.Println("Error getting geocode", err)
		return
	}
	formatted_address := results.Results[0].Formatted_address
	lat := results.Results[0].Geometry.Location.Lat
	lng := results.Results[0].Geometry.Location.Lng

	fmt.Println("Location:", formatted_address)
	// fmt.Printf("lat %v\nlng %v\n", lat, lng)

	weatherResults, err := weatherUtils.GetWeather(lat, lng, *hoursFlag, apiKey)
	if err != nil {
		fmt.Println("Error getting weather", err)
		return
	}

	// fmt.Println("weather results", weatherResults)
	for _, fh := range weatherResults.ForecastHours {
		weatherUtils.PrintWeatherLine(fh)
	}
}
