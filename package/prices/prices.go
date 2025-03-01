package prices

import (
	"bufio"
	"fmt"

	"example.com/tax-calculator/package/conversion"
	"example.com/tax-calculator/package/fileops"
)

type TaxIncludedPriceJobs struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJobs(taxRate float64) *TaxIncludedPriceJobs {
	return &TaxIncludedPriceJobs{
		TaxRate: taxRate,
	}
}

func (jobs *TaxIncludedPriceJobs) LoadData() {
	file, err := fileops.OpenFile("prices.txt")
	if err != nil {
		panic(err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readPrice, err := conversion.StringToFloat(scanner.Text())
		if err != nil {
			panic(err.Error())
		}
		jobs.InputPrices = append(jobs.InputPrices, readPrice)
	}
	file.Close()
}

func (job *TaxIncludedPriceJobs) CalculateTaxes() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%0.2f", price)] = fmt.Sprintf("%0.2f", taxIncludePrice)
	}

	job.TaxIncludedPrices = result
	fileops.WriteToJson(fmt.Sprintf("result_%0.0f.json", job.TaxRate*100), job)
}
