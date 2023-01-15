package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/tomato-net/go-sticker/api/v1"
	"github.com/tomato-net/go-sticker/internal/stock"
	"github.com/tomato-net/go-sticker/pkg/slice"
)

// Handler is a HTTP handler for Gin that returns the last N days of stock data for a given
// symbol, alongside the average close price for that stock over the same time period.
func Handler(repo *stock.Repository, symbol string, days int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := repo.DailyStockData(ctx, symbol, days)
		if err != nil {
			// TODO: Logging and metrics
			ctx.JSON(http.StatusInternalServerError, v1.InternalServerError{
				Error: "failed to get stock data",
			})
			return
		}

		total := slice.Reduce(response, func(memo float64, stock v1.DatedStock) float64 {
			return memo + stock.Stock.Close
		}, 0)

		ctx.JSON(http.StatusOK, v1.StockResponse{
			Average: v1.Stock{
				Close: total / float64(days),
			},
			History: response,
		})
	}
}
