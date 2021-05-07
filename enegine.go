package gaa

import (
	"log"
	"net/http"
	"net/url"
)

type Engine struct {
	router *Router
}

func (e Engine) New() *Engine {
	router := RouterNew(Root)
	return &Engine{
		router: router,
	}
}

func (e *Engine) Get(path string, handle Handle) {
	if handle == nil {
		panic("http: nil handler")
	}
	e.router.Handle("GET", path, handle)
}

func (e *Engine) Post(path string, handle Handle) {
	if handle == nil {
		panic("http: nil handler")
	}
	e.router.Handle("POST", path, handle)
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e.router)
}

func Root(res http.ResponseWriter, req *http.Request, url url.Values) {
	log.Printf("Hello Gaa!\n")
}
