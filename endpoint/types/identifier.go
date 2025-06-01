package types

type Identifier struct {
	MarketplaceID string  `json:"MarketplaceId"`
	ASIN          string  `json:"ASIN"`
	SellerSKU     string  `json:"SellerSKU"`
	ItemCondition string  `json:"ItemCondition"`
	Summary       Summary `json:"summary"`
}

type Identifires struct {
	MarketplaceASIN MarketplaceASIN `json:"MarketplaceASIN"`
	SKUIdentifier   SKUIdentifier   `json:"SKUIdentifier"`
}

type MarketplaceASIN struct {
	MarketplaceId string `json:"MarketplaceId"`
	ASIN          string `json:"ASIN"`
}

type SKUIdentifier struct {
	MarketplaceId string `json:"MarketplaceId"`
	SellerId      string `json:"SellerId"`
	SellerSKU     string `json:"SellerSKU"`
}

type Granularity struct {
	GranularityType string `json:"granularityType"`
	GranularityId   string `json:"granularityId"`
}
