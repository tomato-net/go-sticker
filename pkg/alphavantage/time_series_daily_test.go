package alphavantage

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_TimeSeriesDailyAdjusted(t *testing.T) {
	t.Parallel()

	t.Run("parses and returns time series data", func(t *testing.T) {
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		defer server.Close()
		client := &Client{
			apiKey:     "test-123",
			baseURL:    server.URL,
			httpClient: &http.Client{},
		}

		wantQuery := url.Values{}
		wantQuery.Add("symbol", "MSFT")
		wantQuery.Add("function", "TIME_SERIES_DAILY_ADJUSTED")
		wantQuery.Add("outputsize", "compact")
		wantQuery.Add("apikey", "test-123")
		wantQuery.Add("datatype", "json")
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, wantQuery, r.URL.Query())

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"Time Series (Daily)":{"2023-01-13":{"4. close": "238.51"}}}`)
		})

		gotResponse, gotErr := client.TimeSeriesDailyAdjusted(context.TODO(), TimeSeriesDailyOptions{
			Symbol: "MSFT",
		})

		assert.Nil(t, gotErr)
		assert.Equal(t, 238.51, gotResponse.TimeSeriesData.Sorted()[0].Close)
	})

	t.Run("sets symbol", func(t *testing.T) {
		t.Skip("TODO")
	})

	t.Run("sets outputsize", func(t *testing.T) {
		t.Skip("TODO")
	})

	t.Run("errors if request fails", func(t *testing.T) {
		t.Skip("TODO")
	})
}
