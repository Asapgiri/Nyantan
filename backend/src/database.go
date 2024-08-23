package main

import (
	"errors"
	"fmt"
	"time"
)

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

func list_translations() ([]translation, error) {
    return Example_translations, nil
}

func list_edits(translation_id string) ([]edit_page_list_item, error) {
    trans, _ := select_translation(translation_id)
    edits := make([]edit_page_list_item, trans.Pages)
    for i := 0; i < len(edits); i++ {
        edits[i].Page = i
        edits[i].Accepted = false
        edits[i].Accepter = ""
        edits[i].Users  = trans.Users
        edits[i].Progress = generate_progress(float32(i) / float32(len(edits)))
        edits[i].LastUpdate = int64(i) + time.Now().Unix()
        // FIXME: ...
        edits[i].IImage = trans.Cover
    }

    return edits, nil
}

// =====================================================================================================================
// Selects

func select_translation(id string) (translation, error) {
    for _, trans := range Example_translations {
        fmt.Print(trans.Title)
        fmt.Print(" == ")
        fmt.Println(id)
        if trans.Title == id || trans.Id == id {
            return trans, nil
        }
    }

    return translation{}, errors.New("Not found!")
}

func select_edit(translation_id string, page int) ([]edit, error) {
    _ = translation_id
    _ = page
    return Example_edits, nil
}
