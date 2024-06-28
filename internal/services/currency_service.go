package services

import (
	"encoding/json"
	"net/http"

	"github.com/schtvr/crypto-converter/internal/models"
)

const coinbaseAPI = "https://api.coinbase.com/v2/exchange-rates?currency=USD"

func GetCryptoRates() (map[string]string, error) {
	resp, err := http.Get(coinbaseAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rates models.Payload
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return nil, err
	}

	return rates.Data.Rates, nil
}
