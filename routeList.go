package router

import (
	"errors"
	"net/http"
	"strings"
)

type routeList struct {
	route *route
}

func (r *routeList) Add(f func(w http.ResponseWriter, r *Request), url string) {
	cases := r.urlToCases(url)

	route := &route{handler: f, cases: cases}

	if r.route == nil {
		r.route = route
	} else {
		route.next = r.route

		r.route = route
	}
}

func (r routeList) Find(url string) (fRule route, err error) {
	currentURLCases := r.urlToCases(url)

	currentRule := r.route
	for {
		if currentRule.compareRule(currentURLCases) {
			fRule = *currentRule
			return
		}

		if currentRule.next == nil {
			err = errors.New("not found")
			return
		}

		currentRule = currentRule.next
	}
}

func (r routeList) urlToCases(url string) (cases []string) {
	cases = strings.Split(url, "/")

	return
}

//-------------------------------------------------------------------

type route struct {
	cases     []string
	handler   func(w http.ResponseWriter, r *Request)
	next      *route
	variables map[string]string
}

func (r *route) compareRule(cases []string) (ok bool) {
	var tempVariables = map[string]string{}
	var casesLen = len(r.cases)

	for i := 0; i < casesLen; i++ {
		tempVariables = map[string]string{}
		urlCase := r.cases[i]

		if r.isVariable(urlCase) {
			tempVariables[urlCase[1:]] = cases[i]
		} else if urlCase != cases[i] {
			return
		}
	}

	ok = true
	r.variables = tempVariables

	return
}

func (r *route) isVariable(urlCase string) (ok bool) {
	if len(urlCase) == 0 {
		return
	}

	if urlCase[0:1] == ":" {
		ok = true
	}

	return
}
