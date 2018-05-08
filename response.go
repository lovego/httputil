package httputil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	*http.Response
	body []byte
}

func Get(url string, headers map[string]string, body interface{}) (*Response, error) {
	return DefaultClient.Do(http.MethodGet, url, headers, body)
}

func GetCtx(
	ctx context.Context, opName, url string, headers map[string]string, body interface{},
) (*Response, error) {
	return DefaultClient.DoCtx(ctx, opName, http.MethodGet, url, headers, body)
}

func Post(url string, headers map[string]string, body interface{}) (*Response, error) {
	return DefaultClient.Do(http.MethodPost, url, headers, body)
}

func PostCtx(
	ctx context.Context, opName, url string, headers map[string]string, body interface{},
) (*Response, error) {
	return DefaultClient.DoCtx(ctx, opName, http.MethodPost, url, headers, body)
}

func Head(url string, headers map[string]string, body interface{}) (*Response, error) {
	return DefaultClient.Do(http.MethodHead, url, headers, body)
}

func HeadCtx(
	ctx context.Context, opName, url string, headers map[string]string, body interface{},
) (*Response, error) {
	return DefaultClient.DoCtx(ctx, opName, http.MethodHead, url, headers, body)
}

func Put(url string, headers map[string]string, body interface{}) (*Response, error) {
	return DefaultClient.Do(http.MethodPut, url, headers, body)
}

func PutCtx(
	ctx context.Context, opName, url string, headers map[string]string, body interface{},
) (*Response, error) {
	return DefaultClient.DoCtx(ctx, opName, http.MethodPut, url, headers, body)
}

func Delete(url string, headers map[string]string, body interface{}) (*Response, error) {
	return DefaultClient.Do(http.MethodDelete, url, headers, body)
}

func DeleteCtx(
	ctx context.Context, opName, url string, headers map[string]string, body interface{},
) (*Response, error) {
	return DefaultClient.DoCtx(ctx, opName, http.MethodDelete, url, headers, body)
}

func (resp *Response) Body() []byte {
	return resp.body
}

func (resp *Response) Ok() error {
	if resp.StatusCode != http.StatusOK {
		return resp.CodeError()
	}
	return nil
}

func (resp *Response) Check(codes ...int) error {
	for _, code := range codes {
		if resp.StatusCode == code {
			return nil
		}
	}
	return resp.CodeError()
}

func (resp *Response) CodeError() error {
	return fmt.Errorf(`HTTP %s %s
Unexpected Response: %s
%s`, resp.Request.Method, resp.Request.URL.String(), resp.Status, resp.body,
	)
}

func (resp *Response) Json(data interface{}) error {
	if data == nil {
		return nil
	}
	decoder := json.NewDecoder(bytes.NewBuffer(resp.body))
	decoder.UseNumber()
	return decoder.Decode(&data)
}

func (resp *Response) Json2(data interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(resp.body, &data); err != nil {
		return err
	}
	return nil
}
