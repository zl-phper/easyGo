package server

import "net/http"



type HandlerBaseOnMap struct {

	// key 应该method + url
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBaseOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request){

	key := h.key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok{
		handler(NewContext(writer,request))
	}else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("not found"))
	}

}

func (h *HandlerBaseOnMap ) key(method string,pattern string) string {
	 return method + "#" + pattern
}
