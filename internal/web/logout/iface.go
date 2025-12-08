package logout

import (
	"net/http"
)

type LogoutPage interface {
	HandleLogoutRequest(w http.ResponseWriter, r *http.Request)
}
