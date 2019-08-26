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

func (c *Client) Do(method, url string, headers map[string]string, body interface{}) (*Response, error) {
	req, err := c.makeReq(method, url, headers, body)
	if err != nil {
		return nil, err
	}
	return c.DoReq(req)
}

func (c *Client) DoCtx(
	ctx context.Context, opName, method, url string, headers map[string]string, body interface{},
) (*Response, error) {
	req, err := c.makeReq(method, url, headers, body)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		ctx = tracer.Start(ctx, opName)
		defer tracer.Finish(ctx)
		req = req.WithContext(ctx)
	}
	return c.DoReq(req)
}

func (c *Client) DoReq(req *http.Request) (*Response, error) {
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
	resp, err := c.DoCtx(ctx, opName, method, url, headers, body)
	if err != nil {
		return err
	}
	if err := resp.Ok(); err != nil {
		return err
	}
	return resp.Json(data)
}

func (c *Client) makeReq(
	method, url string, headers map[string]string, body interface{},
) (*http.Request, error) {
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
	return req, nil
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
