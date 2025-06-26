package openweather

import (
	"encoding/json"
	"fmt"
	"github.com/ZaharBorisenko/tg-bot/clients/openweather/models"
	"net/http"
)

type Client struct {
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (o *Client) GetCoordinates(city string) (models.Coordinates, error) {
	url := "http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s"
	resp, err := http.Get(fmt.Sprintf(url, city, o.apiKey))

	if err != nil {
		return models.Coordinates{}, fmt.Errorf("error get coordirate %w", err)
	}

	if resp.StatusCode != 200 {
		return models.Coordinates{}, fmt.Errorf("error fail api %d", resp.StatusCode)
	}

	var coordinatesResponse []models.CoordinatesResponse
	err = json.NewDecoder(resp.Body).Decode(&coordinatesResponse)
	if err != nil {
		return models.Coordinates{}, fmt.Errorf("error unmarshal response %w", err)
	}

	if len(coordinatesResponse) == 0 {
		return models.Coordinates{}, fmt.Errorf("len response == 0")
	}

	return models.Coordinates{
		Lat: coordinatesResponse[0].Lat,
		Lon: coordinatesResponse[0].Lon,
	}, nil

}
