package session

import (
	"net/http"
	"nyantan/logic"
	"nyantan/logger"
	"os"

	"github.com/gorilla/sessions"
)

var log = logger.Logger {
    Color: logger.Colors.Yellow,
    Pretext: "session",
}

type Sessioner struct {
    Auth logic.Auth
    Main string
    Path string
    Dto any
}
//FIXME: Handle fully separately in every function/session!!
//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var store = sessions.NewCookieStore([]byte(os.Getenv("NYANTAN_SESSION_KEY")))
var session Sessioner
var sessionName = "uname"

func Authenticate(r *http.Request) {
    // TODO: Add request aut header
    real_session, _ := store.Get(r, sessionName)
    uname, _ := real_session.Values[sessionName].(string)

    session.Auth.Username = uname
    logic.Authenticate(&session.Auth)
}

func New(w http.ResponseWriter, r *http.Request, uname string) {
    // FIXME: Store auth headers in database with associated user
    store.MaxAge(86400)
    rsess, _ := store.New(r, sessionName)

    rsess.Values[sessionName] = uname
    rsess.Save(r, w)
    session.Auth.Username = uname
}

func Delete(w http.ResponseWriter, r *http.Request) {
    store.MaxAge(-1)
}


// =====================================================================================================================
// Getters

func Get() Sessioner {
    return session
}

// =====================================================================================================================
// Setters

func SetPath(path string) {
    session.Path = path
}

func SetError(msg string) {
    session.Auth.Username = ""
    session.Auth.Error = msg
}

func SetMain(main string) {
    session.Main = main
}

func SetDto(dto any) {
    session.Dto = dto
}
