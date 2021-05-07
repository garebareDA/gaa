package gaa

import (
	"net/http"
	"net/url"
	"strings"
)

type Handle func(http.ResponseWriter, *http.Request, url.Values)

type Router struct {
	tree        *node
	rootHandler Handle
}

func RouterNew(rootHandler Handle) *Router {
	node := node{
		component:    "/",
		isNamedParam: false,
		methods:      make(map[string]Handle),
	}

	return &Router{tree: &node, rootHandler: rootHandler}
}

func (r *Router) Handle(method string, path string, handler Handle) {
	if path[0] != '/' {
		panic("Path has to start with a /.")
	}

	r.tree.addNode(method, path, handler)
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	if handler := node.methods[req.Method]; handler != nil {
		handler(res, req, params)
	} else {
		r.rootHandler(res, req, params)
	}
}
