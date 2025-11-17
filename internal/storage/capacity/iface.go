package capacity

import (
	"context"
)

type CapacityStorage interface {
	StoreCapacity(ctx context.Context, d int) error
	GetCapacity(ctx context.Context) (int, error)
}
