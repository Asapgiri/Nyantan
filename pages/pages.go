package pages

import (
	"io"
	"net/http"
	"nyantan/dbase"
	"nyantan/logger"
	"nyantan/logic"
	"strconv"
    "nyantan/session"
)

var log = logger.Logger {
    Color: logger.Colors.Red,
    Pretext: "pages",
}

func Base_auth_and_render(w http.ResponseWriter, r *http.Request, path string) (string, string) {
    session.SetPath(r.URL.Path)
    session.Authenticate(r)
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
    uname := r.FormValue("form[userName]")
    upass := r.FormValue("form[userPass]")

    if "" != uname {
        user := logic.Auth_login(uname, upass)
        if user.Id != "" {
            session.New(w, r, uname)
        } else {
            session.SetError("Auth Error")
        }
    } else {
        session.SetError("")
    }

    if "" == session.Get().Auth.Username {
        Render(w, fil, nil)
    } else {
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func Register(w http.ResponseWriter, r *http.Request) {
    fil, _ := Base_auth_and_render(w, r, "register.html")
    uname := r.FormValue("form[userName]")
    upass := r.FormValue("form[userPass]")

    if "" != uname {
        if logic.Auth_register(uname, upass) {
            session.New(w, r, uname)
        } else {
            session.SetError("Cannot Register")
        }
    } else {
        session.SetError("")
    }

    if "" == session.Get().Auth.Username {
        Render(w, fil, nil)
    } else {
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func Translate(w http.ResponseWriter, r *http.Request) {
    fil, _ := Base_auth_and_render(w, r, "translate.html")
    translations, err := logic.List_translations(session.Get().Auth)
    if err != nil {
        log.Println(err)
    }
    Render(w, fil, translations)
}

func Logout(w http.ResponseWriter, r *http.Request) {
    session.Authenticate(r)
    session.Delete(w, r)
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
    selected := dbase.Translation{}
    err := selected.Select(r.PathValue("id"))
    if nil != err {
        log.Println(err)
        base_error_render(w, r)
        return
    }

    fil, _ := Base_auth_and_render(w, r, "trans.html")
    pre_rendered := Pre_render(fil, selected)
    Render(w, pre_rendered, nil)
}

func Editor_list(w http.ResponseWriter, r *http.Request) {
    fil, _ := Base_auth_and_render(w, r, "edit-list.html")
    id := r.PathValue("id")

    edits, err := logic.List_edits(id)
    if nil != err {
        log.Println(err)
        base_error_render(w, r)
        return
    }

    epl := logic.Edit_page_list{
        TransId: id,
        Title: id,
        Link: logic.Generate_translation_link(id),
        PageCount: len(edits),
        Edits: edits,
    }

    pre_rendered := Pre_render(fil, epl)
    Render(w, pre_rendered, nil)
}

func Editor(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    page := r.PathValue("page")

    selected := dbase.Translation{}
    err := selected.Select(id)
    if nil != err {
        log.Println(err)
        base_error_render(w, r)
        return
    }

    if !logic.User_in_fandom(session.Get().Auth, selected.Fandom) {
        base_error_render(w, r)
        return
    }

    log.Println(id, page)
    if page == "" {
        Editor_list(w, r)
        return
    }

    page_index, err := strconv.Atoi(page)
    if nil != err {
        log.Println(err)
        base_error_render(w, r)
        return
    }

    edits, err := logic.Select_edit(id, page_index)
    if nil != err {
        log.Println(err)
        base_error_render(w, r)
        return
    }

    edit_list := logic.Edit_list{
        TransId:    selected.Id.Hex(),
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
