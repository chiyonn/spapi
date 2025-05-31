package inventory_test

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"context"

	"github.com/stretchr/testify/assert"

	"github.com/chiyonn/spapi/endpoint/inventory"
	"github.com/chiyonn/spapi/testutil"
)

func TestGetInventorySummaries_Success(t *testing.T) {
	body := testutil.LoadResponseJSON(t, "get_inventory_summary_response.json")
	client := testutil.NewMockedClient(t, func(req *http.Request) *http.Response {
		assert.Equal(t, "/fba/inventory/v1/summaries", req.URL.Path)

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}
	})

	params := &inventory.GetInventorySummariesParams{
		GranularityType: "Marketplace",
		GranularityId:   "maketplace_id",
	}

	api := inventory.NewInventoryAPI(client)
	got, err := api.GetInventorySummaries(context.Background(), params)

	assert.NoError(t, err)

	expected := testutil.LoadResponseStruct[*inventory.GetInventorySummariesResponse](t, "get_inventory_summary_response.json")
	assert.Equal(t, expected, got)
}

func TestGetInventorySummaries_BadJSON(t *testing.T) {
	client := testutil.NewMockedClient(t, func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`this is NOT json`)),
			Header:     make(http.Header),
		}
	})

	params := &inventory.GetInventorySummariesParams{
		GranularityType: "Marketplace",
		GranularityId:   "maketplace_id",
	}

	api := inventory.NewInventoryAPI(client)
	_, err := api.GetInventorySummaries(context.Background(), params)

	assert.Error(t, err) // JSON デコード失敗を期待
}
