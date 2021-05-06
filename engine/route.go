package engine

import "net/url"

type Handle func(*Context, url.Values)

type node struct {
	children     []*node
	component    string
	isNamedParam bool
	methods      map[string]Handle
}

type Router struct {
	tree        *node
	rootHandler Handle
}

func New(rootHandler Handle) *Router {
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
