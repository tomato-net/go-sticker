package v1

import "time"

type StockResponse struct {
	Average Stock        `json:"average"`
	History []DatedStock `json:"history"`
}

type Stock struct {
	Open   float64 `json:"open,omitempty"`
	High   float64 `json:"high,omitempty"`
	Low    float64 `json:"low,omitempty"`
	Close  float64 `json:"close,omitempty"`
	Volume int64   `json:"volume,omitempty"`
}

type DatedStock struct {
	Date  time.Time `json:"date"`
	Stock Stock     `json:"stock"`
}
