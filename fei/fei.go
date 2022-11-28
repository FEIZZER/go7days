package fei

import (
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Serve struct {
	router map[string]HandleFunc
}

func New() *Serve {
	return &Serve{router: map[string]HandleFunc{}}
}

func (server *Serve) AddHandle(method string, pattern string, handle HandleFunc) {
	server.router[method+"-"+pattern] = handle
}

func (serve *Serve) Run(addr string) error {
	return http.ListenAndServe(addr, serve)
}

func (serve *Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.RequestURI
	if handleFunc, ok := serve.router[key]; ok {
		handleFunc(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}
