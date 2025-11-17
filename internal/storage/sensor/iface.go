package sensor

import (
	"context"
	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
)

type SensorAuthStorage interface {
	StoreToken(ctx context.Context, t cstype.ScannerToken) error
	HaveToken(ctx context.Context, t cstype.ScannerToken) (bool, error)
	RevokeToken(ctx context.Context, t cstype.ScannerToken) error
}
