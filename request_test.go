package router

import (
	"testing"
)

func TestRequest_GetVariable(t *testing.T) {
	r := Request{variables: map[string]string{"id": "1"}}

	if v, ok := r.GetVariable("id"); !ok || v != "1" {
		t.Error("TestRequest_GetVariable ERROR")
	} else {
		t.Log("TestRequest_GetVariable DONE")
	}
}
