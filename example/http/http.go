package main

import (
	"github.com/extism/go-pdk/http"
)

//export http_get
func http_get() int32 {
	// create an HTTP Request (withuot relying on WASI), set headers as needed
	req := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos/1")
	req.SetHeader("some-name", "some-value")
	req.SetHeader("another", "again")
	// send the request, get response back (can check status on response via res.Status())
	res := req.Send()

	// zero-copy output to host
	res.Memory().Output()

	return 0
}

//export http_get_client
func http_get_client() int32 {
	// use the DefaultClient from the go-pdk's http pkg
	c := http.DefaultClient

	// create some headers to be set
	headers := map[string]string{
		"some-name": "some-value",
		"another":   "again",
	}

	// use the client to perform a GET request to the specified URL and with
	// the provided headers set on the reques (note: DefaultClient can perform
	// GET, PUT, POST, and DELETE requests via predefined methods)
	res := c.Get("https://jsonplaceholder.typicode.com/todos/1", headers)

	// zero-copy output to host
	res.Memory().Output()

	return 0
}

func main() {}
