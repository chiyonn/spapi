package listingsitem_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/chiyonn/spapi/endpoint/listingsitem"
	"github.com/chiyonn/spapi/testutil"
	"github.com/stretchr/testify/assert"
)

func TestPatchListingsItem(t *testing.T) {
	body := testutil.LoadResponseJSON(t, "patch_listing_items_response.json")
	client := testutil.NewMockedClient(t, func(req *http.Request) *http.Response {
		assert.Equal(t, "/listings/2021-09-01/items/AAA/BBB", req.URL.Path)
		assert.Equal(t, http.MethodPatch, req.Method)

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}
	})

	params := &listingsitem.PatchListingsItemQuery{
		MarketplaceIds: []string{"marketplaceid"},
	}

	api := listingsitem.NewListingsItemsAPI(client)
	got, err := api.PatchListingsItem("AAA", "BBB", params)
	assert.NoError(t, err)

	expected := testutil.LoadResponseStruct[*listingsitem.ListingsItemSubmissionResponse](t, "patch_listing_items_response.json")
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}
