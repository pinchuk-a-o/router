package router

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
)

const (
	TypePost    = "POST"
	TypeGet     = "GET"
	TypePut     = "PUT"
	TypeHead    = "HEAD"
	TypeDelete  = "DELETE"
	TypeOptions = "OPTIONS"
	TypePatch   = "PATCH"
)

type Router struct {
	mux        *http.ServeMux
	routes     map[string]routeList
	handler404 func(w http.ResponseWriter, r *http.Request)
}

func (i *Router) AddURL(url string, f func(w http.ResponseWriter, r *Request), method string) {
	url = i.prepareURL(url)
	key := i.getKey(url, method)

	list := routeList{}
	if _, ok := i.routes[key]; !ok {
		list = i.routes[key]
	}

	list.Add(f, url)

	o := sync.Once{}
	o.Do(func() {
		i.routes = map[string]routeList{}
	})

	i.routes[key] = list
}

func (i Router) getKey(url, method string) string {
	return method + strconv.Itoa(len(strings.Split(url, "/")))
}

func (i Router) prepareURL(url string) (result string) {
	result = url

	if result == "/" {
		result = ""
		return
	}

	result = url[1:]
	if result[len(url)-1:] == "/" {
		result = url[1 : len(url)-1]
	}

	return
}

func (i *Router) Set404Handler(f func(w http.ResponseWriter, r *http.Request)) {
	i.handler404 = f
}

func (i Router) page404(w http.ResponseWriter, r *http.Request) {
	i.handler404(w, r)
}

func default404(w http.ResponseWriter, r *http.Request) {
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
		route    route
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

func NewRouter(mux *http.ServeMux) (r *Router) {
	o := sync.Once{}

	o.Do(func() {
		r = &Router{
			mux:        mux,
			handler404: default404,
		}

		r.mux.HandleFunc("/", r.handler)
	})

	return
}
