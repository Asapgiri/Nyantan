package logic

import "nyantan/dbase"



// =====================================================================================================================
// Edit list

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
    Rect        dbase.Rectangle
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

type Edit_page_list_item struct {
    Page        int
    LastUpdate  int64
    IImage      string
    Progress    dbase.Progress
    Users       []dbase.User
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
