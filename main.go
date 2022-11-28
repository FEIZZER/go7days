package main

import (
	"fei"
	"fmt"
	"net/http"
)

func main() {
	serve := fei.New()
	serve.AddHandle("GET", "/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
		})

	err := serve.Run(":8080")
	fmt.Println(err)
}
