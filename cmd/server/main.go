package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomato-net/go-sticker/internal/stock"
	stockhttp "github.com/tomato-net/go-sticker/internal/stock/http"
	"github.com/tomato-net/go-sticker/pkg/alphavantage"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Print("go stock ticker")

	// TODO: Do not use gin.Default so we can use structured logging middleware
	router := gin.Default()

	client, err := alphavantage.NewClient(alphavantage.ClientOptions{
		APIKey: os.Getenv("API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	repo := stock.NewRepository(client)
	symbol := os.Getenv("SYMBOL")
	days, err := strconv.Atoi(os.Getenv("NDAYS"))
	if err != nil {
		log.Fatalf("error parsing NDAYS env var: %s", err.Error())
	}

	stockHandler := stockhttp.Handler(repo, symbol, days)

	routerV1Group := router.Group("v1")
	routerV1Group.GET("stock", stockHandler)

	// TODO: Register debug pprof handlers when profiling

	log.Print("starting server...")
	if err := router.Run(":" + os.Getenv("HTTP_PORT")); err != nil {
		log.Print(err)
	}

	log.Print("shutting down server...")
}
