package pricecalculator

import (
	"testing"
)

type SetExpected struct {
	set      PriceCalculator
	expected float64
}

type SetExpectedBool struct {
	set      PriceCalculator
	expected bool
}

func Set1() *PriceCalculator {
	pc := New()
	pc.SellPrice = 2500
	pc.SalesTaxFeeRate = 0.20
	pc.FreeDeliveryPrice = 0
	pc.Cost = 100
	pc.PaymentFeeRate = 0.034
	pc.PaymentFixedFee = 0.20
	pc.ChannelFeeRate = 0.09
	pc.ChannelFeeIsCapped = true
	pc.ChannelFeeCappedValue = 20
	pc.SelectProfitRate = 0.79933
	return pc
}

func TestGetSellPriceByProfitRate(t *testing.T) {
	pc := Set1()

	tcs := []SetExpected{
		{*pc, 36059999.999463506},
	}

	for _, tc := range tcs {
		result, err := pc.GetSellPriceByProfitRate()

		if err != nil {
			t.Errorf(err.Error())
		}

		if result != tc.expected {
			t.Errorf("Expected sell price of %f got %f.", tc.expected, result)
		}
	}
}

func TestGetFeesTotal(t *testing.T) {
	pc := Set1()

	tcs := []SetExpected{
		{*pc, 521.8666666666666},
	}

	for _, tc := range tcs {
		result := pc.GetFeesTotal()
		if result != tc.expected {
			t.Errorf("Expected total fees of %f got %f.", tc.expected, result)
		}
	}
}

func TestGetSalesTaxFeesTotal(t *testing.T) {
	pc := Set1()

	tcs := []SetExpected{
		{*pc, 416.6666666666666},
	}

	for _, tc := range tcs {
		result := pc.GetSalesTaxFeesTotal()
		if result != tc.expected {
			t.Errorf("Expected sales tax fees of %f got %f.", tc.expected, result)
		}
	}
}

func TestGetPaymentFeesTotal(t *testing.T) {
	pc := Set1()

	tcs := []SetExpected{
		{*pc, 85.2},
	}

	for _, tc := range tcs {
		result := pc.GetPaymentFeesTotal()
		if result != tc.expected {
			t.Errorf("Expected payment fees of %f got %f.", tc.expected, result)
		}
	}
}

func TestGetChannelFeesTotal(t *testing.T) {
	pc := Set1()

	tcs := []SetExpected{
		{*pc, 20},
	}

	for _, tc := range tcs {
		result := pc.GetChannelFeesTotal()
		if result != tc.expected {
			t.Errorf("Expected channel fees of %f got %f.", tc.expected, result)
		}
	}
}

func TestGetOtherFeesTotal(t *testing.T) {
	pc := Set1()

	tcs := []SetExpected{
		{*pc, 0},
	}

	for _, tc := range tcs {
		result := pc.GetOtherFeesTotal()
		if result != tc.expected {
			t.Errorf("Expected other fees of %f got %f.", tc.expected, result)
		}
	}
}

func TestGetProfitTotal(t *testing.T) {
	pc := Set1()

	tcs := []SetExpected{
		{*pc, 1878.1333333333334},
	}

	for _, tc := range tcs {
		result := pc.GetProfitTotal()
		if result != tc.expected {
			t.Errorf("Expected total profit of %f got %f.", tc.expected, result)
		}
	}
}

func TestIsValidProfitRate(t *testing.T) {
	pc := Set1()

	tcs := []SetExpectedBool{
		{*pc, true},
	}

	for _, tc := range tcs {
		result := pc.IsValidProfitRate()
		if result != tc.expected {
			t.Errorf("Expected profit rate validity of %t got %t.", tc.expected, result)
		}
	}
}
