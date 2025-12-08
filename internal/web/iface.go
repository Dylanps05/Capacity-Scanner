package web

import "net/http"

type Handler interface {
	GetMux() http.Handler
}
