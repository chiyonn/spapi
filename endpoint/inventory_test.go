package endpoint_test

import (
	"testing"

	"github.com/chiyonn/spapi/client"
	"github.com/chiyonn/spapi/endpoint"
	"github.com/chiyonn/spapi/testutil"
)

func TestGetInventorySummary(t *testing.T) {
	httpClient := testutil.NewMockHTTPClient(`{"summary": "test summary"}`, 200)

	cli := &client.Client{
		HttpClient:       httpClient,
		BaseURL:          "https://mock-api.amazon.com",
		RateLimitManager: client.NewRateLimitManager(),
	}

	api := endpoint.NewInventoryAPI(cli)
	res, err := api.GetInventorySummary()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	m, ok := res.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected map[string]interface{}, got %T", res)
	}

	if m["summary"] != "test summary" {
		t.Errorf("Expected summary = 'test summary', got %v", m["summary"])
	}
}
