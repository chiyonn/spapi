package productpricing

import (
	"github.com/chiyonn/spapi/endpoint/model"
)

type GetPricingParams struct {
	MarketplaceIds string `url:"MarketplaceId"`
	ASINs *[]string `url:"Asins"`
	SKUs *[]string `url:"Skus"`
	ItemType string `url:"ItemType"` // "Asin" or "Sku"
	ItemCondition *string `url:"ItemCondition"` // New, Used, Collectible, Refurblished, Club
	OfferType *string `url:"OfferType"` // B2C, B2B
}

type GetPricingResponse struct {
	Payload    *[]GetPricingResult `json:"payload"`
	Errors     []*model.Error               `json:"errors"`
}

type GetPricingResult struct {
	Status string
	SellerSKU *string
	ASIN *string
	Product *Product
}

type Product struct {
	Identifiers        Identifires         `json:"Identifiers"`
	AttributeSets      []any               `json:"AttributeSets"`
	Relationships      []any               `json:"Relationships"`
	CompetitivePricing *CompetitivePricing `json:"CompetitivePricing"`
	SalesRankings      []SalesRanking      `json:"SalesRankings"`
	Offers             []Offer             `json:"Offers"`
}

type Identifires struct {
	MarketplaceASIN MarketplaceASIN `json:"MarketplaceASIN"`
	SKUIdentifier *SKUIdentifier `json:"SKUIdentifier"`
}

type MarketplaceASIN struct {
	MarketplaceId string `json:"MarketplaceId"`
	ASIN string `json:"ASIN"`
}

type SKUIdentifier struct {
	MarketplaceId string `json:"MarketplaceId"`
	SellerId string `json:"SellerId"`
	SellerSKU string `json:"SellerSKU"`
}

type CompetitivePricing struct {
	CompetitivePrices    []CompetitivePrice    `json:"CompetitivePrices"`
	NumberOfOfferListings []OfferListingCount   `json:"NumberOfOfferListings"`
	TradeInValue         *MoneyType            `json:"TradeInValue"`
}

type CompetitivePrice struct {
	CompetitivePriceId   string     `json:"CompetitivePriceId"`
	Price                PriceType  `json:"Price"`
	Condition            string     `json:"condition"`
	Subcondition         string     `json:"subcondition"`
	OfferType            string     `json:"offerType"`
	QuantityTier         int        `json:"quantityTier"`
	QuantityDiscountType string     `json:"quantityDiscountType"`
	SellerId             string     `json:"sellerId"`
	BelongsToRequester   bool       `json:"belongsToRequester"`
}

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

type OfferListingCount struct {
	Count     int    `json:"Count"`
	Condition string `json:"condition"`
}

type SalesRanking struct {
	ProductCategoryId string `json:"ProductCategoryId"`
	Rank              int    `json:"Rank"`
}

type Offer struct {
	OfferType              string               `json:"offerType"`
	BuyingPrice            PriceType            `json:"BuyingPrice"`
	RegularPrice           MoneyType            `json:"RegularPrice"`
	BusinessPrice          *MoneyType           `json:"businessPrice"`
	QuantityDiscountPrices []QuantityDiscount   `json:"quantityDiscountPrices"`
	FulfillmentChannel     string               `json:"FulfillmentChannel"`
	ItemCondition          string               `json:"ItemCondition"`
	ItemSubCondition       string               `json:"ItemSubCondition"`
	SellerSKU              string               `json:"SellerSKU"`
}

type QuantityDiscount struct {
	QuantityTier         int       `json:"quantityTier"`
	QuantityDiscountType string    `json:"quantityDiscountType"`
	ListingPrice         MoneyType `json:"listingPrice"`
}
