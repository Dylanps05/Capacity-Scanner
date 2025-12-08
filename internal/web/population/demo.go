package population

import (
	"html/template"
	"log"
	"net/http"
)

type DemoPopulationHandler struct {
	PopulationHandlerLogic
	landingPage *template.Template
}

func NewDemoPopulationHandler(m *http.ServeMux, l PopulationHandlerLogic) PopulationHandler {
	h := &DemoPopulationHandler{
		PopulationHandlerLogic: l,
	}

	var err error
	h.landingPage, err = template.ParseFiles("./web/template/landing.html")
	if err != nil {
		log.Fatal(err)
	}

	m.HandleFunc("/", h.HandleLandingPageRequest)

	return h
}

func (h *DemoPopulationHandler) HandleLandingPageRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pop, err := h.PopulationHandlerLogic.CurrentPopulation(ctx)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	page_data := struct {
		Population int
	}{
		Population: pop,
	}
	h.landingPage.Execute(w, page_data)
}
