package alphavantage

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime_UnmarshalText(t *testing.T) {
	t.Parallel()

	t.Run("successfully parses time", func(t *testing.T) {
		subject := Time{}
		gotErr := subject.UnmarshalText([]byte("2023-01-15"))
		assert.Nil(t, gotErr)
		assert.Equal(t, "2023-01-15 00:00:00 +0000 UTC", subject.String())
	})

	t.Run("fails to parse if missing format", func(t *testing.T) {
		t.Skip("TODO")
	})
}

func TestTimeSeriesStockData_Sorted(t *testing.T) {
	t.Parallel()

	now := time.Now()

	tests := []struct {
		name           string
		giveTimeSeries TimeSeriesStockData
		wantSortedData []StockData
	}{
		{
			name: "sorts data by time descending",
			giveTimeSeries: TimeSeriesStockData{
				Time(now): StockData{
					Close: 100.0,
				},
				Time(now.Add(-1 * time.Second)): StockData{
					Close: 101.0,
				},
				Time(now.Add(1 * time.Second)): StockData{
					Close: 99.0,
				},
			},
			wantSortedData: []StockData{
				{
					Close: 99.0,
					Time:  now.Add(1 * time.Second),
				},
				{
					Close: 100.0,
					Time:  now,
				},
				{
					Close: 101.0,
					Time:  now.Add(-1 * time.Second),
				},
			},
		},
		{
			name:           "returns no elements if map empty",
			giveTimeSeries: TimeSeriesStockData{},
			wantSortedData: []StockData{},
		},
		// TODO: Add tests for matching times
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotSortedData := test.giveTimeSeries.Sorted()
			assert.ElementsMatch(t, test.wantSortedData, gotSortedData)
		})
	}
}
