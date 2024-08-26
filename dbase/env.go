package dbase

var Example_translations =  []Translation{
    {
        Id:         "2",
        Date:       "Aug. 25",
        Group:      group{ Id: "nihility", Name: "Nihility" },
        Title:      "Devilchi 224*bonus chapter",
        Tags:       []string{},
        Parodies:   []string{"original"},
        Characters: []string{},
        Artist:     []string{},
        Pages:      12,
        Visible:    false,
        Views:      "",
        Cover:      "/trassets/Devilchi%20224*bonus%20chapter/orig/000.png",
        Link:       "/trans/Devilchi%20224*bonus%20chapter",
        Progress:   progress{ Color: "info", Percentage: 21 },
        Users:  []user{
            {Id: "asapgiri", Name: "Asapgiri",
            Roles: []role{
                {Id: "uploader", Name: "Uploader", Color: "primary"},
                {Id: "translator", Name: "Translator", Color: "info"},
                {Id: "proofreader", Name: "Proofreader", Color: "secondary"},
            }},
            {Id: "waster", Name: "Wasted",
            Roles: []role{
                {Id: "translator", Name: "Translator", Color: "info"},
                {Id: "editor", Name: "Editor", Color: "warning"},
            }},
        },
        Externals: []external_link{
        },
    },
}

var Example_edits = []Edit{
    {
        Accepted:   false,
        Accepter:   "",
        LastUpdate: 1,
        Rect:   rectangle{
            X:      45.0,
            Y:      45.0,
            Width:  133.0,
            Height: 155.0,
        },
        Original: accepter{
            SIndex:     0,
            Date:       2,
            Accepted:   true,
            Accepter:   "asapgiri",
            List:       []selectable{
                {
                    Text:       "腹ご",
                    Selected:   true,
                    Author:     "asapgiri",
                    Accepter:   "",
                    Accepted:   false,
                    Date:       3,
                },
                {
                    Text:       "なしに",
                    Selected:   false,
                    Author:     "asapgiri",
                    Accepter:   "wasted",
                    Accepted:   true,
                    Date:       4,
                },
            },
        },
        Translated: accepter{
            SIndex:     1,
            Date:       5,
            Accepted:   false,
            Accepter:   "",
            List:       []selectable{
                {
                    Text:       "text",
                    Selected:   false,
                    Author:     "asapgiri",
                    Accepter:   "",
                    Accepted:   false,
                    Date:       6,
                },
                {
                    Text:       "random",
                    Selected:   true,
                    Author:     "asapgiri",
                    Accepter:   "wasted",
                    Accepted:   true,
                    Date:       7,
                },
                {
                    Text:       "selectable",
                    Selected:   false,
                    Author:     "asapgiri",
                    Accepter:   "",
                    Accepted:   false,
                    Date:       8,
                },
            },
        },
    },
    {
        Accepted:   true,
        Accepter:   "",
        LastUpdate: 9,
        Rect:   rectangle{
            X:      545.0,
            Y:      45.0,
            Width:  133.0,
            Height: 155.0,
        },
        Original: accepter{
            SIndex:     1,
            Date:       10,
            Accepted:   true,
            Accepter:   "asapgiri",
            List:       []selectable{
                {
                    Text:       "腹ご",
                    Selected:   false,
                    Author:     "asapgiri",
                    Accepter:   "",
                    Accepted:   false,
                    Date:       11,
                },
                {
                    Text:       "なしに",
                    Selected:   true,
                    Author:     "asapgiri",
                    Accepter:   "wasted",
                    Accepted:   true,
                    Date:       12,
                },
            },
        },
        Translated: accepter{
            SIndex:     2,
            Date:       13,
            Accepted:   true,
            Accepter:   "",
            List:       []selectable{
                {
                    Text:       "text",
                    Selected:   false,
                    Author:     "asapgiri",
                    Accepter:   "",
                    Accepted:   false,
                    Date:       14,
                },
                {
                    Text:       "random",
                    Selected:   false,
                    Author:     "asapgiri",
                    Accepter:   "wasted",
                    Accepted:   true,
                    Date:       14,
                },
                {
                    Text:       "selectable",
                    Selected:   true,
                    Author:     "asapgiri",
                    Accepter:   "",
                    Accepted:   false,
                    Date:       15,
                },
            },
        },
    },
}
