package logic

import (
	"fmt"
	"nihility/logger"
	"strings"
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

func List_images(path string) [100]string {
    // FIXME:
    log.Println("Listing images for:", path)
    temp := [100]string{}
    for i := 0; i < 100; i++ {
        val := fmt.Sprintf("%s/%03d.png", path, i)
        log.Println(i, ": Adding value =", val)
        temp[i] = val
    }
    return temp
}

func Generate_translation_image_path_original(id string, index int) string {
    log.Println("Generating path for:", id, index)
    return List_images(strings.Join([]string{translatioins_image_base_path, id, orig_dir}, "/"))[index]
}

func Generate_translation_image_path_translated(id string, index int) string {
    return List_images(strings.Join([]string{translatioins_image_base_path, id, trans_dir}, "/"))[index]
}
