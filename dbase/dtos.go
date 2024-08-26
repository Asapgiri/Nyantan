package dbase

// =====================================================================================================================
// Translation

type group struct {
    Id      string
    Name    string
}

type progress struct {
    Color       string
    Percentage  int
}

type role struct {
    Id      string
    Name    string
    Color   string
}

type user struct {
    Id      string
    Name    string
    Roles   []role
}

type external_link struct {
    Website     string
    Title       string
    Link        string
    Language    string
}

type Translation struct {
    Id          string
    Date        string
    Group       group
    Title       string
    Tags        []string
    Parodies    []string
    Characters  []string
    Artist      []string
    Pages       int
    Cover       string
    Link        string
    Visible     bool
    Views       string
    Progress    progress
    Users       []user
    Externals   []external_link
}

// =====================================================================================================================
// Edits

type rectangle struct {
    X       float32
    Y       float32
    Width   float32
    Height  float32
}

type selectable struct {
    Text        string
    Author      string
    Accepter    string
    Selected    bool
    Accepted    bool
    Date        int64
}

type accepter struct {
    SIndex      int
    Accepted    bool
    Accepter    string
    Date        int64
    List        []selectable
}

type Edit struct {
    Rect        rectangle
    LastUpdate  int64
    Accepter    string
    Accepted    bool
    Original    accepter
    Translated  accepter
}

type Edit_list struct {
    TransId     string
    Title       string
    Link        string
    Image       string
    Page        int
    PageCount   int
    Edits       []Edit
}

// =====================================================================================================================
// Edit list

type Edit_page_list_item struct {
    Page        int
    LastUpdate  int64
    IImage      string
    Progress    progress
    Users       []user
    Accepter    string
    Accepted    bool
}

type Edit_page_list struct {
    TransId     string
    Title       string
    Link        string
    PageCount   int
    Edits       []Edit_page_list_item
}
