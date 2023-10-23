package web

import (
	"bytes"
	"net/http"
)

type ConcatString struct {
	buffer bytes.Buffer
}

func (c *ConcatString) Add(String string) *ConcatString {
	c.buffer.WriteString(String)
	return c
}
func (c *ConcatString) Build() string {
	return c.buffer.String()
}

type ResponseModel struct {
	DataContent    ResponseCacheModel
	Error          []error
	HttpStatusCode int
	HttpResponse   *http.Response
}

type ResponseCacheModel struct {
	Data   string
	Header map[string][]string
}
