package apis

import (
	"io"
	"net/http"
	"nyantan/dbase"
	"nyantan/logic"
	"nyantan/pages"
	"nyantan/session"
	"strconv"
)

func Translations(w http.ResponseWriter, r *http.Request) {
    selected := dbase.Translation{}
    if nil != selected.Select(r.PathValue("id")) {
        fil, _ := pages.Base_auth_and_render(w, r, "not_found.html")
        pages.Render(w, fil, nil)
        return
    }

    fil, _ := pages.Base_auth_and_render(w, r, "trans.html")
    pre_rendered := pages.Pre_render(fil, selected)

    io.WriteString(w, pre_rendered)
}

func AddEditSnippet(w http.ResponseWriter, r *http.Request) {
    session.Authenticate(r)
    id := r.PathValue("id")
    page := r.PathValue("page")

    selected := dbase.Translation{}
    err := selected.Select(id)
    if nil != err && !logic.User_in_fandom(session.Get().Auth, selected.Fandom) {
        http.Error(w, "Not permitted!", http.StatusMethodNotAllowed)
        return
    }

    query := r.URL.Query()
    index, err1         := strconv.Atoi(query.Get("index"))
    height, err2        := strconv.ParseFloat(query.Get("height"), 32)
    width, err3         := strconv.ParseFloat(query.Get("width"), 32)
    x, err4             := strconv.ParseFloat(query.Get("x"), 32)
    y, err5             := strconv.ParseFloat(query.Get("y"), 32)
    page_index, err6    := strconv.Atoi(page)

    if nil != err1 || nil != err2 || nil != err3 || nil != err4 || nil != err5 || nil != err6 {
        http.Error(w, "Invalid parameter(s)!", http.StatusMethodNotAllowed)
        return
    }

    edit := logic.Edit{
        Index:  index,
        Rect:   dbase.Rectangle{
            X: float32(x),
            Y: float32(y),
            Height: float32(height),
            Width: float32(width),
        },
    }
    err = logic.NewEdit(session.Get().Auth, id, page_index, edit)
    if nil != err {
        io.WriteString(w, err.Error())
        return
    }

    io.WriteString(w, "OK.")
}
