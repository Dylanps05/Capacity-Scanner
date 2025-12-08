package auth

import (
	"context"
	"net/http"

	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
)

type AuthModule interface {
	NewSession(ctx context.Context) string
	CreateLogin(ctx context.Context, u string, p string) error
	GetSessionFromCookie(ctx context.Context, c *http.Cookie) (string, error)
	ValidateSession(ctx context.Context, s string) bool
	IsSessionAuthed(ctx context.Context, s string) bool
	SetSessionAuthByLogin(ctx context.Context, s string, u string, p string) error
	DeauthSession(ctx context.Context, s string) error
	GetUsernameFromSession(ctx context.Context, s string) (string, error)
	GetUserIDFromSession(ctx context.Context, s string) (cstype.UserID, error)
}
