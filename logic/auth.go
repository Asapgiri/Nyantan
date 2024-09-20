package logic

import (
    "golang.org/x/crypto/bcrypt"
)

type user struct {
    Id              string
    Name            string
    PasswordHash    string
    Groups          map[string][]string
    SiteRoles       []string
}

var Example_user_auts = []user {
}

type Auth struct {
    Error string
    User  string
    Roles []string
}

func Authenticate(a *Auth) {
    //*a = Auth_get_user()
}

func Auth_get_user() Auth {
    return Auth{
        Error: "",
        User: "",
    }
}

func Auth_logout() {}

func find_user(id string) user {
    for _, user := range Example_user_auts {
        if user.Id == id {
            return user
        }
    }
    return user{}
}

func Auth_register(id string, password_clear string) bool {
    // FIXME: Do the actual database search!
    if find_user(id).Id != "" {
        return false
    }

    pwh, _ := bcrypt.GenerateFromPassword([]byte(password_clear), 0)
    new_user := user{
        Id: id,
        PasswordHash: string(pwh),
    }

    // FIXME: Save to database
    Example_user_auts = append(Example_user_auts, new_user)
    log.Printf("Registerd with %s:%s\n", new_user.Id, new_user.PasswordHash)

    return true
}

func Auth_login(id string, password_clear string) user {
    fuser := find_user(id)

    if fuser.Id == "" || nil != bcrypt.CompareHashAndPassword([]byte(fuser.PasswordHash), []byte(password_clear)) {
        return user{}
    }

    return fuser
}
