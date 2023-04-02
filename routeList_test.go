package router

import (
	"net/http"
	"testing"
)

func TestRouteList_Find(t *testing.T) {
	r := &routeList{}

	url1t := "foo1/:name"
	url1 := "foo1/golang"
	url2t := "foo/:id"
	url2 := "foo/1"
	url3t := "bar/:id"
	url3 := "bar/2"

	r.Add(func(w http.ResponseWriter, r *Request) {}, url1t)
	r.Add(func(w http.ResponseWriter, r *Request) {}, url2t)
	r.Add(func(w http.ResponseWriter, r *Request) {}, url3t)

	r1, err1 := r.Find(url1)
	if err1 != nil {
		t.Error(url1 + " not found")
	}

	if r1.variables["name"] != "golang" {
		t.Error(url1 + " variables not found")
	}

	r2, err2 := r.Find(url2)
	if err2 != nil {
		t.Error(url2 + " not found")
	}

	if r2.variables["id"] != "1" {
		t.Error(url2 + " variables not found")
	}

	r3, err3 := r.Find(url3)
	if err3 != nil {
		t.Error(url3 + " not found")
	}

	if r3.variables["id"] != "2" {
		t.Error(url3 + " variables not found")
	}
}
