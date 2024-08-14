package handlers_test

import (
	"testing"

	h "github.com/schtvr/crypto-converter/internal/handlers"
)

func TestConvertHoldings(t *testing.T) {
	usd := float64(100)
	crypto1 := "BTC"
	crypto2 := "ETH"
	want := "$70.00 => 0.0010 BTC $30.00 => 0.0120 ETH"

	mockClient := newMockClient()

	got, _ := h.ConvertHoldings(usd, crypto1, crypto2, mockClient)

	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

type MockClient struct{}

func (c *MockClient) GetCryptoRates() (map[string]string, error) {
	mockRates := map[string]string{
		"BTC": "0.000015",
		"ETH": "0.0004",
	}
	return mockRates, nil
}

func newMockClient() *MockClient {
	return &MockClient{}
}
