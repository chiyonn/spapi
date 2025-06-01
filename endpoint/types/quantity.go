package types

type UnfulfillableQuantity struct {
	TotalUnfulfillableQuantity int `json:"totalUnfulfillableQuantity"`
	CustomerDamagedQuantity    int `json:"customerDamagedQuantity"`
	WarehouseDamagedQuantity   int `json:"warehouseDamagedQuantity"`
	DistributorDamagedQuantity int `json:"distributorDamagedQuantity"`
	CarrierDamagedQuantity     int `json:"carrierDamagedQuantity"`
	DefectiveQuantity          int `json:"defectiveQuantity"`
	ExpiredQuantity            int `json:"expiredQuantity"`
}

type ResearchingQuantity struct {
	TotalResearchingQuantity     int                       `json:"totalResearchingQuantity"`
	ResearchingQuantityBreakdown []ResearchingQuantityEntry `json:"researchingQuantityBreakdown"`
}


type ReservedQuantity struct {
	TotalReservedQuantity        int `json:"totalReservedQuantity"`
	PendingCustomerOrderQuantity int `json:"pendingCustomerOrderQuantity"`
	PendingTransshipmentQuantity int `json:"pendingTransshipmentQuantity"`
	FCProcessingQuantity         int `json:"fcProcessingQuantity"`
}

type ResearchingQuantityEntry struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type QuantityDiscount struct {
	QuantityTier         int       `json:"quantityTier"`
	QuantityDiscountType string    `json:"quantityDiscountType"`
	ListingPrice         MoneyType `json:"listingPrice"`
}
