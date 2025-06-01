package types

type PriceType struct {
	LandedPrice  MoneyType  `json:"LandedPrice"`
	ListingPrice MoneyType  `json:"ListingPrice"`
	Shipping     MoneyType  `json:"Shipping"`
	Points       *Points    `json:"Points"`
}

type MoneyType struct {
	CurrencyCode string  `json:"CurrencyCode"`
	Amount       float64 `json:"Amount"`
}

type Points struct {
	PointsNumber         int       `json:"PointsNumber"`
	PointsMonetaryValue  MoneyType `json:"PointsMonetaryValue"`
}

type LowestPrice struct {
	Condition string `json:"condition"`
	FulfillmentChannel string `json:"fulfillmentChannel"`
	OfferType string `json:"offerType"`
	QuantityTier         int        `json:"quantityTier"`
	QuantityDiscountType string     `json:"quantityDiscountType"`
	PriceType
	Shipping MoneyType `json:"Shipping"`
	Points Points `json:"Points"`
}
