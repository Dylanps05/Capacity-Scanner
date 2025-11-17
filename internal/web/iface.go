package web

import ()

type Handler interface {
	Start(addr string)
}
