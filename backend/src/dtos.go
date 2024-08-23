package main

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

type translation struct {
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

type edit struct {
    Rect        rectangle
    LastUpdate  int64
    Accepter    string
    Accepted    bool
    Original    accepter
    Translated  accepter
}

type edit_list struct {
    TransId     string
    Title       string
    Link        string
    Image       string
    Page        int
    PageCount   int
    Edits       []edit
}

// =====================================================================================================================
// Edit list

type edit_page_list_item struct {
    Page        int
    LastUpdate  int64
    IImage      string
    Progress    progress
    Users       []user
    Accepter    string
    Accepted    bool
}

type edit_page_list struct {
    TransId     string
    Title       string
    Link        string
    PageCount   int
    Edits       []edit_page_list_item
}
