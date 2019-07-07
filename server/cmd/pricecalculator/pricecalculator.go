package pricecalculator

import (
	ec "../errors"
)

// PriceCalculator stores the price, costs, and fees
type PriceCalculator struct {
	SellPrice             float64 `json:"sell_price"`
	FreeDeliveryPrice     float64 `json:"free_delivery_price"`
	Cost                  float64 `json:"cost"`
	SalesTaxFeeRate       float64 `json:"fees.sales_tax.rate"`
	PaymentFeeRate        float64 `json:"fees.payment.rate"`
	PaymentFixedFee       float64 `json:"fees.payment.fixed"`
	ChannelFeeRate        float64 `json:"fees.channel.rate"`
	ChannelFixedFee       float64 `json:"fees.channel.fixed"`
	ChannelFeeIsCapped    bool    `json:"fees.channel.is_capped"`
	ChannelFeeCappedValue float64 `json:"fees.channel.capped_value"`
	OtherFeeRate          float64 `json:"fees.other.rate"`
	OtherFixedFee         float64 `json:"fees.other.fixed"`
	SelectProfitRate      float64 `json:"profit.rate"`
}

// New creates and returns a new instance of PriceCalculator
func New() *PriceCalculator {
	return &PriceCalculator{}
}

// GetSellPriceByProfitRate calculates and returns the sell price given the profit rate
func (pc *PriceCalculator) GetSellPriceByProfitRate() (float64, error) {
	numerator := pc.Cost
	numerator += pc.FreeDeliveryPrice
	numerator += pc.PaymentFixedFee
	numerator += pc.ChannelFixedFee
	numerator += pc.OtherFixedFee

	denominator := (1 / (1 + pc.SalesTaxFeeRate))
	denominator -= pc.PaymentFeeRate
	denominator -= pc.ChannelFeeRate
	denominator -= pc.OtherFeeRate
	denominator -= pc.SelectProfitRate

	sellPriceUncappedFees := numerator / denominator

	if !pc.ChannelFeeIsCapped || pc.ChannelFeeCappedValue == 0 {
		if denominator <= 0 {
			err := ec.New(ec.NegDenUncappedFees, "Negative denomiator with uncapped fees.")
			return 0, err
		}
		return sellPriceUncappedFees, nil
	} else if denominator+pc.ChannelFeeRate <= 0 {
		err := ec.New(ec.NegDenCappedFees, "Negative denominator with capped fees.")
		return 0, err
	}

	channelFees := sellPriceUncappedFees*pc.ChannelFeeRate + pc.ChannelFixedFee

	if channelFees < pc.ChannelFixedFee || channelFees > pc.ChannelFeeCappedValue {
		channelFees = pc.ChannelFeeCappedValue
	}

	sellPriceCappedFees := (numerator + channelFees) / (denominator + pc.ChannelFeeRate)

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
	return pc.SellPrice * (1 - (1 / (1 + pc.SalesTaxFeeRate)))
}

// GetPaymentFeesTotal calculates and returns the payments fees
func (pc *PriceCalculator) GetPaymentFeesTotal() float64 {
	return pc.SellPrice*pc.PaymentFeeRate + pc.PaymentFixedFee
}

// GetChannelFeesTotal calculates and returns the channel fees
func (pc *PriceCalculator) GetChannelFeesTotal() float64 {
	channelFeesUncapped := pc.SellPrice*pc.ChannelFeeRate + pc.ChannelFixedFee
	channelFees := channelFeesUncapped

	if pc.ChannelFeeIsCapped &&
		pc.ChannelFeeCappedValue != 0 &&
		channelFeesUncapped > pc.ChannelFeeCappedValue {
		channelFees = pc.ChannelFeeCappedValue
	}

	return channelFees
}

// GetOtherFeesTotal calculates and returns the other fees
func (pc *PriceCalculator) GetOtherFeesTotal() float64 {
	return pc.SellPrice*pc.OtherFeeRate + pc.OtherFixedFee
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
	denominator := (1 / (1 + pc.SalesTaxFeeRate))
	denominator -= pc.PaymentFeeRate
	denominator -= pc.ChannelFeeRate
	denominator -= pc.OtherFeeRate
	denominator -= pc.SelectProfitRate

	if !pc.ChannelFeeIsCapped || pc.ChannelFeeCappedValue == 0 {
		if denominator <= 0 {
			return false
		}
	} else if denominator+pc.ChannelFeeRate <= 0 {
		return false
	}

	return true
}
