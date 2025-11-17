package sensor

import "net/http"

type SensorAPILogic interface {
	AuthenticateSensor(r *http.Request) (bool, error)
	ParseRequest(r *http.Request) (int, error)
	RecordPopulation(p int) error
}

type SensorAPIHandler interface {
	SensorAPILogic
	HandlePopulationUpdate(w http.ResponseWriter, r *http.Request)
}
