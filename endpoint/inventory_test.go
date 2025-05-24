// endpoint/inventory_test.go
package endpoint_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chiyonn/spapi/endpoint"
	"github.com/chiyonn/spapi/testutil"
)

func TestGetInventorySummary_Success(t *testing.T) {
	client := testutil.NewMockedClient(t, func(req *http.Request) *http.Response {
		assert.Equal(t, "/fba/inventory/v1/summaries", req.URL.Path)

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`{"summary":"value"}`)),
			Header:     make(http.Header),
		}
	})

	api := endpoint.NewInventoryAPI(client)
	got, err := api.GetInventorySummary()

	assert.NoError(t, err)
	assert.Equal(t,
		map[string]interface{}{"summary": "value"},
		got,
	)
}

func TestGetInventorySummary_BadJSON(t *testing.T) {
	client := testutil.NewMockedClient(t, func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`this is NOT json`)),
			Header:     make(http.Header),
		}
	})

	api := endpoint.NewInventoryAPI(client)
	_, err := api.GetInventorySummary()

	assert.Error(t, err) // JSON デコード失敗を期待
}
