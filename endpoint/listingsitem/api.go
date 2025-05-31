package listingsitem

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/chiyonn/spapi/client"
	"github.com/chiyonn/spapi/endpoint"
)

type ListingsItemsAPI struct {
	client *client.Client
}

func NewListingsItemsAPI(client *client.Client) *ListingsItemsAPI {
	return &ListingsItemsAPI{client: client}
}

func (api *ListingsItemsAPI) PatchListingsItem(ctx context.Context, sellerID string, sku string, params *PatchListingsItemQuery) (*ListingsItemSubmissionResponse, error) {
	const rate = 5.0
	const burst = 5
	path := fmt.Sprintf("/listings/2021-09-01/items/%s/%s", sellerID, sku)
	const key = "listingsitem.PatchListingsItem"
	const method = http.MethodPatch

	endpoint, err := endpoint.NewEndpoint(api.client, method, path, rate, burst, key)
	if err != nil {
		return nil, fmt.Errorf("failed to build new endpoint: %w", err)
	}

	endpoint.BuildReq = func() (*http.Request, error) {
		u, err := url.Parse(api.client.BaseURL + path)
		if err != nil {
			return nil, fmt.Errorf("failed to parse url: %w", err)
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
				return nil, fmt.Errorf("failed to parse request body: %w", err)
			}
		}

		req, err := http.NewRequest(method, u.String(), &buf)
		if err != nil {
			return nil, fmt.Errorf("failed to build new request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")
		return req, nil
	}

	endpoint.ParseResp = func(resp *http.Response) (any, error) {
		var res ListingsItemSubmissionResponse
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return nil, fmt.Errorf("failed to decode response body: %w", err)
		}
		return &res, nil
	}

	result, err := endpoint.Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute endpoint: %w", err)
	}

	resp, ok := result.(*ListingsItemSubmissionResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", result)
	}
	return resp, nil
}
