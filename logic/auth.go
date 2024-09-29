package logic

import (
	"nihility/dbase"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
    Error       string
    Username    string
    Groups      []string
    Roles       []string
}

func Authenticate(a *Auth) {
    if a.Username != "" {
        user := dbase.User{Id: a.Username}
        err := user.Find()
        if err != nil {
            *a = Auth{}
            return
        }

        a.Username = user.Id
        a.Roles = user.SiteRoles
        a.Error = ""
    }
}

func Auth_logout() {}

func Auth_register(id string, password_clear string) bool {
    new_user := dbase.User{Id: id}

    if new_user.Find() == nil {
        log.Println("User exists: " + id)
        return false
    }

    pwh, _ := bcrypt.GenerateFromPassword([]byte(password_clear), 0)
    new_user = dbase.User{
        Id: id,
        PasswordHash: string(pwh),
        SiteRoles: []string{dbase.Roles.USER},
    }

    new_user.Register()
    log.Printf("Registerd with %s:%s\n", new_user.Id, string(pwh))

    return true
}

func Auth_login(id string, password_clear string) dbase.User {
    user := dbase.User{Id: id}
    err := user.Find()

    log.Println(user)
    if err != nil || nil != bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password_clear)) {
        log.Println("Returnning with error, not user")
        log.Println(err)
        return dbase.User{}
    }

    return user
}

func User_in_fandom(a Auth, fandom string) bool {
    user := dbase.User{Id: a.Username}
    user.Find()

    for _, r := range user.Fandoms() {
        if r == fandom {
            return true
        }
    }

    return false
}
