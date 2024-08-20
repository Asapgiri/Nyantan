package main

import (
	"io"
	"net/http"
)

type sessioner struct {
    Auth Auth
    Main string
    Path string
    Dto any
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

// =====================================================================================================================
// Basic functios

func get_root(w http.ResponseWriter, r *http.Request) {
    fil, _ := base_auth_and_render(w, r, "index.html")
    render(w, fil, nil)
}

func get_login(w http.ResponseWriter, r *http.Request) {
    fil, _ := base_auth_and_render(w, r, "login.html")

    if "" == session.Auth.User {
        render(w, fil, nil)
    } else {
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func get_translate(w http.ResponseWriter, r *http.Request) {
    fil, _ := base_auth_and_render(w, r, "translate.html")
    translations, _ := translations_list()
    render(w, fil, translations)
}

func get_logout(w http.ResponseWriter, r *http.Request) {
    authenticate(&session.Auth)
    auth_logout()
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func unexpected(w http.ResponseWriter, r *http.Request) {
    fil, typ := base_auth_and_render(w, r, r.URL.Path)

    if "text" == typ {
        render(w, fil, nil)
    } else {
        io.WriteString(w, fil)
    }
}

// =====================================================================================================================
// "Smart" functios

func get_translation(w http.ResponseWriter, r *http.Request, id string) {
    selected, err := translations_select(id)
    if nil != err {
        fil, _ := base_auth_and_render(w, r, "not_found.html")
        render(w, fil, nil)
        return
    }

    fil, _ := base_auth_and_render(w, r, "trans.html")
    pre_rendered := pre_render(fil, selected)
    render(w, pre_rendered, nil)
}

func get_editor(w http.ResponseWriter, r *http.Request, id string) {
    selected, err := translations_select(id)
    if nil != err {
        fil, _ := base_auth_and_render(w, r, "not_found.html")
        render(w, fil, nil)
        return
    }

    fil, _ := base_auth_and_render(w, r, "editor.html")
    pre_rendered := pre_render(fil, selected)
    render(w, pre_rendered, nil)
}
