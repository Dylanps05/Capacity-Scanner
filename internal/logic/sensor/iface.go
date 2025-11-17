package sensor

import (
	"context"
	"net/http"
)

type SensorModule interface {
	AuthenticateSensor(r *http.Request) error
	ParseRequest(r *http.Request) (int, error)
	RecordPopulation(ctx context.Context, p int) error
}
