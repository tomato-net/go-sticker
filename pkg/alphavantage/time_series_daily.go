package alphavantage

import (
	"context"
	"fmt"
	"net/url"
)

// TimeSeriesDailyOptions are the options for the TimeSeriesDaily API calls.
type TimeSeriesDailyOptions struct {
	// Symbol is the symbol of the company that data is being requested for.
	Symbol string

	// OutputSize determines whether the API returns limited or full data. See OutputSize for more details.
	OutputSize OutputSize
}

// TimeSeriesDailyAdjustedResponse is the response from the API for a TimeSeriesDailyAdjusted API call.
type TimeSeriesDailyAdjustedResponse struct {
	// TimeSeriesData is a collection of stock data with time information.
	TimeSeriesData TimeSeriesStockData `json:"Time Series (Daily)"`
}

// This API returns raw (as-traded) daily open/high/low/close/volume values, daily adjusted close values, and historical split/dividend events of the global equity specified,
// covering 20+ years of historical data.
func (c *Client) TimeSeriesDailyAdjusted(ctx context.Context, queryOptions TimeSeriesDailyOptions) (TimeSeriesDailyAdjustedResponse, error) {
	params := url.Values{}
	params.Set(ParamKeyFunction, string(FunctionTimeSeriesDailyAdjusted))
	params.Set(ParamKeySymbol, queryOptions.Symbol)

	outputSize := OutputSizeCompact
	if queryOptions.OutputSize != "" {
		outputSize = queryOptions.OutputSize
	}
	params.Set(ParamKeyOutputSize, string(outputSize))

	resp := TimeSeriesDailyAdjustedResponse{}
	if err := c.Do(ctx, params, &resp); err != nil {
		return resp, fmt.Errorf("c.Do: %w", err)
	}

	return resp, nil
}
