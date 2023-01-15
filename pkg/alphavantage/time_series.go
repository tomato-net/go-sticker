package alphavantage

import (
	"fmt"
	"sort"
	"time"
)

// Time is an alias of the time.Time type, used to enable unmarshalling of the date keys in time series data.
type Time time.Time

func (t *Time) UnmarshalText(data []byte) error {
	format := "2006-01-02"
	time, err := parseTime(string(data), format)
	if err != nil {
		return fmt.Errorf("parseTime: %w", err)
	}

	*t = Time(time)
	return nil
}

func (t Time) String() string {
	return time.Time(t).String()
}

// StockData contains data about stock.
type StockData struct {
	Open             float64   `json:"1. open,string"`
	High             float64   `json:"2. high,string"`
	Low              float64   `json:"3. low,string"`
	Close            float64   `json:"4. close,string"`
	AdjustedClose    float64   `json:"5. adjusted close,string"`
	Volume           int64     `json:"6. volume,string"`
	DividendAmount   float64   `json:"7. dividend amount,string"`
	SplitCoefficient float64   `json:"8. split coefficient,string"`
	Time             time.Time `json:"-"`
}

// TimeSeriesStockData is a collection of StockData with additional time information. This type is returned from the
// AlphaVantage API in TIME_SERIES API calls.
type TimeSeriesStockData map[Time]StockData

// Sorted returns a sorted slice of StockData in time descending order. Necessary as Go maps are not sorted.
func (t TimeSeriesStockData) Sorted() []StockData {
	times := make([]Time, 0, len(t))
	for time := range t {
		times = append(times, time)
	}

	sort.Slice(times, func(i, j int) bool {
		return (time.Time)(times[i]).After((time.Time)(times[j]))
	})

	sorted := make([]StockData, 0, len(t))
	for _, v := range times {
		data := t[v]
		data.Time = (time.Time)(v)
		sorted = append(sorted, data)
	}

	return sorted
}

func parseTime(v string, format string) (time.Time, error) {
	t, err := time.Parse(format, v)
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse time %s: %w", v, err)
	}

	return t, nil
}
