package httputil

import (
	"context"
	"net/http"
)

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
