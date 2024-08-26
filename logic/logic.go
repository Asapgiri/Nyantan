package logic

import "strings"

var translations_path = "/trans"
var translatioins_image_base_path = "/trassets"

var orig_dir = "orig"
var trans_dir = "trld"

func Generate_translation_link(id string) string {
    return translations_path + "/" + id
}

func List_images(path string) []string {
    // FIXME:
    return []string{}
}

func Generate_translation_image_path_original(id string, index int) string {
    return List_images(strings.Join([]string{translatioins_image_base_path, id, orig_dir}, "/"))[index]
}

func Generate_translation_image_path_translated(id string, index int) string {
    return List_images(strings.Join([]string{translatioins_image_base_path, id, trans_dir}, "/"))[index]
}