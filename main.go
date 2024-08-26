package main

import (
	"net/http"
    "nihility/logger"
)

var msg = logger.Logger {
    Color: logger.Colors.Green,
    Pretext: "main",
}

type Hello struct{}


func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var path = r.URL.Path
    msg.Printf("Serving request: %#v\n", path)

    Router(path, w, r)
}

func main() {
    var h Hello
    err := http.ListenAndServe("localhost:8080", h)
    msg.Println(err)
}
