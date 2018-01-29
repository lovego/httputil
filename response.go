package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
