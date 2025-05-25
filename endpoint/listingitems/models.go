package listingitems

import "time"


type PatchListingItemQuery struct {
	MarketplaceIds []string `json:"marketplaceIds"`
	IncludedData *[]string `json:"includedData"`
	Mode *string `json:"mode"`
	IssueLocale *string `json:"issueLocale"`
	Body *ListingItemPatchRequest `json:"body"`
}

type ListingItemPatchRequest struct {
	ProductType string `json:"productType"`
	Patches []PatchOperation `json:"patches"`
}

type PatchOperation struct {
	OP string `json:"op"`
	Path string `json:"path"`
	Value *[]any `json:"value"`
}

type ListingItemSubmissionResponse struct {
	SKU string `json:"sku"`
	Status string `json:"status"`
	SubmissionID string `json:"submissionId"`
	Issues *[]Issue `json:"issues"`
	Identifiers *[]string `json:"identifiers"`
}

type Issue struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Severity string `json:"severity"`
	AttributeNames *[]string `json:"attributeNames"`
	Categories []string `json:"categories"`
	Enforcements *[]IssueEnforcements
}

type IssueEnforcements struct {
	Actions []IssueEnforcementAction `json:"actions"`
	Exemption IssueExemption
}

type IssueEnforcementAction struct {
	Action string `json:"action"`
}

type IssueExemption struct {
	Status string `json:"status"`
	ExpiryDate *time.Time `json:"expiryDate"`
}
