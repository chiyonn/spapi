# Amazon SP-API Go Client

A Go client for the [Amazon Selling Partner API (SP-API)](https://developer-docs.amazon.com/sp-api). This library provides convenient access to supported SP-API endpoints, OAuth2 authentication, request signing, rate limiting, and structured response decoding.

## Features

* 🔐 OAuth2 + AWS Signature Version 4 authentication
* 🚀 Built-in rate limiting
* 📦 Modular endpoint packages (inventory, listings, etc.)
* 🔍 Struct-based request/response handling

---

## Installation

```bash
go get github.com/chiyonn/spapi
```

---

## Usage

### 1. Set environment variables

```bash
export SPAPI_REFRESH_TOKEN=your_refresh_token
export LWA_CLIENT_ID=your_client_id
export LWA_CLIENT_SECRET=your_client_secret
```

### 2. Initialize client and call an endpoint

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/chiyonn/spapi/auth"
	"github.com/chiyonn/spapi/client"
	"github.com/chiyonn/spapi/endpoint/inventory"
)

func main() {
	cfg, _ := auth.NewAuthConfig(
		os.Getenv("SPAPI_REFRESH_TOKEN"),
		os.Getenv("LWA_CLIENT_ID"),
		os.Getenv("LWA_CLIENT_SECRET"),
	)

	cli, _ := client.NewClient(&http.Client{Timeout: 10 * time.Second}, "JP", cfg, client.NewRateLimitManager())
	invAPI := inventory.NewInventoryAPI(cli)

	params := &inventory.GetInventorySummariesParams{
		GranularityType: "Marketplace",
		GranularityId:   "A1VC38T7YXB528", // Japan marketplace
	}

	res, err := invAPI.GetInventorySummaries(params)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", res)
}
```

---

## Available Endpoints

| Method                  | Description                       |
| ----------------------- | --------------------------------- |
| `GetInventorySummaries` | Retrieves FBA inventory summaries |

---

## License

This project is licensed under the MIT License, see the LICENSE.txt file for details
