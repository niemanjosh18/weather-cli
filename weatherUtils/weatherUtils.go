package weatherUtils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func GetApiKey(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open file", path)
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// only need one line from the file
		key := scanner.Text()
		return key, nil
	}
	return "", errors.New("Failed to find API key.")
}

func PrintWeatherLine(fh ForecastHour) {
	time := fh.DisplayDateTime.Hours % 12
	if time == 0 {
		time = 12
	}
	morningText := "AM"
	if fh.DisplayDateTime.Hours >= 12 {
		morningText = "PM"
	}
	temp := fh.FeelsLikeTemperature.Degrees
	desc := fh.WeatherCondition.Description.Text
	precipitation := fh.Precipitation.Probability.Type
	pChance := fh.Precipitation.Probability.Percent
	windSpeed := fh.Wind.Speed.Value

	lineTemplate := "  %2d %2s   %.1fF %30s  %8s  %3d%%    Wind: %3d MPH\n"
	fmt.Printf(lineTemplate, time, morningText, temp, desc, precipitation, pChance, windSpeed)
}
