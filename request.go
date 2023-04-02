package router

import (
	"net/http"
)

type Request struct {
	httpRequest *http.Request
	variables   map[string]string
}

func (r Request) GetVariable(name string) (variable string, ok bool) {
	variable, ok = r.variables[name]
	return
}
