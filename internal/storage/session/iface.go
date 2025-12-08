package session

import (
	"context"

	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
)

type SessionStorage interface {
	StoreSession(ctx context.Context, session_id string) error
	GetSessionUUID(ctx context.Context, token string) (cstype.UserID, error)
	SetSessionUUID(ctx context.Context, token string, uuid cstype.UserID) error
	DeleteSession(ctx context.Context, token string) error
}
