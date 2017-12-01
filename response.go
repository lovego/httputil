package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
)

type Response struct {
	*http.Response
	body []byte
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
