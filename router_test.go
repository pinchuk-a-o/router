package router

import (
	"net/http"
	"testing"
)

func TestRouter_AddURL(t *testing.T) {
	mux := &http.ServeMux{}

	r := NewRouter(mux)

	r.AddURL(TypeGet, "/foo", func(w http.ResponseWriter, r *Request) {})
	r.AddURL(TypeGet, "/bar", func(w http.ResponseWriter, r *Request) {})
	r.AddURL(TypeGet, "/foo/:id", func(w http.ResponseWriter, r *Request) {})
	r.AddURL(TypeGet, "/foo/:name", func(w http.ResponseWriter, r *Request) {})

	url := r.prepareURL("/foo")
	key := r.getKey(url, TypeGet)

	rList, ok := r.routes[key]

	if !ok {
		t.Error("Route list not found")
	}

	rt, err := rList.Find(url)

	if err != nil {
		t.Error("Route not found")
	} else if rt.cases[0] != "foo" {
		t.Error("Route is not correct")
	}
}

func TestRouter_AddURL1(t *testing.T) {
	mux := &http.ServeMux{}

	r := NewRouter(mux)

	r.AddURL(TypeGet, "/foo", func(w http.ResponseWriter, r *Request) {})
	r.AddURL(TypeGet, "/bar", func(w http.ResponseWriter, r *Request) {})
	r.AddURL(TypeGet, "/foo/:id", func(w http.ResponseWriter, r *Request) {})

	url := r.prepareURL("/bar/")
	key := r.getKey(url, TypeGet)

	rList, ok := r.routes[key]

	if !ok {
		t.Error("Route list not found")
	}

	rt, err := rList.Find(url)

	if err != nil {
		t.Error("Route not found")
	} else if rt.cases[0] != "bar" {
		t.Error("Route is not correct")
	}
}

func TestRouter_AddURL2(t *testing.T) {
	mux := &http.ServeMux{}

	r := NewRouter(mux)

	r.AddURL(TypeGet, "/foo/:id", func(w http.ResponseWriter, r *Request) {}).SetRules(map[string]Rule{":id": RuleInteger{}})
	r.AddURL(TypeGet, "/foo/:name", func(w http.ResponseWriter, r *Request) {}).SetRules(map[string]Rule{":name": RuleLetter{}})

	url1 := r.prepareURL("/foo/12")
	url2 := r.prepareURL("/foo/asd")
	url3 := r.prepareURL("/foo/asd-1")
	key1 := r.getKey(url1, TypeGet)
	key2 := r.getKey(url2, TypeGet)
	key3 := r.getKey(url3, TypeGet)

	rList1, ok1 := r.routes[key1]

	if !ok1 {
		t.Error("Route list not found")
	}

	rt, err := rList1.Find(url1)

	if err != nil {
		t.Error("Route not found")
	} else if rt.cases[1] != ":id" {
		t.Error("Route is not correct")
	}
	//------

	rList2, ok2 := r.routes[key2]
	if !ok2 {
		t.Error("Route list not found")
	}

	rt2, err := rList2.Find(url2)

	if err != nil {
		t.Error("Route not found")
	} else if rt2.cases[1] != ":name" {
		t.Error("Route is not correct")
	}

	//------

	rList3, ok3 := r.routes[key3]
	if !ok3 {
		t.Error("Route list not found")
	}

	_, err = rList3.Find(url3)

	if err == nil {
		t.Error("Route found " + url3)
	}
}

func TestRouter_AddURL3(t *testing.T) {
	mux := &http.ServeMux{}

	r := NewRouter(mux)

	r.AddURL(TypeGet, "/", func(w http.ResponseWriter, r *Request) {})

	url1 := r.prepareURL("/")
	key1 := r.getKey(url1, TypeGet)

	rList1, ok1 := r.routes[key1]

	if !ok1 {
		t.Error("Route list not found")
	}

	_, err := rList1.Find(url1)

	if err != nil {
		t.Error("Route not found")
	}
}

func TestRouter_AddURL4(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("No error for not supported method")
		}
	}()
	mux := &http.ServeMux{}

	r := NewRouter(mux)

	r.AddURL("get", "/", func(w http.ResponseWriter, r *Request) {})
}

func TestNewRouter(t *testing.T) {
	mux := &http.ServeMux{}

	r := NewRouter(mux)

	if r == nil {
		t.Error("router not created")
	}
}
