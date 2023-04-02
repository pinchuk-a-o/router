```golang 
package main

import (
	"net/http"

	"github.com/pinchuk-a-o/router"
)

func main() {
	mux := &http.ServeMux{}

	r := router.NewRouter(mux)

	r.Set404Handler(func(w http.ResponseWriter, r *http.Request) {
		// 404 handler
	})

	r.AddURL("/foo", func(w http.ResponseWriter, r *router.Request) { w.Write([]byte("its work!!!")) }, router.TypeGet)
	r.AddURL("/foo/:id", func(w http.ResponseWriter, r *router.Request) {
		id, _ := r.GetVariable("id")
		w.Write([]byte(id))
	}, router.TypeGet)

	http.ListenAndServe(":8080", mux)
}

```