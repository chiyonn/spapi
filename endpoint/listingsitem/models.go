package listingsitem

import "github.com/chiyonn/spapi/endpoint/types"

type PatchListingsItemQuery struct {
	MarketplaceIds []string                 `json:"marketplaceIds"`
	IncludedData   []string                 `json:"includedData"`
	Mode           string                   `json:"mode"`
	IssueLocale    string                   `json:"issueLocale"`
	Body           ListingsItemPatchRequest `json:"body"`
}

type ListingsItemPatchRequest struct {
	ProductType string           `json:"productType"`
	Patches     []PatchOperation `json:"patches"`
}

type PatchOperation struct {
	OP    string `json:"op"`
	Path  string `json:"path"`
	Value []any  `json:"value"`
}

type ListingsItemSubmissionResponse struct {
	SKU          string        `json:"sku"`
	Status       string        `json:"status"`
	SubmissionID string        `json:"submissionId"`
	Issues       []types.Issue `json:"issues"`
	Identifiers  []string      `json:"identifiers"`
}
