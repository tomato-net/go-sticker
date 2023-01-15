package alphavantage

import (
	"context"
	"fmt"
	"net/url"
)

// TODO: Comment
type TimeSeriesDailyOptions struct {
	Symbol     string
	OutputSize OutputSize
}

// TODO: Comment
type TimeSeriesDailyAdjustedResponse struct {
	TimeSeriesData TimeSeriesStockData `json:"Time Series (Daily)"`
}

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
