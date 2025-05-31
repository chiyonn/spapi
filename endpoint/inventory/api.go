// inventory/inventory.go
package inventory

import (
	"context"

	"github.com/chiyonn/spapi/client"
	"github.com/chiyonn/spapi/endpoint"
)

type InventoryAPI struct {
	client *client.Client
}

func NewInventoryAPI(c *client.Client) *InventoryAPI {
	return &InventoryAPI{client: c}
}

func (api *InventoryAPI) GetInventorySummaries(
	ctx context.Context,
	params *GetInventorySummariesParams,
) (*GetInventorySummariesResponse, error) {
	const (
		path  = "/fba/inventory/v1/summaries"
		key   = "inventory.GetInventorySummaries"
		rate  = 2.0
		burst = 2
	)

	ep, err := endpoint.NewJSONGet[GetInventorySummariesParams, GetInventorySummariesResponse](
		api.client,
		path,
		key,
		rate,
		burst,
	)
	if err != nil {
		return nil, err
	}
	return ep.Do(ctx, params)
}
