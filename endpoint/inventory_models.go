package endpoint

import (
	"time"

	"github.com/chiyonn/spapi/endpoint/internal/model"
)

type GetInventorySummariesParams struct {
	Details         *bool      `query:"details"`
	GranularityType string    `query:"granularityType"`
	GranularityId   string    `query:"granularityId"`
	StartDateTime   *time.Time `query:"startDateTime"`
	SellerSkus      *[]string  `query:"sellerSkus"`
	SellerSku       *string    `query:"sellerSku"`
	NextToken       *string    `query:"nextToken"`
	MarketplaceIds  []string  `query:"marketplaceIds"`
}

//func (p *GetInventorySummariesParams) Stringfy() string {
//	return queryutil.StructToQuery(p).Encode()
//}

type GetInventorySummariesResponse struct {
	Payload    *GetInventorySummariesResult `json:"payload"`
	Pagination *model.Pagination            `json:"pagination"`
	Errors     []*model.Error 				`json:"errors"`
}

type GetInventorySummariesResult struct {
	Granularity        model.Granularity   `json:"granularity"`
	InventorySummaries []InventorySummary  `json:"inventorySummaries"`
}

type InventorySummary struct {
	ASIN             *string            `json:"asin"`
	FNSKU            *string            `json:"fnSku"`
	SellerSKU        *string            `json:"sellerSku"`
	Condition        *string            `json:"condition"`
	InventoryDetails *InventoryDetails  `json:"inventoryDetails"`
	LastUpdatedTime  *time.Time         `json:"lastUpdatedTime"`
	ProductName      *string            `json:"productName"`
	TotalQuantity    *int	            `json:"totalQuantity"`
	Stores           *[]string          `json:"stores"`
}

type InventoryDetails struct {
	FulfillableQuantity      *int                 `json:"fulfillableQuantity"`
	InboundWorkingQuantity   *int                 `json:"inboundWorkingQuantity"`
	InboundShippedQuantity   *int                 `json:"inboundShippedQuantity"`
	InboundReceivingQuantity *int                 `json:"inboundReceivingQuantity"`
	ReservedQuantity         *ReservedQuantity    `json:"reservedQuantity"`
	ResearchingQuantity      *ResearchingQuantity `json:"researchingQuantity"`
	UnfulfillableQuantity    *UnfulfillableQuantity `json:"unfulfillableQuantity"`
}

type ReservedQuantity struct {
	TotalReservedQuantity        *int `json:"totalReservedQuantity"`
	PendingCustomerOrderQuantity *int `json:"pendingCustomerOrderQuantity"`
	PendingTransshipmentQuantity *int `json:"pendingTransshipmentQuantity"`
	FCProcessingQuantity         *int `json:"fcProcessingQuantity"`
}

type ResearchingQuantity struct {
	TotalResearchingQuantity     *int                     `json:"totalResearchingQuantity"`
	ResearchingQuantityBreakdown []ResearchingQuantityEntry `json:"researchingQuantityBreakdown"`
}

type ResearchingQuantityEntry struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type UnfulfillableQuantity struct {
	TotalUnfulfillableQuantity *int `json:"totalUnfulfillableQuantity"`
	CustomerDamagedQuantity    *int `json:"customerDamagedQuantity"`
	WarehouseDamagedQuantity   *int `json:"warehouseDamagedQuantity"`
	DistributorDamagedQuantity *int `json:"distributorDamagedQuantity"`
	CarrierDamagedQuantity     *int `json:"carrierDamagedQuantity"`
	DefectiveQuantity          *int `json:"defectiveQuantity"`
	ExpiredQuantity            *int `json:"expiredQuantity"`
}
