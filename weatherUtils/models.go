package weatherUtils

type ForecastHour struct {
	Interval        any `json:"interval"`
	DisplayDateTime struct {
		Year      int    `json:"year"`
		Month     int    `json:"month"`
		Hours     int    `json:"hours"`
		UtcOffset string `json:"utcOffset"`
	}
	WeatherCondition struct {
		Description struct {
			Text string `json:"text"`
		} `json:"description"`
	} `json:"weatherCondition"`
	FeelsLikeTemperature struct {
		Degrees float64 `json:"degrees"`
	} `json:"feelsLikeTemperature"`
	Precipitation struct {
		Probability struct {
			Percent int    `json:"percent"`
			Type    string `json:"type"`
		} `json:"probability"`
	} `json:"precipitation"`
	Wind struct {
		Speed struct {
			Value int `json:"value"`
		} `json:"speed"`
	} `json:"wind"`
}

type WeatherResponse struct {
	ForecastHours []ForecastHour `json:"forecastHours"`
	NextPageToken string         `json:"nextPageToken"`
}

// request from geocoding
type GeometryComponent struct {
	Bounds   any `json:"bounds"`
	Location struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"location"`
	LocationType string `json:"location_type"`
	Viewport     any    `json:"viewport"`
}

type ResultComponents struct {
	Address_components []any             `json:"address_components"`
	Formatted_address  string            `json:"formatted_address"`
	Geometry           GeometryComponent `json:"geometry"`
	PlaceId            any               `json:"place_id"`
	Types              any               `json:"types"`
}

type Geocode struct {
	Results []ResultComponents `json:"results"`
}
