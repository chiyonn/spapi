package productpricing

import (
	"github.com/chiyonn/spapi/endpoint/types"
)

type GetPricingParams struct {
	MarketplaceIds string `url:"MarketplaceId"`
	ASINs []string `url:"Asins"`
	SKUs []string `url:"Skus"`
	ItemType string `url:"ItemType"` // "Asin" or "Sku"
	ItemCondition string `url:"ItemCondition"` // New, Used, Collectible, Refurblished, Club
	OfferType string `url:"OfferType"` // B2C, B2B
}

type GetItemOffersBatchParams struct {
	Requests []GetListingOffersParams `json:"requests"`
}

type GetListingOffersParams struct {
	URI string `json:"uri"` // e.g. "/products/pricing/v0/{items|listings}/B000P6Q7MY/offers"
	Method string `json:"method"`
	Header string `json:"header"` 
	MarketplaceID string `json:"MarketplaceId"`
	ItemCondition string `json:"ItemCondition"` // New, Used, Collectible, Refurblished, Club
	CustomerType string `json:"CustomerType"` // Consumer, Business
}

type GetPricingResponse struct {
	Payload    []GetPricingResult `json:"payload"`
	Errors     []types.Error               `json:"errors"`
}

type GetItemOffersBatchResponse struct {
	Responses []GetItemOffersResponse `json:"responses"`
}

type GetItemOffersResponse struct {
	Headers string `json:"headers"`
	Status types.Status `json:"status"`
	Body GetItemOffersResult `json:"body"`
}

type GetPricingResult struct {
	Status string
	SellerSKU string
	ASIN string
	Product Product
}

type GetItemOffersResult struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	MarketplaceID string `json:"MarketplaceID"`
	ASIN string `json:"ASIN"`
	SKU string `json:"SKU"`
	ItemCondition string `json:"ItemCondition"`
	Status string `json:"status"`
	Identifier types.Identifier `json:"Identifier"`
}

type Product struct {
	Identifiers        types.Identifires         `json:"Identifiers"`
	AttributeSets      []any               `json:"AttributeSets"`
	Relationships      []any               `json:"Relationships"`
	CompetitivePricing CompetitivePricing `json:"CompetitivePricing"`
	SalesRankings      []SalesRanking      `json:"SalesRankings"`
	Offers             []types.Offer             `json:"Offers"`
}

type CompetitivePricing struct {
	CompetitivePrices    []CompetitivePrice    `json:"CompetitivePrices"`
	NumberOfOfferListings []OfferListingCount   `json:"NumberOfOfferListings"`
	TradeInValue         types.MoneyType            `json:"TradeInValue"`
}

type CompetitivePrice struct {
	CompetitivePriceId   string     `json:"CompetitivePriceId"`
	Price                types.PriceType  `json:"Price"`
	Condition            string     `json:"condition"`
	Subcondition         string     `json:"subcondition"`
	OfferType            string     `json:"offerType"`
	QuantityTier         int        `json:"quantityTier"`
	QuantityDiscountType string     `json:"quantityDiscountType"`
	SellerId             string     `json:"sellerId"`
	BelongsToRequester   bool       `json:"belongsToRequester"`
}

type OfferListingCount struct {
	Count     int    `json:"Count"`
	Condition string `json:"condition"`
}

type SalesRanking struct {
	ProductCategoryId string `json:"ProductCategoryId"`
	Rank              int    `json:"Rank"`
}
