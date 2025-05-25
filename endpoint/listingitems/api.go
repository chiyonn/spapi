package listingitems

import (
	"context"
	"encoding/json"
	"bytes"
	"fmt"
	"net/url"
	"net/http"

	"github.com/chiyonn/spapi/client"
	"github.com/chiyonn/spapi/endpoint"
)

type ListingItemsAPI struct {
	client *client.Client
}

func NewListingItemAPI(client *client.Client) *ListingItemsAPI {
	return &ListingItemsAPI{client: client}
}

func (api *ListingItemsAPI) PatchListingItem(sellerID string, sku string, params *PatchListingItemQuery) (*ListingItemSubmissionResponse, error) {
	const rate = 5.0
	const burst = 5
	path := fmt.Sprintf("/listings/2021-08-01/items/%s/%s", sellerID, sku)
	const key = "inventory.GetInventorySummaries"
	const method = http.MethodPatch

	endpoint, err := endpoint.NewEndpoint(api.client, method, path, rate, burst, key)
	if err != nil {
		return nil, err
	}

	endpoint.BuildReq = func() (*http.Request, error) {
		u, err := url.Parse(api.client.BaseURL + path)
		if err != nil {
			return nil, err
		}
		q := u.Query()
		for _, v := range params.MarketplaceIds {
			q.Add("marketplaceIds", v)
		}
		if params.IncludedData != nil {
			for _, v := range *params.IncludedData {
				q.Add("includedData", v)
			}
		}
		if params.Mode != nil {
			q.Set("mode", *params.Mode)
		}
		if params.IssueLocale != nil {
			q.Set("issueLocale", *params.IssueLocale)
		}
		u.RawQuery = q.Encode()

		var buf bytes.Buffer
		if params.Body != nil {
			if err := json.NewEncoder(&buf).Encode(params.Body); err != nil {
				return nil, err
			}
		}

		req, err := http.NewRequest(method, u.String(), &buf)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		return req, nil
	}

	endpoint.ParseResp = func(resp *http.Response) (any, error) {
		var res ListingItemSubmissionResponse
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return nil, err
		}
		return &res, nil
	}

	result, err := endpoint.Do(context.Background())
	if err != nil {
		return nil, err
	}

	resp, ok := result.(*ListingItemSubmissionResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", result)
	}
	return resp, nil
}
