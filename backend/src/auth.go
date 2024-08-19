package main

type Auth struct {
    Error string
    User  string
    Roles []string
}

func authenticate(a *Auth) {
    *a = auth_get_user()
}

func auth_get_user() Auth {
    return Auth{
        Error: "why, hello there",
        User: "Asapgiri",
    }
}

func auth_logout() {}
