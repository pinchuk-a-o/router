package router

import (
	"errors"
	"net/http"
	"strings"
)

type routeList struct {
	route *Route
}

func (r *routeList) Add(f func(w http.ResponseWriter, r *Request), url string) *Route {
	cases := r.urlToCases(url)

	rt := &Route{handler: f, cases: cases}

	if r.route == nil {
		r.route = rt
	} else {
		rt.next = r.route

		r.route = rt
	}

	return rt
}

func (r *routeList) Find(url string) (fRule *Route, err error) {
	currentURLCases := r.urlToCases(url)

	currentRule := r.route
	for {
		if currentRule.compareRule(currentURLCases) {
			fRule = currentRule
			return
		}

		if currentRule.next == nil {
			err = errors.New("not found")
			return
		}

		currentRule = currentRule.next
	}
}

func (r *routeList) urlToCases(url string) (cases []string) {
	cases = strings.Split(url, "/")

	return
}

//-------------------------------------------------------------------

// Route ...
type Route struct {
	cases     []string
	handler   func(w http.ResponseWriter, r *Request)
	next      *Route
	variables map[string]string
	rules     Rules
}

// SetRules ...
func (r *Route) SetRules(rules map[string]Rule) *Route {
	r.rules.setRules(rules)

	return r
}

func (r *Route) compareRule(currentURLCases []string) (ok bool) {
	var tempVariables = map[string]string{}
	var casesLen = len(r.cases)

	for i := 0; i < casesLen; i++ {
		tempVariables = map[string]string{}
		urlCase := r.cases[i]

		if r.isVariable(urlCase) && r.rules.compare(urlCase, currentURLCases[i]) {
			tempVariables[urlCase[1:]] = currentURLCases[i]
		} else if urlCase != currentURLCases[i] {
			return
		}
	}

	ok = true
	r.variables = tempVariables

	return
}

func (r *Route) isVariable(urlCase string) (ok bool) {
	if len(urlCase) == 0 {
		return
	}

	if urlCase[0:1] == ":" {
		ok = true
	}

	return
}
