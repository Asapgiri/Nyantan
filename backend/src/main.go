package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

var pages = map[string]func(w http.ResponseWriter, r *http.Request) {
    "/":            get_root,
    "/index":       get_root,

    // Authentications
    "/login":       get_login,
    "/logout":       get_login,

    // Translation related
    "/translate":   get_translate,
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var path = r.URL.Path
    fmt.Printf("Serving r: %#v\n", path)

    page := pages[path]
    if nil == page {
        unexpected(w, r)
    } else {
        page(w, r)
    }
}

func main() {
    var h Hello
    http.ListenAndServe("localhost:4000", h)
}
