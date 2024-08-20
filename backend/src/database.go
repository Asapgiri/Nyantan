package main

import "fmt"

func translations_list() ([]translation, error) {
    return Example_translations, nil
}

func translations_select(id string) (translation, error) {
    for _, trans := range Example_translations {
        fmt.Print(trans.Title)
        fmt.Print(" == ")
        fmt.Println(id)
        if trans.Title == id || trans.Id == id {
            return trans, nil
        }
    }

    return translation{}, nil
}
