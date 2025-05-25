package inventory

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chiyonn/spapi/client"
	"github.com/chiyonn/spapi/endpoint"
	"github.com/google/go-querystring/query"
)

type InventoryAPI struct {
	client *client.Client
}

func NewInventoryAPI(client *client.Client) *InventoryAPI {
	return &InventoryAPI{client: client}
}

func (api *InventoryAPI) GetInventorySummaries(params *GetInventorySummariesParams) (*GetInventorySummariesResponse, error) {
	const rate = 2.0
	const burst = 2
	const path = "/fba/inventory/v1/summaries"
	const key = "inventory.GetInventorySummaries"
	const method = http.MethodGet

	endpoint, err := endpoint.NewEndpoint(api.client, method, path, rate, burst, key)
	if err != nil {
		return nil, err
	}

	endpoint.BuildReq = func() (*http.Request, error) {
		req, err := http.NewRequest(method, api.client.BaseURL+path, nil)
		if err != nil {
			return nil, err
		}

		if params != nil {
			values, err := query.Values(params)
			if err != nil {
				return nil, fmt.Errorf("failed to encode query params: %w", err)
			}
			req.URL.RawQuery = values.Encode()
		}

		return req, nil
	}

	endpoint.ParseResp = func(resp *http.Response) (any, error) {
		var res GetInventorySummariesResponse
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return nil, err
		}
		return &res, nil
	}

	result, err := endpoint.Do(context.Background())
	if err != nil {
		return nil, err
	}

	resp, ok := result.(*GetInventorySummariesResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", result)
	}
	return resp, nil
}
