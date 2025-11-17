package capacity

import (
	"context"
	"git.11spades.net/CivilMatch/civilmatch/internal/cstype"
)

type CapacityStorage interface {
	StoreCapacity(ctx context.Context, d int) error
	GetCapacity(ctx context.Context) (int, error)
}
