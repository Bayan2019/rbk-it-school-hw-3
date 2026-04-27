package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/Bayan2019/rbk-it-school-hw-3/internal/domain"
)

type weatherService interface {
	CreateHistory(ctx context.Context, userID int64, cityWeather domain.CityWeatherInput) (domain.WeatherHistoryResponse, error)
	WeatherHistoryOfUser(ctx context.Context, userID int64, filter domain.WeatherHistoryFilter) ([]domain.WeatherHistoryResponse, error)
}

type weatherProvider interface {
	GetCurrentWeather(ctx context.Context, lat, lon float64) (domain.ProviderWeatherResponse, error)
}

type weatherHistoryResponse struct {
	Data []domain.WeatherHistoryResponse `json:"data"`
}

type WeatherHandler struct {
	CityService     cityService
	WeatherService  weatherService
	WeatherProvider weatherProvider
}

////// methods
////// methods
////// methods

func (h *WeatherHandler) GetWeatherOfUserCities(w http.ResponseWriter, r *http.Request) {
	userID, err := parseIDParam(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error(), Message: "invalid path variable"})
		return
	}

	cities, err := h.CityService.ListOfUser(r.Context(), userID, domain.ListCitiesFilter{})
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: err.Error(), Message: "couldn't get cities of user"})
		return
	}

	results := []domain.CityWeatherResponse{}

	for _, city := range cities {
		res, err := h.WeatherProvider.GetCurrentWeather(r.Context(), city.Lat, city.Lon)
		if err != nil {
			h.handleError(w, err)
			return
		}
		results = append(results, domain.CityWeatherResponse{
			City:        city.City,
			Temperature: res.Temperature,
			Description: res.Description,
		})
	}

	for _, result := range results {
		_, err = h.WeatherService.CreateHistory(r.Context(), userID, domain.CityWeatherInput{
			City:        result.City,
			Temperature: result.Temperature,
			Description: result.Description,
		})
		if err != nil {
			h.handleError(w, err)
			return
		}
	}

	writeJSON(w, http.StatusOK, results)
}

func (h *WeatherHandler) GetWeatherHistoryOfUser(w http.ResponseWriter, r *http.Request) {
	filter := domain.WeatherHistoryFilter{
		Limit:  parseIntQuery(r, "limit", 20),
		Offset: parseIntQuery(r, "offset", 0),
		City:   r.URL.Query().Get("city"),
	}
	userID, err := parseIDParam(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error(), Message: "invalid path variable"})
		return
	}

	result, err := h.WeatherService.WeatherHistoryOfUser(r.Context(), userID, filter)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: err.Error(), Message: "couldn't get weather history"})
		return
	}

	writeJSON(w, http.StatusOK, weatherHistoryResponse{Data: result})
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
////// accommodating functions

func (h *WeatherHandler) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidUserID), errors.Is(err, domain.ErrInvalidUserInput):
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
	case errors.Is(err, domain.ErrCityNotFound):
		writeJSON(w, http.StatusNotFound, errorResponse{Error: err.Error()})
	default:
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal server error"})
		// writeJSON(w, http.StatusInternalServerError, errorResponse{Error: err.Error()})
	}
}
