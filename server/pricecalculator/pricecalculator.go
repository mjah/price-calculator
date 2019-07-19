package pricecalculator

import (
	ec "github.com/mjah/price-calculator/server/errors"
)

// PriceCalculator stores the price, costs, and fees
type PriceCalculator struct {
	SellPrice         float64 `json:"sell_price"`
	FreeDeliveryPrice float64 `json:"free_delivery_price"`
	Cost              float64 `json:"cost"`
	Fees              struct {
		SalesTax struct {
			Rate float64 `json:"rate"`
		} `json:"sales_tax"`
		Payment struct {
			Rate  float64 `json:"rate"`
			Fixed float64 `json:"fixed"`
		} `json:"payment"`
		Channel struct {
			Rate        float64 `json:"rate"`
			Fixed       float64 `json:"fixed"`
			IsCapped    bool    `json:"is_capped"`
			CappedValue float64 `json:"capped_value"`
		} `json:"channel"`
		Other struct {
			Rate  float64 `json:"rate"`
			Fixed float64 `json:"fixed"`
		} `json:"other"`
	} `json:"fees"`
	Profit struct {
		Rate float64 `json:"rate"`
	} `json:"profit"`
}

// New creates and returns a new instance of PriceCalculator
func New() *PriceCalculator {
	return &PriceCalculator{}
}

// GetSellPriceByProfitRate calculates and returns the sell price given the profit rate
func (pc *PriceCalculator) GetSellPriceByProfitRate() (float64, error) {
	numerator := pc.Cost
	numerator += pc.FreeDeliveryPrice
	numerator += pc.Fees.Payment.Fixed
	numerator += pc.Fees.Channel.Fixed
	numerator += pc.Fees.Other.Fixed

	denominator := (1 / (1 + pc.Fees.SalesTax.Rate))
	denominator -= pc.Fees.Payment.Rate
	denominator -= pc.Fees.Channel.Rate
	denominator -= pc.Fees.Other.Rate
	denominator -= pc.Profit.Rate

	sellPriceUncappedFees := numerator / denominator

	if !pc.Fees.Channel.IsCapped || pc.Fees.Channel.CappedValue == 0 {
		if denominator <= 0 {
			err := ec.New(ec.NegDenUncappedFees, "Negative denomiator with uncapped fees.")
			return 0, err
		}
		return sellPriceUncappedFees, nil
	} else if denominator+pc.Fees.Channel.Rate <= 0 {
		err := ec.New(ec.NegDenCappedFees, "Negative denominator with capped fees.")
		return 0, err
	}

	channelFees := sellPriceUncappedFees*pc.Fees.Channel.Rate + pc.Fees.Channel.Fixed

	if channelFees < pc.Fees.Channel.Fixed || channelFees > pc.Fees.Channel.CappedValue {
		channelFees = pc.Fees.Channel.CappedValue
	}

	sellPriceCappedFees := (numerator + channelFees) / (denominator + pc.Fees.Channel.Rate)

	return sellPriceCappedFees, nil
}

// GetFeesTotal calculates and returns the total fees
func (pc *PriceCalculator) GetFeesTotal() float64 {
	feesTotal := pc.GetSalesTaxFeesTotal()
	feesTotal += pc.GetChannelFeesTotal()
	feesTotal += pc.GetPaymentFeesTotal()
	feesTotal += pc.GetOtherFeesTotal()
	return feesTotal
}

// GetSalesTaxFeesTotal calculates and returns the sales tax
func (pc *PriceCalculator) GetSalesTaxFeesTotal() float64 {
	return pc.SellPrice * (1 - (1 / (1 + pc.Fees.SalesTax.Rate)))
}

// GetPaymentFeesTotal calculates and returns the payments fees
func (pc *PriceCalculator) GetPaymentFeesTotal() float64 {
	return pc.SellPrice*pc.Fees.Payment.Rate + pc.Fees.Payment.Fixed
}

// GetChannelFeesTotal calculates and returns the channel fees
func (pc *PriceCalculator) GetChannelFeesTotal() float64 {
	channelFeesUncapped := pc.SellPrice*pc.Fees.Channel.Rate + pc.Fees.Channel.Fixed
	channelFees := channelFeesUncapped

	if pc.Fees.Channel.IsCapped &&
		pc.Fees.Channel.CappedValue != 0 &&
		channelFeesUncapped > pc.Fees.Channel.CappedValue {
		channelFees = pc.Fees.Channel.CappedValue
	}

	return channelFees
}

// GetOtherFeesTotal calculates and returns the other fees
func (pc *PriceCalculator) GetOtherFeesTotal() float64 {
	return pc.SellPrice*pc.Fees.Other.Rate + pc.Fees.Other.Fixed
}

// GetProfitTotal calculates and returns the profit total
func (pc *PriceCalculator) GetProfitTotal() float64 {
	profitTotal := pc.SellPrice
	profitTotal -= pc.GetFeesTotal()
	profitTotal -= pc.Cost
	profitTotal -= pc.FreeDeliveryPrice
	return profitTotal
}

// IsValidProfitRate checks if the profit rate is valid
func (pc *PriceCalculator) IsValidProfitRate() bool {
	denominator := (1 / (1 + pc.Fees.SalesTax.Rate))
	denominator -= pc.Fees.Payment.Rate
	denominator -= pc.Fees.Channel.Rate
	denominator -= pc.Fees.Other.Rate
	denominator -= pc.Profit.Rate

	if !pc.Fees.Channel.IsCapped || pc.Fees.Channel.CappedValue == 0 {
		if denominator <= 0 {
			return false
		}
	} else if denominator+pc.Fees.Channel.Rate <= 0 {
		return false
	}

	return true
}
