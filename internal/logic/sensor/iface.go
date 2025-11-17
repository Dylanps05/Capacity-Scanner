package sensor

import (
	"net/http"
)

type SensorModule interface {
	AuthenticateSensor(r *http.Request) (bool, error)
	ParseRequest(r *http.Request) (int, error)
	RecordPopulation(p int) error
}
