package gaa

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestServer(t *testing.T) {
	router := RouterNew(Root)
	m := MiddleNew()
	router.Handle("GET", "/", m.Then(func(w http.ResponseWriter, r *http.Request, u url.Values) {
		fmt.Fprint(w, "hello")
	}))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("req faital")
	}
	router.ServeHTTP(w, req)

	if 200 != w.Code {
		t.Fatalf("status code %d", w.Code)
	}

	if w.Body.String() != "hello" {
		t.Fatalf("body %s faital", w.Body.String())
	}
}
