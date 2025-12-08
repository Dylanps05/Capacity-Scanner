package mmw

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
	"github.com/Dylanps05/Capacity-Scanner/internal/cstype/ctxkey"
)

type SessionMiddlewareLogic interface {
	NewSession(ctx context.Context) string
	ValidateSession(ctx context.Context, s string) bool
	GetUserIDFromSession(ctx context.Context, s string) (cstype.UserID, error)
}

type SessionMiddleware struct {
	SessionMiddlewareLogic
	next http.Handler
}

func NewSessionMiddleware(next http.Handler, l SessionMiddlewareLogic) http.Handler {
	return &SessionMiddleware{
		SessionMiddlewareLogic: l,
		next:                   next,
	}
}

func (m *SessionMiddleware) createSession(ctx context.Context) *http.Cookie {
	new_session := m.NewSession(ctx)
	new_session_cookie := http.Cookie{
		Name:  "session",
		Value: new_session,

		Expires: time.Now().Add(time.Minute * 30),

		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	return &new_session_cookie
}

func (m *SessionMiddleware) ServeFreshSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session_cookie := m.createSession(ctx)
	r.AddCookie(session_cookie)
	http.SetCookie(w, session_cookie)
}

// TODO: CSRF vulnerable, probably.
func (m *SessionMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer m.next.ServeHTTP(w, r)

	// Extract session cookie
	session_cookie := r.CookiesNamed("session")

	// Has session?
	no_session := len(session_cookie) == 0
	if no_session {
		m.ServeFreshSession(w, r)
		return
	}

	// Session valid?
	ctx := r.Context()
	session := session_cookie[0].Value
	invalid_session := !m.ValidateSession(ctx, session)
	if invalid_session {
		m.ServeFreshSession(w, r)
		return
	}

	// Session OK
	auth_ctx := context.WithValue(ctx, ctxkey.Session, session)
	uid, err := m.GetUserIDFromSession(ctx, session)
	if err != nil {
		log.Fatal("SessionHandler: Failed to get UID from session:\n" + err.Error())
	}
	auth_ctx = context.WithValue(auth_ctx, ctxkey.UID, uid)
	r = r.WithContext(auth_ctx)
}
