package endpoint

import (
	"io"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/chiyonn/spapi/client"
	"github.com/google/go-querystring/query"
)

type reqEnc[Req any] func(r *http.Request, p *Req) error

type Endpoint[Req any, Res any] struct {
	c      *client.Client
	Method string
	Path   string
	Rate   float64
	Burst  int
	Key    string

	encode reqEnc[Req] // nil ならデフォルト
	decode func(*http.Response, *Res) error
}

func NewJSONGet[Req any, Res any](
	cl *client.Client,
	path string,
	key string,
	rate float64,
	burst int,
) (*Endpoint[Req, Res], error) {
	if cl == nil {
		return nil, ErrNoClient
	}

	ep := &Endpoint[Req, Res]{
		c:      cl,
		Method: http.MethodGet,
		Path:   path,
		Rate:   rate,
		Burst:  burst,
		Key:    key,
		decode: func(resp *http.Response, v *Res) error {
			return json.NewDecoder(resp.Body).Decode(v)
		},
		encode: func(r *http.Request, p *Req) error {
			if p == nil || isEmptyStruct(*p) {
				return nil
			}
			v, err := query.Values(p)
			if err != nil {
				return err
			}
			r.URL.RawQuery = v.Encode()
			return nil
		},
	}
	if err := cl.RateLimitManager.Register(key, rate, burst); err != nil {
		return nil, err
	}

	return ep, nil
}

func NewJSONPatch[Req any, Res any](
	cl *client.Client,
	path, key string,
	rate float64, burst int,
) (*Endpoint[Req, Res], error) {
	if cl == nil {
		return nil, ErrNoClient
	}

	ep := &Endpoint[Req, Res]{
		c:      cl,
		Method: http.MethodPatch,
		Path:   path,
		Rate:   rate,
		Burst:  burst,
		Key:    key,
		decode: func(resp *http.Response, v *Res) error {
			return json.NewDecoder(resp.Body).Decode(v)
		},
	}

	ep.encode = func(r *http.Request, p *Req) error {
		if p == nil {
			return nil
		}

		vals, err := query.Values(p)
		if err != nil {
			return err
		}
		if len(vals) != 0 {
			r.URL.RawQuery = vals.Encode()
		}

		rv := reflect.ValueOf(p).Elem()
		if f := rv.FieldByName("Body"); f.IsValid() && !f.IsNil() {
			buf := new(bytes.Buffer)
			if err := json.NewEncoder(buf).Encode(f.Interface()); err != nil {
				return err
			}
			r.Body = io.NopCloser(buf)
			r.ContentLength = int64(buf.Len())
			r.Header.Set("Content-Type", "application/json")
		}
		return nil
	}

	if err := cl.RateLimitManager.Register(key, rate, burst); err != nil {
		return nil, err
	}
	return ep, nil
}

func (ep *Endpoint[Req, Res]) Do(
	ctx context.Context, params *Req,
) (*Res, error) {
	if ep.c == nil || ep.c.HTTPClient == nil {
		return nil, errors.New("client nil")
	}
	if err := ep.c.RateLimitManager.Wait(ctx, ep.Key); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(ep.Method, ep.c.BaseURL+ep.Path, nil)
	if err != nil {
		return nil, err
	}
	if ep.encode != nil {
		if err := ep.encode(req, params); err != nil {
			return nil, err
		}
	}

	token, err := ep.c.Auth.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-amz-access-token", token)

	resp, err := ep.c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out Res
	if ep.decode != nil {
		if err := ep.decode(resp, &out); err != nil {
			return nil, err
		}
	}
	return &out, nil
}

func isEmptyStruct(v any) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Struct && rv.NumField() == 0
}

var ErrNoClient = errors.New("client must be set")
