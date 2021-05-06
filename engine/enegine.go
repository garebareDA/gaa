package engine

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/justinas/alice"
)

type Engine struct {
	handler alice.Chain
}

type ContextFunc func(*Context)

func loggingHandler(next http.Handler) http.Handler {
	log := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		times := time.Now()
		log.Printf("[%s] %q %v \n", r.Method, r.URL.String(), times)
	}
	return http.HandlerFunc(log)
}

func recoverHandler(next http.Handler) http.Handler {
	rec := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(rec)
}

func (e Engine) New() *Engine {
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)
	return &Engine{
		handler: commonHandlers,
	}
}

func (e Engine) Get(path string, header ContextFunc) {
	if header == nil {
		panic("http: nil header")
	}

}

func (e Engine) Post(path string, c ContextFunc) {

}
