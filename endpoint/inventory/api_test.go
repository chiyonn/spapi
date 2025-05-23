package inventory_test

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chiyonn/spapi/endpoint/inventory"
	"github.com/chiyonn/spapi/testutil"
)

func loadResponseJSON(t *testing.T, name string) string {
	t.Helper()

	path := filepath.Join("testdata", name)
	bytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read test data file %s: %v", path, err)
	}
	return string(bytes)
}

func loadResponseStruct[T any](t *testing.T, name string) T {
	t.Helper()
	path := filepath.Join("testdata", name)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read test JSON: %v", err)
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("failed to unmarshal test JSON: %v", err)
	}
	return result
}

func TestGetInventorySummaries_Success(t *testing.T) {
	body := loadResponseJSON(t, "get_inventory_summary_response.json")
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
	got, err := api.GetInventorySummaries(params)

	assert.NoError(t, err)

	expected := loadResponseStruct[*inventory.GetInventorySummariesResponse](t, "get_inventory_summary_response.json")
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
	_, err := api.GetInventorySummaries(params)

	assert.Error(t, err) // JSON デコード失敗を期待
}
