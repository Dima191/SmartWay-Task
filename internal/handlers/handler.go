package handlers

import "net/http"

type Handler interface {
	Register(router http.Handler) error
}
