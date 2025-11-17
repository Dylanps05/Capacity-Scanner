package population

import (
	"context"
	"net/http"
)

type PopulationHandlerLogic interface {
	CurrentPopulation(ctx context.Context) (int, error)
}

type PopulationHandler interface {
	PopulationHandlerLogic
	HandleLandingPageRequest(w http.ResponseWriter, r *http.Request)
}
