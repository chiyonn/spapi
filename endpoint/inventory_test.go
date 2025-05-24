// endpoint/inventory_test.go
package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/chiyonn/spapi/client"
	"github.com/chiyonn/spapi/endpoint"
	"github.com/chiyonn/spapi/testutil"
	"github.com/stretchr/testify/assert"
)

/* ------------------  テスト補助  ------------------ */

// noopAuthenticator は auth を完全に無効化するダミー実装。
// プロダクションの Authenticator インターフェースに合わせて置き換えてください。
type noopAuthenticator struct{}

func (n *noopAuthenticator) GetAccessToken(ctx context.Context) (string, error) {
	return "dummy-access-token", nil
}

// newMockedClient は InventoryAPI テスト用の最小 Client を返す。
// ・HTTP ラウンドトリップは呼び出し元が渡す handler でモック
// ・RateLimitManager も実態依存を持たない形で生成
// ・Auth には noopAuthenticator を注入し、認可ロジックを完全に除外
func newMockedClient(t *testing.T, handler testutil.RoundTripFunc) *client.Client {
	t.Helper()

	return &client.Client{
		BaseURL:          "https://mock.api",
		HTTPClient:       testutil.NewMockHTTPClient(handler),
		RateLimitManager: client.NewRateLimitManager(),
		Auth:             &noopAuthenticator{},
	}
}

/* ------------------  テストケース  ------------------ */

func TestGetInventorySummary_Success(t *testing.T) {
	client := newMockedClient(t, func(req *http.Request) *http.Response {
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
	client := newMockedClient(t, func(req *http.Request) *http.Response {
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
