package gaa

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

type Engine struct {
	router *Router
	middle Middle
}

func EngineNew() *Engine {
	router := RouterNew(Root)
	middle := MiddleNew(loggingHandler)
	return &Engine{
		router: router,
		middle: middle,
	}
}

func loggingHandler(next Handle) Handle {
	fn := func(w http.ResponseWriter, r *http.Request, u url.Values) {
		t1 := time.Now()
		next(w, r, u)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return fn
}

func (e *Engine) Get(path string, handle Handle) {
	if handle == nil {
		panic("http: nil handler")
	}
	e.router.Handle("GET", path, e.middle.Then(handle))
}

func (e *Engine) Post(path string, handle Handle) {
	if handle == nil {
		panic("http: nil handler")
	}
	e.router.Handle("POST", path, e.middle.Then(handle))
}

func (e *Engine) Run(addr string) {
	log.Println("Liten on" + addr)
	http.ListenAndServe(addr, e.router)
}

func Root(res http.ResponseWriter, req *http.Request, url url.Values) {

}

func (e *Engine) Middle(constructors ...Constructor) {
	add := e.middle.Append(constructors)
	e.middle = add
}
