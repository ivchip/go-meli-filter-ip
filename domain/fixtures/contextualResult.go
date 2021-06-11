package fixtures

import "github.com/ivchip/go-meli-filter-ip/domain"

func GenerateIP() string {
	return "170.78.40.99"
}

func GenerateContextualResult() domain.ContextualResult {
	return domain.ContextualResult{
		CurrencyQuote: 3606,
	}
}
