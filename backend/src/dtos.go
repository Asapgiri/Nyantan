package main

type group struct {
    Id      string
    Name    string
}

type progress struct {
    Color       string
    Percentage  string
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
