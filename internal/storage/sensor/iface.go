package sensor

import (
	"context"
	"git.11spades.net/CivilMatch/civilmatch/internal/cstype"
)

type SensorAuthStorage interface {
	StoreToken(ctx context.Context) error
}
