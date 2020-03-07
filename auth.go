package main

import (
	"net/http"

	"github.com/video-server/api/session"
)

var HEADER_FILELD_SESSION = "X-Session-Id"
var HEADER_FILELD_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FILELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FILELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter) {

}
