package gaa

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServer(t *testing.T) {
	router := RouterNew(Root)
	m := MiddleNew()
	router.Handle("GET", "/", m.Then(func(w http.ResponseWriter, r *http.Request, u url.Values) {
		fmt.Fprint(w, "hello")
	}))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), "hello")
}

func TestPostServer(t *testing.T) {
	router := RouterNew(Root)
	m := MiddleNew()
	router.Handle("POST", "/ps", m.Then(func(w http.ResponseWriter, r *http.Request, u url.Values) {
		fmt.Fprint(w, r.Body)
	}))

	w := httptest.NewRecorder()

	body := bytes.NewBufferString("{\"msg\":{\"name\":\"foo\"}}")
	req, _ := http.NewRequest("POST", "/ps", body)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Body.String(), "{{\"msg\":{\"name\":\"foo\"}}}")
	assert.Equal(t, w.Code, 200)
}
