package login

import (
	"net/http"
)

type LoginPage interface {
	HandleLoginPageRequest(w http.ResponseWriter, r *http.Request)
	HandleLoginRequest(w http.ResponseWriter, r *http.Request)
}
