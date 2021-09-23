package httputil

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lovego/errs"
)

type Response struct {
	*http.Response
	body          []byte
	UnmarshalFunc func(data []byte, v interface{}) error
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
	return setErrorData(errs.Newf("", "Unexpected Response: %s", resp.Status), resp)
}

func (resp *Response) Json(data interface{}) error {
	if data == nil {
		return nil
	}

	if err := resp.GetUnmarshalFunc()(resp.body, data); err != nil {
		return setErrorData(errs.Wrap(err).(*errs.Error), resp)
	}
	if d, ok := data.(interface {
		ValidateResponse(resp *Response) error
	}); ok {
		return d.ValidateResponse(resp)
	}
	return nil
}

func (resp *Response) GetUnmarshalFunc() func(data []byte, v interface{}) error {
	if resp.UnmarshalFunc != nil {
		return resp.UnmarshalFunc
	}
	return json.Unmarshal
}

type CodeMessageData struct {
	Code, Message string
	Data          interface{}
}

func (cmd *CodeMessageData) ValidateResponse(resp *Response) error {
	switch cmd.Code {
	case "ok":
		return nil
	case "":
		return setErrorData(errs.New("", "Unexpected Response Body"), resp)
	default:
		return errs.New(cmd.Code, cmd.Message)
	}
}

func setErrorData(err *errs.Error, resp *Response) *errs.Error {
	return err.SetData(fmt.Sprintf("HTTP %s %s\nResponse Body: %s",
		resp.Request.Method, resp.Request.URL.String(), resp.body,
	))
}
