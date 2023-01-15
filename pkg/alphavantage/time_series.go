package alphavantage

import (
	"fmt"
	"sort"
	"time"
)

// TODO: Comment - Override Time to enable unmarshalling
type Time time.Time

func (t *Time) UnmarshalText(data []byte) error {
	format := "2006-01-02"
	d, err := parseDate(string(data), format)
	if err != nil {
		return fmt.Errorf("parseDate: %w", err)
	}

	*t = Time(d)

	return nil
}

func (t Time) String() string {
	return time.Time(t).String()
}

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

type TimeSeriesStockData map[Time]StockData

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

func parseDate(v string, format string) (time.Time, error) {
	t, err := time.Parse(format, v)
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse date %s: %w", v, err)
	}

	return t, nil
}
