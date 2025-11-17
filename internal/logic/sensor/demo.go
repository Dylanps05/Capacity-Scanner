package sensor

import (
	"context"
	"fmt"
	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/capacity"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/sensor"
	"net/http"
	"reflect"
	"strconv"
)

type DefaultSensorModule struct {
	sensor.SensorAuthStorage
	capacity.CapacityStorage
}

func NewDefaultSensorModule(s sensor.SensorAuthStorage, c capacity.CapacityStorage) SensorModule {
	l := &DefaultSensorModule{
		SensorAuthStorage: s,
		CapacityStorage:   c,
	}

	return l
}

func (l *DefaultSensorModule) AuthenticateSensor(r *http.Request) error {
	ctx := r.Context()
	token := cstype.ScannerToken(r.Header.Get("Authentication"))
	ok, err := l.SensorAuthStorage.HaveToken(ctx, token)
	if err != nil {
		return fmt.Errorf("%T: Failed to query token registry\n%w", reflect.TypeOf(l), err)
	}
	if !ok {
		return fmt.Errorf("%T: bad token", reflect.TypeOf(l))
	}
	return nil
}

func (l *DefaultSensorModule) ParseRequest(r *http.Request) (int, error) {
	pop_delta_s := r.FormValue("delta")
	pop_delta, err := strconv.ParseInt(pop_delta_s, 10, 32)
	if err != nil {
		return int(pop_delta), fmt.Errorf("%T: Failed to parse pop delta\n%w", reflect.TypeOf(l), err)
	}
	return int(pop_delta), nil
}

func (l *DefaultSensorModule) RecordPopulation(ctx context.Context, p int) error {
	old_pop, err := l.CapacityStorage.GetCapacity(ctx)
	if err != nil {
		return fmt.Errorf("%T: Failed to query for old pop\n%w", reflect.TypeOf(l), err)
	}

	new_pop := old_pop + p
	err = l.StoreCapacity(ctx, new_pop)
	if err != nil {
		return fmt.Errorf("%T: Failed to store new pop\n%w", reflect.TypeOf(l), err)
	}

	return nil
}
