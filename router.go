package main

import (
	"net/http"
	"nyantan/apis"
	"nyantan/pages"
)

func setup_routes() {
    http.HandleFunc("GET /",                    pages.Root)
    http.HandleFunc("GET /index",               pages.Root)
    http.HandleFunc("GET /index.html",          pages.Root)

    // Authentications
    http.HandleFunc("GET  /login",              pages.Login)
    http.HandleFunc("POST /login",              pages.Login)
    http.HandleFunc("GET  /register",           pages.Register)
    http.HandleFunc("POST /register",           pages.Register)
    http.HandleFunc("GET  /logout",             pages.Logout)

    // Translation related
    http.HandleFunc("GET /translate",           pages.Translate)

    http.HandleFunc("GET /trans/{id}",          pages.Translation)
    http.HandleFunc("GET /editor/{id}",         pages.Editor)
    http.HandleFunc("GET /editor/{id}/{page}",  pages.Editor)

    http.HandleFunc("GET /api/translations/{id}",   apis.Translations)
    http.HandleFunc("GET /api/editor/{id}/{page}",  apis.AddEditSnippet)
}
