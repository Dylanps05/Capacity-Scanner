package sensor

import "net/http"
import "github.com/Dylanps05/Capacity-Scanner/internal/cstype"

type SensorAPILogic interface {
	AuthenticateSensor(t cstype.ScannerToken) (bool, error)
	ParseRequest(r *http.Request) (int, error)
	RecordPopulation(p int) error
}

type SensorAPIHandler interface {
	SensorAPILogic
	HandlePopulationUpdate(w http.ResponseWriter, r *http.Request)
}
