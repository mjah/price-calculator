package main

import (
	"net/http"

	pc "./pricecalculator"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	case getSellPriceByProfitRate:
		result, err = priceCalc.GetSellPriceByProfitRate()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			c.AbortWithStatus(http.StatusNotAcceptable)
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

	// price group
	price := api.Group("/price")
	{
		price.GET("/by_profit_rate", func(c *gin.Context) {
			executeRequest(c, getSellPriceByProfitRate)
		})
	}

	// fees group
	fees := api.Group("/fees")
	{
		fees.GET("/total", func(c *gin.Context) {
			executeRequest(c, getFeesTotal)
		})
		fees.GET("/sales_tax", func(c *gin.Context) {
			executeRequest(c, getSalesTaxFeesTotal)
		})
		fees.GET("/payment", func(c *gin.Context) {
			executeRequest(c, getPaymentFeesTotal)
		})
		fees.GET("/channel", func(c *gin.Context) {
			executeRequest(c, getChannelFeesTotal)
		})
		fees.GET("/other", func(c *gin.Context) {
			executeRequest(c, getOtherFeesTotal)
		})
	}

	// profit group
	profit := api.Group("/profit")
	{
		profit.GET("/total", func(c *gin.Context) {
			executeRequest(c, getProfitTotal)
		})
		profit.GET("/is_valid", func(c *gin.Context) {
			executeRequest(c, isValidProfitRate)
		})
	}

	r.Run(":9096")
}
