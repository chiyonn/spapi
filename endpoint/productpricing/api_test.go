package productpricing_test

import (
	"io"
	"context"
	"testing"
	"strings"
	"net/http"

	"github.com/stretchr/testify/assert"

	"github.com/chiyonn/spapi/endpoint/productpricing"
	"github.com/chiyonn/spapi/testutil"
)

func TestGetPricing_Success(t *testing.T) {
	body := testutil.LoadResponseJSON(t, "get_pricing_response.json")
	client := testutil.NewMockedClient(t, func(req *http.Request) *http.Response {
		assert.Equal(t, "/products/pricing/v0/price", req.URL.Path)

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}
	})

	params := &productpricing.GetPricingParams{
		MarketplaceIds: "marketpkaceid",
		ASINs: &[]string{"ASIN01", "ASIN02"},
		ItemType: "Asin",
	}

	api := productpricing.NewProductPricingAPI(client)
	got, err := api.GetPricing(context.Background(), params)

	assert.NoError(t, err)

	expected := testutil.LoadResponseStruct[*productpricing.GetPricingResponse](t, "get_pricing_response.json")
	assert.Equal(t, expected, got)
}
