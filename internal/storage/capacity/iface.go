package capacity

import (
	"context"
	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
)

type CapacityStorage interface {
	StoreCapacity(ctx context.Context, d int) error
	GetCapacity(ctx context.Context) (int, error)
}
