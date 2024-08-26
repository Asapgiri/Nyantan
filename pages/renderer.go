package pages

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var file_types = map[string]string {
    "html": "text",
    "css":  "text",
}

func read_artifact(path string, header http.Header) (string, string) {
    var dir_path string

    ex, err := os.Executable()
    if nil != err {
        panic(err)
    }

    parts := strings.Split(path, ".")
    file_type := parts[len(parts)-1]
    if "html" == file_type {
        dir_path = filepath.Dir(ex) + "/" + html_path
    } else {
        dir_path = filepath.Dir(ex) + "/" + artifact_path
    }

    file_read, err := os.ReadFile(dir_path + path)
    if nil != err {
        not_found, _ := os.ReadFile(filepath.Dir(ex) + "/" + html_path + "not_found.html")
        return string(not_found), "text"
    }

    if nil != header {
        _, file_ok := file_types[file_type]
        if file_ok {
            header.Set("Content-Type", file_types[file_type] + "/" + file_type)
        }
    }

    return string(file_read), file_type
}

func Render(w http.ResponseWriter, temp string, dto any) {
    tmp, err := template.ParseFiles(base_template_path)
    if nil != err {
        io.WriteString(w, "Templating error!")
        return
    }

    session.Main = temp
    session.Dto = dto

    var tpl bytes.Buffer
    tmp.Execute(&tpl, session)
    main, err := template.New("Main").Parse(tpl.String())
    if nil != err {
        io.WriteString(w, "Templating error 2!" + err.Error())
        return
    }

    main.Execute(w, session)
}

// Prerender does not support session if you don't pass it...
func Pre_render(temp string, dto any) string {
    var tpl bytes.Buffer

    tmp, err := template.New("Dto").Parse(temp)
    if nil != err {
        return err.Error()
    }
    tmp.Execute(&tpl, dto)

    return tpl.String()
}
