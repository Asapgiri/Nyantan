package main

import (
	"net/http"
	"nihility/apis"
	"nihility/logger"
	"nihility/pages"
	"strings"
)

var log = logger.Logger {
    Color: logger.Colors.Blue,
    Pretext: "router",
}

var routes = map[string]func(w http.ResponseWriter, r *http.Request) {
    "/":            pages.Root,
    "/index":       pages.Root,

    // Authentications
    "/login":       pages.Login,
    "/logout":      pages.Logout,

    // Translation related
    "/translate":   pages.Translate,
}

var funny_pages = map[string]func(w http.ResponseWriter, r *http.Request, id string) {
    // Translation related
    "/trans":               pages.Translation,
    "/editor":              pages.Editor,

    // Apis
    "/api/translations":    apis.Translations,
}

func Router(path string, w http.ResponseWriter, r *http.Request) {
    page := routes[path]
    if nil != page {
        page(w, r)
        return
    }

    for k, page := range funny_pages {
        log.Println(k, path)
        if strings.HasPrefix(path, k) {
            ffile := strings.TrimPrefix(path, k) //strip /pretag
            ffile = strings.TrimPrefix(ffile, "/") //strip trailing / if persists
            page(w, r, ffile)
            return
        }
    }

    pages.Unexpected(w, r)
}
