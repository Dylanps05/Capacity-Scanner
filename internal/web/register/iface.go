package register

import (
	"net/http"
)

type RegisterPage interface {
	HandleRegisterPageRequest(w http.ResponseWriter, r *http.Request)
	HandleRegisterSubmission(w http.ResponseWriter, r *http.Request)
}
