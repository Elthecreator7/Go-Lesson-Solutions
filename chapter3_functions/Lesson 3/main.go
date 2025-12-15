package main

var basicPriceInDollars int = 100
var premiumPriceInDollars int = 150
var enterprisePriceInDollars int = 500

func getMonthlyPrice(tier string) int {
	var priceInPennies int
	switch tier {
	case "basic":
		priceInPennies = basicPriceInDollars * 100
	case "premium":
		priceInPennies = premiumPriceInDollars * 100
	case "enterprise":
		priceInPennies = enterprisePriceInDollars * 100
	}
	return priceInPennies
}
