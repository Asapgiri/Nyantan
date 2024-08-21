package main

import (
	"errors"
	"fmt"
)

// =====================================================================================================================
// Listings

func list_translations() ([]translation, error) {
    return Example_translations, nil
}

func list_edit(translation_id string, page int) ([]edit, error) {
    _ = translation_id
    _ = page
    return Example_edits, nil
}

// =====================================================================================================================
// Selects

func select_translations(id string) (translation, error) {
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
