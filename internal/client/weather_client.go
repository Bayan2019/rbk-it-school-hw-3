package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Bayan2019/rbk-it-school-hw-3/internal/domain"
)

type WeatherClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewWeatherClient(httpClient *http.Client) *WeatherClient {
	return &WeatherClient{
		httpClient: httpClient,
		baseURL:    "https://api.open-meteo.com/v1/forecast",
	}
}

type openMeteoResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
		Weathercode int     `json:"weathercode"`
		Time        string  `json:"time"`
	} `json:"current_weather"`
}

////// methods
////// methods
////// methods

func (c *WeatherClient) GetCurrentWeather(ctx context.Context, lat, lon float64) (*domain.ProviderWeatherResponse, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("parse base url: %w", err)
	}

	q := u.Query()
	q.Set("latitude", fmt.Sprintf("%.4f", lat))
	q.Set("longitude", fmt.Sprintf("%.4f", lon))
	q.Set("current_weather", "true")
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call external api: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external api returned status: %d", resp.StatusCode)
	}

	var result openMeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode external api response: %w", err)
	}

	return &domain.ProviderWeatherResponse{
		Temperature: result.CurrentWeather.Temperature,
		WindSpeed:   result.CurrentWeather.Windspeed,
		WeatherCode: result.CurrentWeather.Weathercode,
		Time:        result.CurrentWeather.Time,
		Description: mapWeatherCode(result.CurrentWeather.Weathercode),
	}, nil
}

////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions

func mapWeatherCode(code int) string {
	switch code {
	case 0:
		return "Ясно"
	case 1, 2, 3:
		return "Переменная облачность"
	case 45, 48:
		return "Туман"
	case 51, 53, 55:
		return "Морось"
	case 61, 63, 65:
		return "Дождь"
	case 71, 73, 75:
		return "Снег"
	case 95:
		return "Гроза"
	default:
		return "Неизвестно"
	}
}
