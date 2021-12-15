package server

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(req interface{}) error {

	r := c.R
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, req)

	if err != nil {
		return err
	}

	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)

	respJon, err := json.Marshal(resp)

	if err != nil {
		return err
	}

	_, err = c.W.Write(respJon)

	return err
}


func (c *Context) WriteJsonSuc(resp interface{}) error {
 	return c.WriteJson(http.StatusOK,resp)
}


func (c *Context) BadRequest(resp interface{}) error {
	return c.WriteJson(http.StatusNotFound,resp)
}