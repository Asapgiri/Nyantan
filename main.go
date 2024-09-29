package main

import (
	"net/http"
	"nihility/logger"
	"nihility/pages"
    "nihility/dbase"
)

var msg = logger.Logger {
    Color: logger.Colors.Green,
    Pretext: "main",
}

type Hello struct{}


func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var path = r.URL.Path
    msg.Printf("Serving request: %#v\n", path)

    pages.Unexpected(w, r)
}

func main() {
    dbase.Connect()
    setup_routes()

    err := http.ListenAndServe("localhost:8080", nil)
    msg.Println(err)
}
