package gaa

import (
	"net/http"
	"net/url"
)

type Context struct {
	res http.ResponseWriter
	req *http.Request
}

func (c *Context) Form() url.Values {
	c.req.ParseForm()
	return c.req.Form
}

func (c *Context) URL() *url.URL {
	return c.req.URL
}

func (c *Context) Method() string {
	return c.req.Method
}
