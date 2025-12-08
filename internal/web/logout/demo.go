package logout

import (
	"context"
	"net/http"
)

type DemoLogoutPageLogic interface {
	GetSessionFromCookie(ctx context.Context, c *http.Cookie) (string, error)
	DeauthSession(ctx context.Context, s string) error
}

type DemoLogoutPage struct {
	logic DemoLogoutPageLogic
}

func NewDemoLogoutPage(mux *http.ServeMux, l DemoLogoutPageLogic) LogoutPage {
	p := DemoLogoutPage{
		logic: l,
	}

	mux.HandleFunc("/logout", p.HandleLogoutRequest)

	return &p
}

func (p *DemoLogoutPage) HandleLogoutRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session_cookie, err := r.Cookie("session")
	if err != nil {
		w.Write([]byte("Bad cookie"))
		return
	}
	session, err := p.logic.GetSessionFromCookie(ctx, session_cookie)
	if err != nil {
		errmsg := "DemoLogoutPage: Failed to extract session from cookie:\n" + err.Error()
		w.Write([]byte(errmsg))
		return
	}

	err = p.logic.DeauthSession(ctx, session)
	if err != nil {
		errmsg := "DemoLogoutPage: Failed to deauth session:\n" + err.Error()
		w.Write([]byte(errmsg))
		return
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
