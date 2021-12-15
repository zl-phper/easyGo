package server

import "net/http"

type Handler interface {
	http.Handler
	Route(method string, path string, handleFunc func(ctx *Context))
}

type HandlerBaseOnMap struct {

	// key 应该method + url
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBaseOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	key := h.key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("not found"))
	}

}

func (h *HandlerBaseOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

func (s *HandlerBaseOnMap) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	key := s.key(method, pattern)
	s.handlers[key] = handleFunc
}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBaseOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
