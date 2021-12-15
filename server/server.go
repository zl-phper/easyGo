package server

import "net/http"

type Server interface {

	// method POST GET PUT DELET

	Route(method string, path string, handleFunc func(ctx *Context))

	Start(address string) error

	Md5() string
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handleFunc func(ctx *Context)) {

	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		ctx :=  NewContext(writer,request)
		handleFunc(ctx)
	})
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}
func (s *sdkHttpServer) Md5() string {
	return GetMD5("http://www.baidu.com")
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{Name: name}
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context{
	return &Context{
		R: request,
		W: writer,
	}
}

