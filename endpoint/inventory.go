package endpoint

import (
	"context"
	"net/http"
	"encoding/json"

	"github.com/chiyonn/spapi/client"
)

type InventoryAPI struct {
	client *client.Client
}

func NewInventoryAPI(client *client.Client) *InventoryAPI {
	return &InventoryAPI{client: client}
}

func (api *InventoryAPI) GetInventorySummary() (any, error) {
	const path = "/fba/inventory/v1/summaries"

	endpoint, err := NewEndpoint(api.client, http.MethodGet, path, 2, 2, "inventory.GetInventorySummary")
	if err != nil {
		return nil, err
	}

	endpoint.BuildReq = func() (*http.Request, error) {
		return http.NewRequest(http.MethodGet, api.client.BaseURL+path, nil)
	}

	endpoint.ParseResp = func(resp *http.Response) (any, error) {
		var res map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return nil, err
		}
		return res, nil
	}

	return endpoint.Do(context.Background())
}
