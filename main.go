package main

import (
	"net/http"
	"nihility/apis"
	"nihility/logger"
	"nihility/pages"
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
    http.HandleFunc("GET /",            pages.Root)
    http.HandleFunc("GET /index",       pages.Root)
    http.HandleFunc("GET /index.html",  pages.Root)

    // Authentications
    http.HandleFunc("GET  /login",      pages.Login)
    http.HandleFunc("POST /login",      pages.Login)
    http.HandleFunc("GET  /register",   pages.Register)
    http.HandleFunc("POST /register",   pages.Register)
    http.HandleFunc("GET  /logout",     pages.Logout)

    // Translation related
    http.HandleFunc("GET /translate",   pages.Translate)

    http.HandleFunc("GET /trans/{id}",          pages.Translation)
    http.HandleFunc("GET /editor/{id}",         pages.Editor)
    http.HandleFunc("GET /editor/{id}/{page}",  pages.Editor)

    http.HandleFunc("GET /api/translations/{id}", apis.Translations)


    err := http.ListenAndServe("localhost:8080", nil)
    msg.Println(err)
}
