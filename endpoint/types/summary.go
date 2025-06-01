package types

type Summary struct {
	TotalOfferCount int32 `json:"TotalOfferCount"`
	NumberOfOffers []NumberOfOffer `json:"NubmerOfOffers"`
	LowestPrices []LowestPrice `json:"LowestPrices"`
	Offers []ItemOffer `json:"Offers"`
}


