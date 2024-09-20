package apis

import (
	"io"
	"net/http"
	"nihility/dbase"
	"nihility/pages"
)

func Translations(w http.ResponseWriter, r *http.Request) {
    selected, err := dbase.Select_translation(r.PathValue("id"))
    if nil != err {
        fil, _ := pages.Base_auth_and_render(w, r, "not_found.html")
        pages.Render(w, fil, nil)
        return
    }

    fil, _ := pages.Base_auth_and_render(w, r, "trans.html")
    pre_rendered := pages.Pre_render(fil, selected)

    io.WriteString(w, pre_rendered)
}
