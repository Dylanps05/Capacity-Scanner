package population

import "net/http"

type PopulationHandlerLogic interface {
	CurrentPopulation() (int, error)
}

type PopulationHandler interface {
	PopulationHandlerLogic
	HandleLandingPageRequest(w http.ResponseWriter, r *http.Request)
}
