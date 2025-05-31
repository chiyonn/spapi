package productpricing

import (
	"context"

	"github.com/chiyonn/spapi/client"
	"github.com/chiyonn/spapi/endpoint"
)

type ProductPricingAPI struct {
	client *client.Client
}

func NewProductPricingAPI(c *client.Client) *ProductPricingAPI {
	return &ProductPricingAPI{client: c}
}

func (api *ProductPricingAPI) GetPricing(
	ctx context.Context,
	params *GetPricingParams,
) (*GetPricingResponse, error) {
	const (
		rate  = 0.5
		burst = 1
		key   = "productpricing.GetPricing"
		path  = "/products/pricing/v0/price"
	)

	ep, err := endpoint.NewJSONGet[GetPricingParams, GetPricingResponse](
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
