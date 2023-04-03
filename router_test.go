package router

import (
	"net/http"
	"testing"
)

func TestRouter_AddURL(t *testing.T) {
	mux := &http.ServeMux{}

	r := NewRouter(mux)

	r.AddURL("/foo", func(w http.ResponseWriter, r *Request) {}, TypeGet)
	r.AddURL("/bar", func(w http.ResponseWriter, r *Request) {}, TypeGet)
	r.AddURL("/foo/:id", func(w http.ResponseWriter, r *Request) {}, TypeGet)

	url := r.prepareURL("/foo")
	key := r.getKey(url, TypeGet)

	rList, ok := r.routes[key]

	if !ok {
		t.Error("route list not found")
	}

	rt, err := rList.Find(url)

	if err != nil {
		t.Error("route not found")
	} else if rt.cases[0] != "foo" {
		t.Error("route is not correct")
	}
}

func TestRouter_AddURL1(t *testing.T) {
	mux := &http.ServeMux{}

	r := NewRouter(mux)

	r.AddURL("/foo", func(w http.ResponseWriter, r *Request) {}, TypeGet)
	r.AddURL("/bar", func(w http.ResponseWriter, r *Request) {}, TypeGet)
	r.AddURL("/foo/:id", func(w http.ResponseWriter, r *Request) {}, TypeGet)

	url := r.prepareURL("/bar/")
	key := r.getKey(url, TypeGet)

	rList, ok := r.routes[key]

	if !ok {
		t.Error("route list not found")
	}

	rt, err := rList.Find(url)

	if err != nil {
		t.Error("route not found")
	} else if rt.cases[0] != "bar" {
		t.Error("route is not correct")
	}
}

func TestNewRouter(t *testing.T) {
	mux := &http.ServeMux{}

	r := NewRouter(mux)

	if r == nil {
		t.Error("router not created")
	}
}
