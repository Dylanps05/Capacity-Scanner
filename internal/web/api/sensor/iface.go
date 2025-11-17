package sensor

import "net/http"
import "context"

type SensorAPILogic interface {
	AuthenticateSensor(r *http.Request) error
	ParseRequest(r *http.Request) (int, error)
	RecordPopulation(ctx context.Context, p int) error
}

type SensorAPIHandler interface {
	SensorAPILogic
	HandlePopulationUpdate(w http.ResponseWriter, r *http.Request)
}
