package main

import (
	"io"
	"net/http"
)

type AntinioHandler int

func (h AntinioHandler) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	io.WriteString(res, req.RequestURI)
}

func main() {
	var h AntinioHandler

	http.ListenAndServe(":9000", h)
}