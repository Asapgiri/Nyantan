package dbase

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// =====================================================================================================================
// Translation

type Group struct {
    Id      string
    Name    string
}

type TrRole struct {
    Id          primitive.ObjectID `bson:"_id"`
    Fandom      string
    Username    string
    roles       string
}

type Progress struct {
    Color       string
    Percentage  int
}

type Role struct {
    Id      string
    Name    string
    Color   string
}

type User struct {
    _ID             primitive.ObjectID `bson:"_id"`
    Id              string `bson:"username"`
    PasswordHash    string `bson:"passhash"`
    Name            string
    SiteRoles       []string
}

type external_link struct {
    Website     string
    Title       string
    Link        string
    Language    string
}

type Translation struct {
    Id          primitive.ObjectID `bson:"_id"`
    Date        primitive.DateTime
    Fandom      string
    Title       string
    Tags        []string
    Parodies    []string
    Characters  []string
    Artist      []string
    Pages       int
    Cover       string
    Link        string
    Visible     bool
    Views       int
    Progress    Progress
    Users       []string
    Externals   []external_link
}

// =====================================================================================================================
// Edits

type Rectangle struct {
    X       float32
    Y       float32
    Width   float32
    Height  float32
}

type Edit struct {
    Id              primitive.ObjectID `bson:"_id"`
    Date            primitive.DateTime
    LastUpdated     primitive.DateTime
    Index           int
    Fandom          string
    Author          string
    Accepter        string
    Accepted        bool
    TranslationId   primitive.ObjectID
    Page            int
    Rectangle       Rectangle
    Original        int     // selected
    Translated      int     // selected
}

type Edit_snippet struct {
    Id              primitive.ObjectID `bson:"_id"`
    Edit            primitive.ObjectID
    Author          string
    Original        bool
    Text            string
}

type Edit_combined struct {
    Edit        Edit
    Snippets    []Edit_snippet
}
