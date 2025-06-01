package inventory

import (
	"time"

	"github.com/chiyonn/spapi/endpoint/types"
)

type GetInventorySummariesParams struct {
	Details         bool      `url:"details"`
	GranularityType string    `url:"granularityType"`
	GranularityId   string    `url:"granularityId"`
	StartDateTime   time.Time `url:"startDateTime"`
	SellerSkus      []string  `url:"sellerSkus"`
	SellerSku       string    `url:"sellerSku"`
	NextToken       string    `url:"nextToken"`
	MarketplaceIds  []string  `url:"marketplaceIds"`
}

type GetInventorySummariesResponse struct {
	Payload    GetInventorySummariesResult `json:"payload"`
	Pagination types.Pagination            `json:"pagination"`
	Errors     []types.Error               `json:"errors"`
}

type GetInventorySummariesResult struct {
	Granularity        types.Granularity  `json:"granularity"`
	InventorySummaries []InventorySummary `json:"inventorySummaries"`
}

type InventorySummary struct {
	ASIN             string             `json:"asin"`
	FNSKU            string             `json:"fnSku"`
	SellerSKU        string             `json:"sellerSku"`
	Condition        string             `json:"condition"`
	InventoryDetails InventoryDetails   `json:"inventoryDetails"`
	LastUpdatedTime  types.NullableTime `json:"lastUpdatedTime"`
	ProductName      string             `json:"productName"`
	TotalQuantity    int                `json:"totalQuantity"`
	Stores           []string           `json:"stores"`
}

type InventoryDetails struct {
	FulfillableQuantity      int                         `json:"fulfillableQuantity"`
	InboundWorkingQuantity   int                         `json:"inboundWorkingQuantity"`
	InboundShippedQuantity   int                         `json:"inboundShippedQuantity"`
	InboundReceivingQuantity int                         `json:"inboundReceivingQuantity"`
	ReservedQuantity         types.ReservedQuantity      `json:"reservedQuantity"`
	ResearchingQuantity      types.ResearchingQuantity   `json:"researchingQuantity"`
	UnfulfillableQuantity    types.UnfulfillableQuantity `json:"unfulfillableQuantity"`
}
