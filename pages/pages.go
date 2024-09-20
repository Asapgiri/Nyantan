package pages

import (
	"io"
	"net/http"
	"nihility/dbase"
	"nihility/logger"
	"nihility/logic"
	"strconv"
)

var log = logger.Logger {
    Color: logger.Colors.Red,
    Pretext: "pages",
}

type sessioner struct {
    Auth logic.Auth
    Main string
    Path string
    Dto any
}
//FIXME: Handle fully separately in every function/session!!
var session sessioner

var artifact_path string = "artifacts/"
var html_path string = "html/"
var base_template_path string = html_path + "base.html"

func Base_auth_and_render(w http.ResponseWriter, r *http.Request, path string) (string, string) {
    session.Path = r.URL.Path
    logic.Authenticate(&session.Auth)
    return read_artifact(path, w.Header())
}

// =====================================================================================================================
// Basic functios

func Root(w http.ResponseWriter, r *http.Request) {
    if "/" == r.URL.Path {
        fil, _ := Base_auth_and_render(w, r, "index.html")
        Render(w, fil, nil)
    } else {
        Unexpected(w, r)
    }
}

func Login(w http.ResponseWriter, r *http.Request) {
    fil, _ := Base_auth_and_render(w, r, "login.html")

    if "" == session.Auth.User {
        Render(w, fil, nil)
    } else {
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func Translate(w http.ResponseWriter, r *http.Request) {
    fil, _ := Base_auth_and_render(w, r, "translate.html")
    translations, _ := dbase.List_translations()
    Render(w, fil, translations)
}

func Logout(w http.ResponseWriter, r *http.Request) {
    logic.Authenticate(&session.Auth)
    logic.Auth_logout()
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Unexpected(w http.ResponseWriter, r *http.Request) {
    fil, typ := Base_auth_and_render(w, r, r.URL.Path)

    if "text" == typ {
        Render(w, fil, nil)
    } else {
        io.WriteString(w, fil)
    }
}

// =====================================================================================================================
// "Smart" functios

func base_error_render(w http.ResponseWriter, r *http.Request) {
    fil, _ := Base_auth_and_render(w, r, "not_found.html")
    Render(w, fil, nil)
}

func Translation(w http.ResponseWriter, r *http.Request) {
    selected, err := dbase.Select_translation(r.PathValue("id"))
    if nil != err {
        base_error_render(w, r)
        return
    }

    fil, _ := Base_auth_and_render(w, r, "trans.html")
    pre_rendered := Pre_render(fil, selected)
    Render(w, pre_rendered, nil)
}

func Editor_list(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    edits, err := dbase.List_edits(id)
    if nil != err {
        base_error_render(w, r)
        return
    }

    epl := dbase.Edit_page_list{
        TransId: id,
        Title: id,
        Link: logic.Generate_translation_link(id),
        PageCount: len(edits),
        Edits: edits,
    }

    fil, _ := Base_auth_and_render(w, r, "edit-list.html")
    pre_rendered := Pre_render(fil, epl)
    Render(w, pre_rendered, nil)
}

func Editor(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    page := r.PathValue("page")

    log.Println(id, page)
    if page == "" {
        Editor_list(w, r)
        return
    }

    page_index, err := strconv.Atoi(page)
    if nil != err {
        base_error_render(w, r)
        return
    }

    // FIXME: should be only needed data
    selected, err := dbase.Select_translation(id)
    if nil != err {
        base_error_render(w, r)
        return
    }

    edits, err := dbase.Select_edit(id, page_index)
    if nil != err {
        base_error_render(w, r)
        return
    }
    edit_list := dbase.Edit_list{
        TransId:    selected.Id,
        // FIXME: sould be setted with prefixes and paths
        Title:      selected.Title,
        Link:       selected.Link,
        Image:      logic.Generate_translation_image_path_original(id, page_index),
        Page:       page_index,
        PageCount:  selected.Pages,
        Edits:      edits,
    }

    fil, _ := Base_auth_and_render(w, r, "editor.html")
    pre_rendered := Pre_render(fil, edit_list)
    Render(w, pre_rendered, nil)
}
