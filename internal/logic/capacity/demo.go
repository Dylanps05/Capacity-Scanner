package capacity

import (
	"context"
	"fmt"

	"github.com/Dylanps05/Capacity-Scanner/internal/storage/capacity"
	"reflect"
)

type DefaultCapacityModule struct {
	capacity.CapacityStorage
}

func NewDefaultCapacityModule(s capacity.CapacityStorage) CapacityModule {
	l := &DefaultCapacityModule{
		CapacityStorage: s,
	}

	return l
}

func (l *DefaultCapacityModule) CurrentPopulation(ctx context.Context) (int, error) {
	pop, err := l.CapacityStorage.GetCapacity(ctx)
	if err != nil {
		return pop, fmt.Errorf("%T: Failed to query for current population\n%w", reflect.TypeOf(l), err)
	}
	return pop, err
}
