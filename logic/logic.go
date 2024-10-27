package logic

import (
	"cmp"
	"fmt"
	"nyantan/dbase"
	"nyantan/logger"
	"slices"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var log = logger.Logger {
    Color: logger.Colors.Cyan,
    Pretext: "logic",
}

var translations_path = "/trans"

var translatioins_image_base_path = "/trassets"
var orig_dir = "orig"
var trans_dir = "trld"

func Generate_translation_link(id string) string {
    return translations_path + "/" + id
}

func Generate_translation_image_path_original(id string, index int) string {
    log.Println("Generating path for:", id, index)
    path := strings.Join([]string{translatioins_image_base_path, id, orig_dir}, "/")
    return fmt.Sprintf("%s/%03d.png", path, index)
}

func Generate_translation_image_path_translated(id string, index int) string {
    path := strings.Join([]string{translatioins_image_base_path, id, trans_dir}, "/")
    return fmt.Sprintf("%s/%03d.png", path, index)
}


// =====================================================================================================================
// Translations and Edits

var color_list = []string{
    "dark",
    "secondary",
    "info",
    "success",
}

func generate_progress(percentage float32) dbase.Progress {
    return dbase.Progress{
        Color: color_list[int(float32(len(color_list)) * percentage)],
        Percentage: int(percentage * 100),
    }
}

func List_edits(translation_id string) ([]Edit_page_list_item, error) {
    trans := dbase.Translation{}
    trans.Select(translation_id)
    edits := make([]Edit_page_list_item, trans.Pages + 1) // +1 for cover
    users := make([]dbase.User, len(trans.Users))
    for i := 0; i < len(users); i++ {
        users[i].Id = trans.Users[i]
        users[i].Find()
    }
    for i := 0; i < len(edits); i++ {
        edits[i].Page = i
        edits[i].Accepted = false
        edits[i].Accepter = ""
        edits[i].Users  = users
        edits[i].Progress = generate_progress(float32(i) / float32(len(edits)))
        edits[i].LastUpdate = time.Now()
        // FIXME: ...
        edits[i].IImage = Generate_translation_image_path_original(translation_id, i)
    }

    return edits, nil
}

func List_translations(auth Auth) ([]dbase.Translation, error) {
    user := dbase.User{Id: auth.Username}
    user.Find()

    tr := dbase.Translation{}
    fandoms := user.Fandoms()
    flist := make([]string, len(fandoms))
    for i, f := range fandoms {
        flist[i] = f.Fandom
    }
    log.Println(flist)
    return tr.List(flist)
}

func Select_edit(id string, page int) ([]Edit, error) {
    var eret []Edit

    trans := dbase.Translation{}
    trans.Select(id)
    edits, err := dbase.Select_edit(trans.Id, page)
    if nil != err {
        return eret, err
    }
    log.Println("Edits:")
    log.Println(edits)
    log.Println(len(edits))

    eret = make([]Edit, len(edits))
    for i, e := range edits {
        eret[i].Id = e.Edit.Id.Hex()
        eret[i].LastUpdate = e.Edit.LastUpdated.Time()
        eret[i].Date = e.Edit.Date.Time()
        eret[i].Rect = e.Edit.Rectangle
        eret[i].Index = e.Edit.Index
        eret[i].Accepted = e.Edit.Accepted
        eret[i].Accepter = e.Edit.Accepter
        eret[i].Original = accepter{
            SIndex: e.Edit.Original,
            Accepted: false, // TODO
            Accepter: "", // TODO
            List: []selectable{},
        }
        eret[i].Translated = accepter{
            SIndex: e.Edit.Translated,
            Accepted: false, // TODO
            Accepter: "", // TODO
            List: []selectable{},
        }
        o, t := 0, 0
        for _, snip := range e.Snippets {
            if snip.Original {
                eret[i].Original.List = append(eret[i].Original.List, selectable{
                    Text: snip.Text,
                    Author: snip.Author,
                    Accepter: "", // TODO
                    Accepted: false, // TODO
                    Selected: o == eret[i].Original.SIndex,
                })
                o++
            } else {
                eret[i].Translated.List = append(eret[i].Translated.List, selectable{
                    Text: snip.Text,
                    Author: snip.Author,
                    Accepter: "", // TODO
                    Accepted: false, // TODO
                    Selected: t == eret[i].Translated.SIndex,
                })
                t++
            }
        }
    }

    slices.SortFunc(eret, func(i, j Edit) int {
        return cmp.Compare(i.Index, j.Index)
    })

    return eret, nil
}

func NewEdit(auth Auth, trId string, page int, edit Edit) error {
    trans := dbase.Translation{}
    err := trans.Select(trId)
    if nil != err {
        return err
    }

    newEdit := dbase.Edit{
        Id: primitive.NewObjectID(),
        Date: primitive.NewDateTimeFromTime(time.Now()),
        LastUpdated: primitive.NewDateTimeFromTime(time.Now()),
        Index: edit.Index,
        Fandom: trans.Fandom,
        Author: auth.Username,
        TranslationId: trans.Id,
        Page: page,
        Rectangle: edit.Rect,
    }
    log.Printf("Saving: ")
    log.Println(newEdit)
    err = newEdit.Add()

    return err
}

//func newSnippet(auth Auth, edit)
