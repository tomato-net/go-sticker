package v1

import "time"

// StockResponse is a model for the stock API response.
type StockResponse struct {
	// Average is a calculated average of stock data over collection of History data.
	Average Stock `json:"average"`
	// History is a collection of historical stock prices.
	History []DatedStock `json:"history"`
}

// Stock is a model for stock data, used in the stock API.
type Stock struct {
	Open   float64 `json:"open,omitempty"`
	High   float64 `json:"high,omitempty"`
	Low    float64 `json:"low,omitempty"`
	Close  float64 `json:"close,omitempty"`
	Volume int64   `json:"volume,omitempty"`
}

// DatedStock is Stock information with added date information.
type DatedStock struct {
	Date  time.Time `json:"date"`
	Stock Stock     `json:"stock"`
}
