package population

import (
	"html/template"
	"net/http"

	"github.com/Dylanps05/Capacity-Scanner/internal/cstype/ctxkey"
)

type DemoPopulationHandler struct {
	PopulationHandlerLogic
	landingPage *template.Template
}

func NewDemoPopulationHandler(m *http.ServeMux, l PopulationHandlerLogic) PopulationHandler {
	h := &DemoPopulationHandler{
		PopulationHandlerLogic: l,
		landingPage:            template.New("./web/template/landing.html"),
	}

	m.HandleFunc("/", h.HandleLandingPageRequest)

	return h
}

func (h *DemoPopulationHandler) HandleLandingPageRequest(w http.ResponseWriter, r *http.Request) {
	rsp := http.Response{
		StatusCode: http.StatusOK,
	}
	defer rsp.Write(w)

	ctx := r.Context()
	pop, err := h.PopulationHandlerLogic.CurrentPopulation(ctx)
	if err != nil {
		w.Write([]byte(err.Error()))
		rsp.StatusCode = http.StatusInternalServerError
		return
	}

	page_data := struct {
		UID        any
		Population int
	}{
		UID:        ctx.Value(ctxkey.UID),
		Population: pop,
	}
	h.landingPage.Execute(w, page_data)
}
