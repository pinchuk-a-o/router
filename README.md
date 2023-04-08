```golang 
package main

import (
	"net/http"

	. "github.com/pinchuk-a-o/router"
)

func main() {
	mux := &http.ServeMux{}

	r := NewRouter(mux)

	r.Set404Handler(func(w http.ResponseWriter, r *http.Request) {
		// 404 handler
	})

	r.AddURL(
		TypeGet,
		"/foo",
		func(w http.ResponseWriter, r *Request) {
			w.Write([]byte("its work!!!"))
		})
	r.AddURL(
		TypeGet,
		"/foo/:id",
		func(w http.ResponseWriter, r *Request) {
			id, _ := r.GetVariable("id")
			w.Write([]byte(id))
		}).SetRules(map[string]Rule{":id": RuleInteger{}})

	http.ListenAndServe(":8080", mux)
}

```