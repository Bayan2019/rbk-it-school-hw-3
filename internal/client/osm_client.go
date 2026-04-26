package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Bayan2019/rbk-it-school-hw-3/internal/config"
	"github.com/Bayan2019/rbk-it-school-hw-3/internal/service"
	"golang.org/x/time/rate"
)

type OsmClient struct {
	client    *http.Client
	baseURL   string
	limiter   *rate.Limiter
	userAgent string
}

func NewOsmClient(cfg config.ApiConfig) *OsmClient {
	return &OsmClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL:   "https://nominatim.openstreetmap.org",
		limiter:   cfg.Limiter,
		userAgent: cfg.UserAgent,
	}
}

func (client *OsmClient) GetInfoOfCity(ctx context.Context, city string) (service.Place, error) {
	var place service.Place

	url := fmt.Sprintf("%s/search?city=%s&format=json", client.baseURL, city)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// fmt.Println("Error creating request:", err)
		return place, fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Set("User-Agent", client.userAgent)

	res, err := client.client.Do(req)
	if err != nil {
		return place, err
	}
	defer res.Body.Close()

	var places []service.Place

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&places)
	if err != nil {
		return place, err
	}

	if len(places) < 1 {
		return place, fmt.Errorf("city %s not found", city)
	}

	place = places[0]

	return place, nil
}
