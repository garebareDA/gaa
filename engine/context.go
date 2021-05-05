package engine

import (
	"net/http"
)

type Context struct {
	res http.ResponseWriter
	req *http.Request
}
