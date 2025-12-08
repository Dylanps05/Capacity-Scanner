package login

import (
	"context"
	"html/template"
	"log"
	"net/http"
)

type DemoLoginPageLogic interface {
	GetSessionFromCookie(ctx context.Context, c *http.Cookie) (string, error)
	SetSessionAuthByLogin(ctx context.Context, s string, u string, p string) error
}

type DemoLoginPage struct {
	logic     DemoLoginPageLogic
	loginPage *template.Template
}

func NewDemoLoginPage(mux *http.ServeMux, l DemoLoginPageLogic) LoginPage {
	p := &DemoLoginPage{
		logic: l,
	}

	var err error = nil
	p.loginPage, err = template.ParseFiles("./web/template/login.html")
	if err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("/login", p.HandleLoginPageRequest)
	mux.HandleFunc("/login/submit", p.HandleLoginRequest)

	return p
}

func (p *DemoLoginPage) HandleLoginPageRequest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	bad_auth := query.Get("badAuth")
	loginPageData := struct {
		LoginFailed bool
	}{
		LoginFailed: bad_auth == "true",
	}

	p.loginPage.Execute(w, loginPageData)
}

func (p *DemoLoginPage) HandleLoginRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	method_incorrect := r.Method != http.MethodPost
	if method_incorrect {
		w.Write([]byte("403"))
		return
	}

	session_cookie, err := r.Cookie("session")
	if err != nil {
		w.Write([]byte("Bad cookie"))
		return
	}
	session, err := p.logic.GetSessionFromCookie(ctx, session_cookie)

	username := r.FormValue("username")
	password := r.FormValue("password")
	err = p.logic.SetSessionAuthByLogin(ctx, session, username, password)
	if err != nil {
		http.Redirect(w, r, "/login?badAuth=true", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
