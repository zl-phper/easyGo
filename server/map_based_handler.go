package server

import (
	"net/http"
)

type Handler interface {
	ServerHTTP(c *Context)
	Route(method string, path string, handleFunc handleFunc)
}

type HandlerBaseOnMap struct {

	// key 应该method + url
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBaseOnMap) ServerHTTP(c *Context) {

	key := h.key(c.R.Method, c.R.URL.Path)

	if handler, ok := h.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("not found"))
	}

}

func (h *HandlerBaseOnMap) key(method string, pattern string) string {
	return method  + "#" + pattern
}

func (s *HandlerBaseOnMap) Route(method string, pattern string, handleFunc handleFunc) {
	key := s.key(method, pattern)
	 
	s.handlers[key] = handleFunc
}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBaseOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
