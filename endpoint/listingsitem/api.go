package listingsitem

import (
	"context"
	"fmt"

	"github.com/chiyonn/spapi/client"
	"github.com/chiyonn/spapi/endpoint"
)

type ListingsItemsAPI struct{ client *client.Client }

func NewListingsItemsAPI(c *client.Client) *ListingsItemsAPI { return &ListingsItemsAPI{c} }

func (api *ListingsItemsAPI) PatchListingsItem(
	ctx context.Context,
	sellerID, sku string,
	params *PatchListingsItemQuery,
) (*ListingsItemSubmissionResponse, error) {

	const (
		rate  = 5.0
		burst = 5
		key   = "listingsitem.PatchListingsItem"
	)

	path := fmt.Sprintf("/listings/2021-09-01/items/%s/%s", sellerID, sku)

	ep, err := endpoint.NewJSONPatch[PatchListingsItemQuery, ListingsItemSubmissionResponse](
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
