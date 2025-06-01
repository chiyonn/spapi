package types

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


type NumberOfOffer struct {
	Condition string `json:"Condition"`
	FulfillmentChannel string `json:"fulfillmentChannel"`
	OfferCount int32 `json:"OfferCount"`
}

type ItemOffer struct {
	MyOffer bool `json:"MyOffer"`
	OfferType string `json:"OfferType"`
	SubCondition string `json:"SubCondition"`
	SellerId string `json:"SellerID"`
	ConditionNotes string `json:"ConditionNotes"`
	SellerFeedbackRating SellerFeedbackRating `json:"SellerFeedbackRating"`
}
