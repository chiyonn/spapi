package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chiyonn/spapi/client"
)

type InventoryAPI struct {
	client *client.Client
}

func NewInventoryAPI(client *client.Client) *InventoryAPI {
	return &InventoryAPI{client: client}
}

func (api *InventoryAPI) GetInventorySummary() (*GetInventorySummariesResponse, error) {
	const path = "/fba/inventory/v1/summaries"

	endpoint, err := NewEndpoint(api.client, http.MethodGet, path, 2, 2, "inventory.GetInventorySummary")
	if err != nil {
		return nil, err
	}

	endpoint.BuildReq = func() (*http.Request, error) {
		return http.NewRequest(http.MethodGet, api.client.BaseURL+path, nil)
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
