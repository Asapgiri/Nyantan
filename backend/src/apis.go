package main

import (
	"io"
	"net/http"
)

func api_translations(w http.ResponseWriter, r *http.Request, id string) {
    selected, err := select_translation(id)
    if nil != err {
        fil, _ := base_auth_and_render(w, r, "not_found.html")
        render(w, fil, nil)
        return
    }

    fil, _ := base_auth_and_render(w, r, "trans.html")
    pre_rendered := pre_render(fil, selected)

    io.WriteString(w, pre_rendered)
}
