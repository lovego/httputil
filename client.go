package httputil

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/lovego/tracer"
)

type Client struct {
	BaseUrl string
	Client  *http.Client
}

func (c *Client) Get(url string, headers map[string]string, body interface{}) (*Response, error) {
	return c.Do(http.MethodGet, url, headers, body)
}

func (c *Client) GetCtx(
	ctx context.Context, opName, url string, headers map[string]string, body interface{},
) (*Response, error) {
	return c.DoCtx(ctx, opName, http.MethodGet, url, headers, body)
}

func (c *Client) GetJson(url string, headers map[string]string, body, data interface{}) error {
	return c.DoJson(http.MethodGet, url, headers, body, data)
}

func (c *Client) GetJsonCtx(ctx context.Context, opName, url string, headers map[string]string, body, data interface{}) error {
	return c.DoJsonCtx(ctx, opName, http.MethodGet, url, headers, body, data)
}

func (c *Client) Post(url string, headers map[string]string, body interface{}) (*Response, error) {
	return c.Do(http.MethodPost, url, headers, body)
}

func (c *Client) PostCtx(
	ctx context.Context, opName, url string, headers map[string]string, body interface{},
) (*Response, error) {
	return c.DoCtx(ctx, opName, http.MethodPost, url, headers, body)
}
func (c *Client) PostJson(url string, headers map[string]string, body, data interface{}) error {
	return c.DoJson(http.MethodPost, url, headers, body, data)
}

func (c *Client) PostJsonCtx(ctx context.Context, opName, url string, headers map[string]string, body, data interface{}) error {
	return c.DoJsonCtx(ctx, opName, http.MethodPost, url, headers, body, data)
}

func (c *Client) Head(url string, headers map[string]string, body interface{}) (*Response, error) {
	return c.Do(http.MethodHead, url, headers, body)
}

func (c *Client) HeadCtx(
	ctx context.Context, opName, url string, headers map[string]string, body interface{},
) (*Response, error) {
	return c.DoCtx(ctx, opName, http.MethodHead, url, headers, body)
}
func (c *Client) HeadJson(url string, headers map[string]string, body, data interface{}) error {
	return c.DoJson(http.MethodHead, url, headers, body, data)
}

func (c *Client) HeadJsonCtx(ctx context.Context, opName, url string, headers map[string]string, body, data interface{}) error {
	return c.DoJsonCtx(ctx, opName, http.MethodHead, url, headers, body, data)
}

func (c *Client) Put(url string, headers map[string]string, body interface{}) (*Response, error) {
	return c.Do(http.MethodPut, url, headers, body)
}

func (c *Client) PutCtx(
	ctx context.Context, opName, url string, headers map[string]string, body interface{},
) (*Response, error) {
	return c.DoCtx(ctx, opName, http.MethodPut, url, headers, body)
}

func (c *Client) PutJson(url string, headers map[string]string, body, data interface{}) error {
	return c.DoJson(http.MethodPut, url, headers, body, data)
}

func (c *Client) PutJsonCtx(ctx context.Context, opName, url string, headers map[string]string, body, data interface{}) error {
	return c.DoJsonCtx(ctx, opName, http.MethodPut, url, headers, body, data)
}

func (c *Client) Delete(url string, headers map[string]string, body interface{}) (*Response, error) {
	return c.Do(http.MethodDelete, url, headers, body)
}

func (c *Client) DeleteCtx(
	ctx context.Context, opName, url string, headers map[string]string, body interface{},
) (*Response, error) {
	return c.DoCtx(ctx, opName, http.MethodDelete, url, headers, body)
}

func (c *Client) DeleteJson(url string, headers map[string]string, body, data interface{}) error {
	return c.DoJson(http.MethodDelete, url, headers, body, data)
}

func (c *Client) DeleteJsonCtx(ctx context.Context, opName, url string, headers map[string]string, body, data interface{}) error {
	return c.DoJsonCtx(ctx, opName, http.MethodDelete, url, headers, body, data)
}

func (c *Client) Do(method, url string, headers map[string]string, body interface{}) (*Response, error) {
	bodyReader, err := makeBodyReader(body)
	if err != nil {
		return nil, err
	}
	if c.BaseUrl != `` {
		url = c.BaseUrl + url
	}
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header[k] = []string{v}
	}
	resp, err := c.Client.Do(req)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Response{Response: resp, body: respBody}, nil
}

func (c *Client) DoCtx(
	ctx context.Context, opName, method, url string, headers map[string]string, body interface{},
) (*Response, error) {
	defer tracer.StartSpan(ctx, opName).Finish()
	return c.Do(method, url, headers, body)
}

func (c *Client) DoJson(method, url string, headers map[string]string, body, data interface{}) error {
	resp, err := c.Do(method, url, headers, body)
	if err != nil {
		return err
	}
	if err := resp.Ok(); err != nil {
		return err
	}
	return resp.Json(data)
}

func (c *Client) DoJsonCtx(
	ctx context.Context, opName, method, url string, headers map[string]string, body, data interface{},
) error {
	defer tracer.StartSpan(ctx, opName).Finish()
	return c.DoJson(method, url, headers, body, data)
}

func makeBodyReader(data interface{}) (io.Reader, error) {
	if data == nil {
		return nil, nil
	}
	var reader io.Reader
	switch body := data.(type) {
	case io.Reader:
		reader = body
	case string:
		if len(body) > 0 {
			reader = strings.NewReader(body)
		}
	case []byte:
		if len(body) > 0 {
			reader = bytes.NewBuffer(body)
		}
	default:
		if !isNil(body) {
			buf, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			reader = bytes.NewBuffer(buf)
		}
	}
	return reader, nil
}

func isNil(data interface{}) bool {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
		return v.IsNil()
	default:
		return false
	}
}
