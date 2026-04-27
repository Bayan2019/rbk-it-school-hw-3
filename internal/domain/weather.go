package domain

type ProviderWeatherResponse struct {
	Temperature float64
	WindSpeed   float64
	WeatherCode int
	Time        string
	Description string `json:"description"`
}

type WeatherResult struct {
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Temperature    float64 `json:"temperature"`
	WindSpeed      float64 `json:"wind_speed"`
	WeatherCode    int     `json:"weather_code"`
	Time           string  `json:"time"`
	Description    string  `json:"description"`
	Recommendation string  `json:"recommendation"`
}
