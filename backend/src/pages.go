package main

import (
	"net/http"
    "io"
	"strconv"
	"strings"
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
    translations, _ := list_translations()
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

func base_error_render(w http.ResponseWriter, r *http.Request) {
    fil, _ := base_auth_and_render(w, r, "not_found.html")
    render(w, fil, nil)
}

func get_translation(w http.ResponseWriter, r *http.Request, id string) {
    selected, err := select_translation(id)
    if nil != err {
        base_error_render(w, r)
        return
    }

    fil, _ := base_auth_and_render(w, r, "trans.html")
    pre_rendered := pre_render(fil, selected)
    render(w, pre_rendered, nil)
}

func get_editor_list(w http.ResponseWriter, r *http.Request, id string) {
    edits, err := list_edits(id)
    if nil != err {
        base_error_render(w, r)
        return
    }

    epl := edit_page_list{
        TransId: id,
        Title: id,
        Link: generate_translation_link(id),
        PageCount: len(edits),
        Edits: edits,
    }

    fil, _ := base_auth_and_render(w, r, "edit-list.html")
    pre_rendered := pre_render(fil, epl)
    render(w, pre_rendered, nil)
}

func get_editor(w http.ResponseWriter, r *http.Request, id string) {
    splits := strings.Split(id, "/")
    print_r(splits)
    if 2 > len(splits) {
        get_editor_list(w, r, id)
        return
    }

    t_id := splits[0]
    page_index, err := strconv.Atoi(splits[1])
    if nil != err {
        base_error_render(w, r)
        return
    }

    // FIXME: should be only needed data
    selected, err := select_translation(t_id)
    if nil != err {
        base_error_render(w, r)
        return
    }

    edits, err := select_edit(id, page_index)
    if nil != err {
        base_error_render(w, r)
        return
    }
    edit_list := edit_list{
        TransId:    selected.Id,
        // FIXME: sould be setted with prefixes and paths
        Title:      selected.Title,
        Link:       selected.Link,
        Image:      selected.Cover,
        Page:       page_index,
        PageCount:  selected.Pages,
        Edits:      edits,
    }

    fil, _ := base_auth_and_render(w, r, "editor.html")
    pre_rendered := pre_render(fil, edit_list)
    render(w, pre_rendered, nil)
}
