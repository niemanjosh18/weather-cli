package weatherUtils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetGeocode(city string, apiKey string) (Geocode, error) {
	var results Geocode
	url := "https://maps.googleapis.com/maps/api/geocode/json?address=" + city + "&key=" + apiKey

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.StatusCode)
		return results, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body", err)
		return results, err
	}

	// fmt.Printf("Response Body: \n%s\n", body)

	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Printf("Error parsing json \n%v\n", err)
		return results, err
	}
	return results, nil
}

func GetWeather(lat float64, lng float64, hours int, apiKey string) (WeatherResponse, error) {
	weatherTemplate := `https://weather.googleapis.com/v1/forecast/hours:lookup?key=%s&location.latitude=%f&location.longitude=%f&hours=%d&unitsSystem=IMPERIAL`
	weatherUrl := fmt.Sprintf(weatherTemplate, apiKey, lat, lng, hours)
	resp, err := http.Get(weatherUrl)
	var weatherResults WeatherResponse

	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
	}

	defer resp.Body.Close()

	// fmt.Println("resp code:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.StatusCode)
		return weatherResults, errors.New("Error. Non-OK StatusCode")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body", err)
		return weatherResults, err
	}

	// fmt.Printf("Response Body: \n%s\n", body)

	err = json.Unmarshal(body, &weatherResults)
	if err != nil {
		fmt.Printf("Error parsing json \n%v\n", err)
		return weatherResults, err
	}

	return weatherResults, nil
}
