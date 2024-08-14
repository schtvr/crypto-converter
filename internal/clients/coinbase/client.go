package clients

import (
	"encoding/json"
	"net/http"

	"github.com/schtvr/crypto-converter/internal/models"
)

type Client struct {
	url string
}

func (c *Client) GetCryptoRates() (map[string]string, error) {
	resp, err := http.Get(c.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rates models.Payload
	err = json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		return nil, err
	}

	return rates.Data.Rates, nil
}

func NewCoinbaseClient() *Client {
	return &Client{
		url: "https://api.coinbase.com/v2/exchange-rates?currency=USD",
	}
}
