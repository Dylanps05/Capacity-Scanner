package web

import (
	"github.com/Dylanps05/Capacity-Scanner/internal/logic"
	"net/http"
)

type DefaultHandler struct {
	mux http.Handler
}

func (h *DefaultHandler) buildMux(c logic.Controller) {
	mux := http.NewServeMux()

	static_fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", static_fs))

	h.mux = mux
}

func (h *DefaultHandler) wrapMux() {
	m := h.mux

	h.mux = m
}

func NewDefaultHandler(c logic.Controller) Handler {
	h := &DefaultHandler{}
	h.buildMux(c)

	return h
}

func (h *DefaultHandler) Start(addr string) {
	http.ListenAndServe(addr, h.mux)
}
