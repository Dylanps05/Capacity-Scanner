package mmw

import (
	"net/http"

	"github.com/Dylanps05/Capacity-Scanner/internal/logic"
)

type MuxMiddleware interface {
	http.Handler
	logic.Controller
	GetMux() http.Handler
}
