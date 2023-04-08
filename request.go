package router

import (
	"net/http"
)

// Request ...
type Request struct {
	httpRequest *http.Request
	variables   map[string]string
}

// GetVariable ...
func (r Request) GetVariable(name string) (variable string, ok bool) {
	variable, ok = r.variables[name]
	return
}
