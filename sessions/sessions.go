package sessions

import (
	"net/http"
    gsessions "github.com/gorilla/sessions"
)

var store = gsessions.NewCookieStore([]byte("random"))

func Get(req *http.Request) (*gsessions.Session, error) {
    return store.Get(req, "default-session-name")
}

func GetEmail(req *http.Request, email string) (*gsessions.Session, error) {
    return store.Get(req, email)
}