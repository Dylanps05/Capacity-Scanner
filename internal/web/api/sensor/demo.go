package sensor

import (
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
