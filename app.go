package main

import (
	"example.com/tax-calculator/package/prices"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, value := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJobs(value)
		priceJob.CalculateTaxes()
	}

}
