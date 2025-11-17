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
	no_token := token == ""
	if no_token {
		w.Write([]byte("Bad auth: no token"))
		return
	}
	token_ok, err := h.SensorAPILogic.AuthenticateSensor(token)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if !token_ok {
		w.Write([]byte("Bad auth: token not registered"))
		return
	}

	pop, err := h.SensorAPILogic.ParseRequest(r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	h.SensorAPILogic.RecordPopulation(pop)
}
