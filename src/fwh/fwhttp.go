package fwh

import (
	"fmt"
	"net/http"
)

func Start() {
	fmt.Println("start")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello, %q", r.URL.Query()["who"]) })
	http.ListenAndServe(":8080", nil)
}
