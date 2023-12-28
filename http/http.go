package http

import (
	"encoding/json"
	"io"
	h "net/http"
	"strings"

	"github.com/extism/go-pdk/internal"
	"github.com/extism/go-pdk/memory"
)

type HTTPRequestMeta struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}

type HTTPRequest struct {
	meta HTTPRequestMeta
	body []byte
}

type HTTPResponse struct {
	memory *memory.Memory
	status uint16
}

func (r HTTPResponse) Memory() *memory.Memory {
	return r.memory
}

func (r HTTPResponse) Body() []byte {
	buf := make([]byte, r.memory.Len)
	r.memory.Load(buf)
	return buf
}

func (r HTTPResponse) Status() uint16 {
	return r.status
}

func NewRequest(method string, url string) *HTTPRequest {
	return &HTTPRequest{
		meta: HTTPRequestMeta{
			Url:    url,
			Method: strings.ToUpper(method),
		},
		body: nil}
}

func (r *HTTPRequest) SetHeader(key string, value string) *HTTPRequest {
	if r.meta.Headers == nil {
		r.meta.Headers = make(map[string]string)
	}
	r.meta.Headers[key] = value
	return r
}

func read(r io.Reader) []byte {
	buf := make([]byte, 8)

	slice := make([]byte, 0)

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			// there is no more data to read
			break
		}

		if n > 0 {
			for i, b := range buf {
				if n == i {
					break
				}

				slice = append(slice, b)
			}
		}
	}

	return slice
}

func (r *HTTPRequest) SetBody(body io.Reader) *HTTPRequest {
	r.body = read(body)
	return r
}

func (r *HTTPRequest) Send() HTTPResponse {
	enc, _ := json.Marshal(r.meta)

	offset, clength := internal.AllocBytes(enc)

	req := &memory.Memory{
		Off: offset,
		Len: clength,
	}
	defer req.Free()

	bOff, cCLen := internal.AllocBytes(r.body)
	data := &memory.Memory{
		Off: bOff,
		Len: cCLen,
	}
	defer data.Free()

	off := req.Off.HTTPRequest(data.Off)
	length := off.Length()
	status := uint16(internal.HTTPStatusCode())

	mem := &memory.Memory{
		Off: off,
		Len: length,
	}

	return HTTPResponse{
		mem,
		status,
	}
}

type client struct {
}

var DefaultClient = &client{}

func (c *client) Do(req *HTTPRequest, headers map[string]string) HTTPResponse {
	if headers != nil || len(headers) != 0 {
		for k, v := range headers {
			req.SetHeader(k, v)
		}
	}

	return req.Send()
}

func (c *client) Get(url string, headers map[string]string) HTTPResponse {
	req := NewRequest(h.MethodGet, url)

	if headers != nil || len(headers) != 0 {
		for k, v := range headers {
			req.SetHeader(k, v)
		}
	}

	return req.Send()
}

func (c *client) Post(url string, headers map[string]string, contentType string, body io.Reader) HTTPResponse {
	req := NewRequest(h.MethodPost, url)

	if headers != nil || len(headers) != 0 {
		for k, v := range headers {
			req.SetHeader(k, v)
		}
	}

	req.SetHeader("Content-Type", contentType)

	req.SetBody(body)

	return req.Send()
}

func (c *client) Put(url string, headers map[string]string, contentType string, body io.Reader) HTTPResponse {
	req := NewRequest(h.MethodPut, url)

	if headers != nil || len(headers) != 0 {
		for k, v := range headers {
			req.SetHeader(k, v)
		}
	}

	req.SetHeader("Content-Type", contentType)

	req.SetBody(body)

	return req.Send()
}

func (c *client) Delete(url string, headers map[string]string) HTTPResponse {
	req := NewRequest(h.MethodDelete, url)

	if headers != nil || len(headers) != 0 {
		for k, v := range headers {
			req.SetHeader(k, v)
		}
	}

	return req.Send()
}
