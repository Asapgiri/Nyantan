package dbase

type roles struct {
    USER        string
    MODERATOR   string
    ADMIN       string
}

var Roles = roles {
    USER:       "USER",
    MODERATOR:  "MODERATOR",
    ADMIN:      "ADMIN",
}
