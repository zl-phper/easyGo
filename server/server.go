package server

import (
	"net/http"
)

type Server interface {

	// method POST GET PUT DELET
	Route(method string, path string, handleFunc func(ctx *Context))

	Start(address string) error

	Md5() string
}

type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {

	s.handler.Route(method, pattern, handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		c := NewContext(writer, request)
		s.root(c)
	})

	return http.ListenAndServe(address, nil)
}

func (s *sdkHttpServer) Md5() string {
	return GetMD5("http://www.baidu.com")
}

func NewHttpServer(name string, builders ...FilterBuilder) Server {

	handler := NewHandlerBasedOnMap()
	var root Filter = handler.ServerHTTP


	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{
		Name:    name,
		handler: NewHandlerBasedOnMap(),
		root:    root,
	}
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		R: request,
		W: writer,
	}
}
