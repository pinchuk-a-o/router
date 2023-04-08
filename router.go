package router

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Type const`s
const (
	TypePost    = "POST"
	TypeGet     = "GET"
	TypePut     = "PUT"
	TypeHead    = "HEAD"
	TypeDelete  = "DELETE"
	TypeOptions = "OPTIONS"
	TypePatch   = "PATCH"
)

// Router ...
type Router struct {
	mux        *http.ServeMux
	routes     map[string]routeList
	handler404 func(w http.ResponseWriter, r *http.Request)
}

// AddURL ...
func (i *Router) AddURL(method string, url string, f func(w http.ResponseWriter, r *Request)) *Route {
	if !i.checkMethod(method) {
		panic("method not supported")
	}

	url = i.prepareURL(url)
	key := i.getKey(url, method)

	list := routeList{}
	if _, ok := i.routes[key]; ok {
		list = i.routes[key]
	}

	rt := list.Add(f, url)

	i.routes[key] = list

	return rt
}

func (i *Router) getKey(url, method string) string {
	return method + strconv.Itoa(len(strings.Split(url, "/")))
}

func (i *Router) prepareURL(url string) (result string) {
	result = url

	if result == "/" {
		result = ""
		return
	}

	result = url[1:]
	if result[len(result)-1:] == "/" {
		result = result[0 : len(result)-1]
	}

	return
}

// Set404Handler ...
func (i *Router) Set404Handler(f func(w http.ResponseWriter, r *http.Request)) {
	i.handler404 = f
}

func (i *Router) page404(w http.ResponseWriter, r *http.Request) {
	i.handler404(w, r)
}

func default404(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("404 not found"))
	if err != nil {
		return
	}
}

func (i *Router) handler(w http.ResponseWriter, r *http.Request) {
	var (
		rList    routeList
		ok       bool
		url, key string
		route    *Route
		err      error
	)

	url = i.prepareURL(r.RequestURI)
	key = i.getKey(url, r.Method)

	if rList, ok = i.routes[key]; !ok {
		i.page404(w, r)
		return
	}

	route, err = rList.Find(url)

	if err != nil {
		i.page404(w, r)
		return
	}

	request := &Request{httpRequest: r, variables: route.variables}
	route.handler(w, request)
}

func (i *Router) checkMethod(method string) bool {
	switch method {
	case TypeGet, TypeHead, TypeOptions, TypeDelete, TypePatch, TypePost, TypePut:
		return true
	}

	return false
}

// NewRouter ...
func NewRouter(mux *http.ServeMux) (r *Router) {
	o := sync.Once{}

	o.Do(func() {
		r = &Router{
			mux:        mux,
			handler404: default404,
			routes:     make(map[string]routeList),
		}

		r.mux.HandleFunc("/", r.handler)
	})

	return
}
