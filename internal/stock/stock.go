package stock

import (
	"context"
	"fmt"

	v1 "github.com/tomato-net/go-sticker/api/v1"
	"github.com/tomato-net/go-sticker/pkg/alphavantage"
)

// Repository is used to query stock data about a given stock symbol.
type Repository struct {
	client *alphavantage.Client
}

func NewRepository(client *alphavantage.Client) *Repository {
	return &Repository{
		client: client,
	}
}

// DailyStockData queries the last N days of stock data for a given symbol and returns that data back in
// an ordered slice.
func (r *Repository) DailyStockData(ctx context.Context, symbol string, days int) ([]v1.DatedStock, error) {
	data, err := r.client.TimeSeriesDailyAdjusted(ctx, alphavantage.TimeSeriesDailyOptions{
		Symbol: symbol,
	})
	if err != nil {
		return nil, fmt.Errorf("client.TimeSeriesDailyAdjusted: %w", err)
	}

	sortedData := data.TimeSeriesData.Sorted()
	days = min(len(sortedData), days)

	datedStocks := make([]v1.DatedStock, days)
	for i, stockData := range data.TimeSeriesData.Sorted()[:days] {
		datedStocks[i] = convertStockDataToDatedStock(stockData)
	}

	return datedStocks, nil
}

func convertStockDataToDatedStock(dirty alphavantage.StockData) v1.DatedStock {
	return v1.DatedStock{
		Date: dirty.Time,
		Stock: v1.Stock{
			Open:   dirty.Open,
			High:   dirty.High,
			Low:    dirty.Low,
			Close:  dirty.Close,
			Volume: dirty.Volume,
		},
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
