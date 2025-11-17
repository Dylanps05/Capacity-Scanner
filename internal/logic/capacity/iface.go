package capacity

import (
	"context"
)

type CapacityModule interface {
	CurrentPopulation(ctx context.Context) (int, error)
}
