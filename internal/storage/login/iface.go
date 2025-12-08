package login

import (
	"context"

	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
)

type LoginStorage interface {
	StoreLogin(ctx context.Context, login cstype.Login) (cstype.Login, error)
	GetLoginByUsername(ctx context.Context, username string) (cstype.Login, error)
	GetLoginByUUID(ctx context.Context, uuid cstype.UserID) (cstype.Login, error)
	SetLoginByUUID(ctx context.Context, login cstype.Login) (cstype.Login, error)
	DeleteLoginByUUID(ctx context.Context, uuid cstype.UserID) (cstype.Login, error)
}
