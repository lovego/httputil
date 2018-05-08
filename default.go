package httputil

import (
	"context"
	"net/http"
	"time"
)

var DefaultClient = &Client{Client: &http.Client{Timeout: 10 * time.Second}}

func GetJson(url string, headers map[string]string, body, data interface{}) error {
	return DefaultClient.DoJson(http.MethodGet, url, headers, body, data)
}

func PostJson(url string, headers map[string]string, body, data interface{}) error {
	return DefaultClient.DoJson(http.MethodPost, url, headers, body, data)
}

func HeadJson(url string, headers map[string]string, body, data interface{}) error {
	return DefaultClient.DoJson(http.MethodHead, url, headers, body, data)
}

func PutJson(url string, headers map[string]string, body, data interface{}) error {
	return DefaultClient.DoJson(http.MethodPut, url, headers, body, data)
}

func DeleteJson(url string, headers map[string]string, body, data interface{}) error {
	return DefaultClient.DoJson(http.MethodDelete, url, headers, body, data)
}

func GetJsonCtx(ctx context.Context, opName, url string, headers map[string]string, body, data interface{}) error {
	return DefaultClient.DoJsonCtx(ctx, opName, http.MethodGet, url, headers, body, data)
}

func PostJsonCtx(ctx context.Context, opName, url string, headers map[string]string, body, data interface{}) error {
	return DefaultClient.DoJsonCtx(ctx, opName, http.MethodPost, url, headers, body, data)
}

func HeadJsonCtx(ctx context.Context, opName, url string, headers map[string]string, body, data interface{}) error {
	return DefaultClient.DoJsonCtx(ctx, opName, http.MethodHead, url, headers, body, data)
}

func PutJsonCtx(ctx context.Context, opName, url string, headers map[string]string, body, data interface{}) error {
	return DefaultClient.DoJsonCtx(ctx, opName, http.MethodPut, url, headers, body, data)
}

func DeleteJsonCtx(ctx context.Context, opName, url string, headers map[string]string, body, data interface{}) error {
	return DefaultClient.DoJsonCtx(ctx, opName, http.MethodDelete, url, headers, body, data)
}
