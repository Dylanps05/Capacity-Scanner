package sensor

import (
	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
	"net/http"
)

type DemoSensorAPIHandler struct {
	SensorAPILogic
}

func NewDemoSensorAPIHandler(m *http.ServeMux, l SensorAPILogic) SensorAPIHandler {
	h := &DemoSensorAPIHandler{
		SensorAPILogic: l,
	}

	m.HandleFunc("/api/sensor/update", h.HandlePopulationUpdate)

	return h
}

func (h *DemoSensorAPIHandler) HandlePopulationUpdate(w http.ResponseWriter, r *http.Request) {
	token := cstype.ScannerToken(r.Header.Get("Authentication"))
	err := h.SensorAPILogic.AuthenticateSensor(r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	pop, err := h.SensorAPILogic.ParseRequest(r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	h.SensorAPILogic.RecordPopulation(r.Context(), pop)
}
