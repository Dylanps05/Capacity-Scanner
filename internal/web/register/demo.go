package register

import (
	"context"
	"html/template"
	"log"
	"net/http"
)

type DemoRegisterPageLogic interface {
	CreateLogin(ctx context.Context, u string, p string) error
	GetSessionFromCookie(ctx context.Context, c *http.Cookie) (string, error)
}

type DemoRegisterPage struct {
	logic        DemoRegisterPageLogic
	registerPage *template.Template
}

func NewDemoRegisterPage(mux *http.ServeMux, l DemoRegisterPageLogic) RegisterPage {
	p := &DemoRegisterPage{
		logic: l,
	}

	var err error
	p.registerPage, err = template.ParseFiles("./web/template/register.html")
	if err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("/register", p.HandleRegisterPageRequest)
	mux.HandleFunc("/register/submit", p.HandleRegisterSubmission)

	return p
}

func (p *DemoRegisterPage) HandleRegisterPageRequest(w http.ResponseWriter, r *http.Request) {
	p.registerPage.Execute(w, nil)
}

func (p *DemoRegisterPage) HandleRegisterSubmission(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	method_incorrect := r.Method != http.MethodPost
	if method_incorrect {
		w.Write([]byte("403"))
		return
	}

	session_cookie, err := r.Cookie("session")
	if err != nil {
		w.Write([]byte("Bad cookie.\n" + err.Error()))
		return
	}
	_, err = p.logic.GetSessionFromCookie(ctx, session_cookie)
	if err != nil {
		w.Write([]byte("Failed to get session fom cookie.\n" + err.Error()))
		return
	}

	un := r.FormValue("username")
	pw := r.FormValue("password")
	err = p.logic.CreateLogin(ctx, un, pw)
	if err != nil {
		w.Write([]byte("Failed to register account.\n" + err.Error()))
		return
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
