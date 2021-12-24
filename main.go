package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	server2 "selfgo/server"
	"time"
)

type SignUpReq struct {
	Email             string `json："email"`
	PassWord          string `json："password"`
	ConfirmedPassword string `json："confirmed_password"`
}

type commonResponse struct {
	Data interface{}
	Code int
}

func handler(ctx *server2.Context) {
	fmt.Fprintf(ctx.W, "hi there, i havle %s!", ctx.R.URL.Path[1:])
}

func queryParams(ctx *server2.Context) {

	data, _ := json.Marshal(ctx.R.URL)

	fmt.Fprint(ctx.W, string(data))
}

func signUp(ctx *server2.Context) {
	req := &SignUpReq{}

	err := ctx.ReadJson(req)

	if err != nil {
		fmt.Printf("err is %v", err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}

	err = ctx.WriteJsonSuc(resp)

	if err != nil {
		fmt.Printf("err is %v", err)
		return
	}

}

func GoRoutine() {

	go func() {
		time.Sleep(10 * time.Second)
	}()
	fmt.Println("i am here")
}

func main() {

	server := server2.NewHttpServer("test-server", server2.MetricsFilterBuilder,server2.MetricsFilterBuilder2)

	server.Route(http.MethodGet, "/signup", signUp)

 	server.Start(":8080")
}



