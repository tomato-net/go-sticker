package alphavantage

type Function string
type OutputSize string

const (
	// ParamKeyAPIKey is the query parameter key for the apikey data
	ParamKeyAPIKey = "apikey"

	// ParamKeyFunction is the query parameter key for the function data
	ParamKeyFunction = "function"
	// FunctionTimeSeriesDailyAdjusted is the function name for the TIME_SERIES_DAILY_ADJUSTED API.
	FunctionTimeSeriesDailyAdjusted Function = "TIME_SERIES_DAILY_ADJUSTED"

	// ParamKeyOutputSize is the query parameter key for the outputsize data. This determines whether
	// the API returns only recent data or all data over the last 20 years.
	ParamKeyOutputSize = "outputsize"
	// OutputSizeCompact will return only data from the last 20 days.
	OutputSizeCompact OutputSize = "compact"
	// OutputSizeFull will return data from the last 20 years.
	OutputSizeFull OutputSize = "full"

	// ParamKeySymbol is the query parameter key for the symbol data. Used to narrow results down to one company.
	ParamKeySymbol = "symbol"

	// ParamKeyDataType is the query parameter key for the datatype data. This client only supports JSON.
	ParamKeyDataType = "datatype"
	DataTypeJSON     = "json"
)
