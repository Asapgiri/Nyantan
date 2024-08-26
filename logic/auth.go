package logic

type Auth struct {
    Error string
    User  string
    Roles []string
}

func Authenticate(a *Auth) {
    *a = Auth_get_user()
}

func Auth_get_user() Auth {
    return Auth{
        Error: "why, hello there",
        User: "Asapgiri",
    }
}

func Auth_logout() {}
