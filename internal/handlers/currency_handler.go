package handlers

import (
	"fmt"
	"strconv"

	"github.com/schtvr/crypto-converter/internal/services"
)

func ConvertHoldings(usdAmount float64, crypto1, crypto2 string) (string, error) {
	rates, err := services.GetCryptoRates()
	if err != nil {
		return "", err
	}

	rate1, err1 := strconv.ParseFloat(rates[crypto1], 64)
	rate2, err2 := strconv.ParseFloat(rates[crypto2], 64)

	if err1 != nil || err2 != nil {
		return "", fmt.Errorf("invalid cryptocurrency provided")
	}

	amount1 := usdAmount * 0.70 / rate1
	amount2 := usdAmount * 0.30 / rate2

	return fmt.Sprintf("$%.2f => %.4f %s $%.2f => %.4f %s", usdAmount*0.70, amount1, crypto1, usdAmount*0.30, amount2, crypto2), nil
}
