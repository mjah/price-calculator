package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	pc "github.com/mjah/price-calculator/server/pricecalculator"
)

// Request types
const (
	getSellPriceByProfitRate int = 0
	getFeesTotal             int = 1
	getSalesTaxFeesTotal     int = 2
	getPaymentFeesTotal      int = 3
	getChannelFeesTotal      int = 4
	getOtherFeesTotal        int = 5
	getProfitTotal           int = 6
	isValidProfitRate        int = 7
	allResults               int = 8
)

// parseRequestMiddleware binds JSON to PriceCalculator struct
func parseRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		priceCalc := pc.New()
		err := c.BindJSON(&priceCalc)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Set("PC", priceCalc)
	}
}

// executeRequest execute requests given the request type
func executeRequest(c *gin.Context, reqType int) {
	priceCalc := c.MustGet("PC").(*pc.PriceCalculator)

	var result float64
	var err error

	switch reqType {
	case allResults:
		resultGetSellPriceByProfitRate, err := priceCalc.GetSellPriceByProfitRate()
		if err != nil {
			resultGetSellPriceByProfitRate = 0
		}
		resultGetFeesTotal := priceCalc.GetFeesTotal()
		resultGetSalesTaxFeesTotal := priceCalc.GetSalesTaxFeesTotal()
		resultGetPaymentFeesTotal := priceCalc.GetPaymentFeesTotal()
		resultGetChannelFeesTotal := priceCalc.GetChannelFeesTotal()
		resultGetOtherFeesTotal := priceCalc.GetOtherFeesTotal()
		resultGetProfitTotal := priceCalc.GetProfitTotal()
		resultIsValidProfitRate := priceCalc.IsValidProfitRate()
		c.JSON(http.StatusOK, gin.H{
			"getSellPriceByProfitRate": resultGetSellPriceByProfitRate,
			"getFeesTotal":             resultGetFeesTotal,
			"getSalesTaxFeesTotal":     resultGetSalesTaxFeesTotal,
			"getPaymentFeesTotal":      resultGetPaymentFeesTotal,
			"getChannelFeesTotal":      resultGetChannelFeesTotal,
			"getOtherFeesTotal":        resultGetOtherFeesTotal,
			"getProfitTotal":           resultGetProfitTotal,
			"isValidProfitRate":        resultIsValidProfitRate,
		})
		return
	case getSellPriceByProfitRate:
		result, err = priceCalc.GetSellPriceByProfitRate()
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": err.Error(),
			})
			return
		}
	case getFeesTotal:
		result = priceCalc.GetFeesTotal()
	case getSalesTaxFeesTotal:
		result = priceCalc.GetSalesTaxFeesTotal()
	case getPaymentFeesTotal:
		result = priceCalc.GetPaymentFeesTotal()
	case getChannelFeesTotal:
		result = priceCalc.GetChannelFeesTotal()
	case getOtherFeesTotal:
		result = priceCalc.GetOtherFeesTotal()
	case getProfitTotal:
		result = priceCalc.GetProfitTotal()
	case isValidProfitRate:
		// look into alternative way for result variable to handle different data types
		resultBool := priceCalc.IsValidProfitRate()
		c.JSON(http.StatusOK, gin.H{
			"result": resultBool,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// main contains the API endpoints
func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(parseRequestMiddleware())

	api := r.Group("/v1")

	// all results
	api.POST("/all_results", func(c *gin.Context) {
		executeRequest(c, allResults)
	})

	// price group
	price := api.Group("/price")
	{
		price.POST("/by_profit_rate", func(c *gin.Context) {
			executeRequest(c, getSellPriceByProfitRate)
		})
	}

	// fees group
	fees := api.Group("/fees")
	{
		fees.POST("/total", func(c *gin.Context) {
			executeRequest(c, getFeesTotal)
		})
		fees.POST("/sales_tax", func(c *gin.Context) {
			executeRequest(c, getSalesTaxFeesTotal)
		})
		fees.POST("/payment", func(c *gin.Context) {
			executeRequest(c, getPaymentFeesTotal)
		})
		fees.POST("/channel", func(c *gin.Context) {
			executeRequest(c, getChannelFeesTotal)
		})
		fees.POST("/other", func(c *gin.Context) {
			executeRequest(c, getOtherFeesTotal)
		})
	}

	// profit group
	profit := api.Group("/profit")
	{
		profit.POST("/total", func(c *gin.Context) {
			executeRequest(c, getProfitTotal)
		})
		profit.POST("/is_valid", func(c *gin.Context) {
			executeRequest(c, isValidProfitRate)
		})
	}

	r.Run(":9096")
}
