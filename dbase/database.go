package dbase

import (
	"errors"
	"nihility/logger"
	"nihility/logic"
	"time"
)

var log = logger.Logger {
    Color: logger.Colors.Purple,
    Pretext: "database",
}

// =====================================================================================================================
// Listings

var color_list = []string{
    "dark",
    "secondary",
    "info",
    "success",
}

func generate_progress(percentage float32) progress {
    return progress{
        Color: color_list[int(float32(len(color_list)) * percentage)],
        Percentage: int(percentage * 100),
    }
}

func List_translations() ([]Translation, error) {
    return Example_translations, nil
}

func List_edits(translation_id string) ([]Edit_page_list_item, error) {
    trans, _ := Select_translation(translation_id)
    edits := make([]Edit_page_list_item, trans.Pages)
    for i := 0; i < len(edits); i++ {
        edits[i].Page = i
        edits[i].Accepted = false
        edits[i].Accepter = ""
        edits[i].Users  = trans.Users
        edits[i].Progress = generate_progress(float32(i) / float32(len(edits)))
        edits[i].LastUpdate = int64(i) + time.Now().Unix()
        // FIXME: ...
        edits[i].IImage = logic.Generate_translation_image_path_original(translation_id, i)
    }

    return edits, nil
}

// =====================================================================================================================
// Selects

func Select_translation(id string) (Translation, error) {
    for _, trans := range Example_translations {
        log.Println(trans.Title, "==", id)
        if trans.Title == id || trans.Id == id {
            return trans, nil
        }
    }

    return Translation{}, errors.New("Not found!")
}

func Select_edit(translation_id string, page int) ([]Edit, error) {
    _ = translation_id
    _ = page
    return Example_edits, nil
}
