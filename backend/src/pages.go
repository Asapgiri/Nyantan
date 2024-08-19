package main

import (
	"io"
	"net/http"
)

type sessioner struct {
    Auth Auth
    Main string
    Path string
}
//FIXME: Handle fully separately in every function/session!!
var session sessioner

var artifact_path string = "artifacts/"
var html_path string = "html/"
var base_template_path string = html_path + "base.html"

func base_auth_and_render(w http.ResponseWriter, r *http.Request, path string) (string, string) {
    session.Path = r.URL.Path
    authenticate(&session.Auth)
    return read_artifact(path, w.Header())
}

func get_root(w http.ResponseWriter, r *http.Request) {
    fil, _ := base_auth_and_render(w, r, "index.html")
    render(w, fil)
}

func get_login(w http.ResponseWriter, r *http.Request) {
    fil, _ := base_auth_and_render(w, r, "login.html")

    if "" == session.Auth.User {
        render(w, fil)
    } else {
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func get_translate(w http.ResponseWriter, r *http.Request) {
    fil, _ := base_auth_and_render(w, r, "translate.html")
    render(w, fil)
}

func get_logout(w http.ResponseWriter, r *http.Request) {
    authenticate(&session.Auth)
    auth_logout()
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func unexpected(w http.ResponseWriter, r *http.Request) {
    fil, typ := base_auth_and_render(w, r, r.URL.Path)

    if "text" == typ {
        render(w, fil)
    } else {
        io.WriteString(w, fil)
    }
}
