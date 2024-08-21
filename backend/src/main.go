package main


import (
    "fmt"
    "net/http"
    "strings"
)

type Hello struct{}

var pages = map[string]func(w http.ResponseWriter, r *http.Request) {
    "/":            get_root,
    "/index":       get_root,

    // Authentications
    "/login":       get_login,
    "/logout":      get_logout,

    // Translation related
    "/translate":   get_translate,
}

var funny_pages = map[string]func(w http.ResponseWriter, r *http.Request, id string) {
    // Translation related
    "/trans":               get_translation,
    "/editor":              get_editor,

    // Apis
    "/api/translations":    api_translations,
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var path = r.URL.Path
    fmt.Printf("Serving request: %s %#v\n", r.Header.Get("X-Forwarded-For"), path)

    page := pages[path]
    if nil != page {
        page(w, r)
        return
    }

    for k, page := range funny_pages {
        fmt.Println(k)
        fmt.Println(path)
        if strings.HasPrefix(path, k) {
            ffile := strings.TrimPrefix(path, k) //strip /pretag
            ffile = strings.TrimPrefix(ffile, "/") //strip trailing / if persists
            page(w, r, ffile)
            return
        }
    }

    unexpected(w, r)
}

func main() {
    var h Hello
    err := http.ListenAndServe("localhost:8080", h)
    fmt.Println(err)
}
