package mmw

import (
	"net/http"

	"github.com/Dylanps05/Capacity-Scanner/internal/logic"
)

type DefaultMuxMiddleware struct {
	http.Handler
	logic.Controller
}

func NewDefaultMuxMiddleware(h http.Handler, c logic.Controller) MuxMiddleware {
	m := &DefaultMuxMiddleware{
		Handler:    h,
		Controller: c,
	}
	m.wrapMux()

	return m
}

func (m *DefaultMuxMiddleware) wrapMux() {
	h := m.Handler

	h = NewSessionMiddleware(h, m)

	m.Handler = h
}

func (m *DefaultMuxMiddleware) GetMux() http.Handler {
	return m.Handler
}
