package alphavantage

type Function string
type OutputSize string

const (
	ParamKeyAPIKey = "apikey"

	ParamKeyFunction                         = "function"
	FunctionTimeSeriesDailyAdjusted Function = "TIME_SERIES_DAILY_ADJUSTED"

	ParamKeyOutputSize            = "outputsize"
	OutputSizeCompact  OutputSize = "compact"
	OutputSizeFull     OutputSize = "full"

	ParamKeySymbol = "symbol"

	ParamKeyDataType = "datatype"
	DataTypeJSON     = "json"
)
