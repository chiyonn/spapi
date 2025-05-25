package listingitems_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/chiyonn/spapi/endpoint/listingitems"
	"github.com/chiyonn/spapi/testutil"
	"github.com/stretchr/testify/assert"
)

func TestPatchListingItem(t *testing.T) {
	body := testutil.LoadResponseJSON(t, "patch_listing_items_response.json")
	client := testutil.NewMockedClient(t, func(req *http.Request) *http.Response {
		assert.Equal(t, "/listings/2021-08-01/items/AAA/BBB", req.URL.Path)
		assert.Equal(t, http.MethodPatch, req.Method)

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}
	})

	params := &listingitems.PatchListingItemQuery{
		MarketplaceIds: []string{"marketplaceid"},
	}

	api := listingitems.NewListingItemAPI(client)
	got, err := api.PatchListingItem("AAA", "BBB", params)
	assert.NoError(t, err)

	expected := testutil.LoadResponseStruct[*listingitems.ListingItemSubmissionResponse](t, "patch_listing_items_response.json")
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}
