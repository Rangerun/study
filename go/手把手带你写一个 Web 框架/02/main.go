package main

import (
	//"coredemo/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		//Handler: framework.NewCore(),
		Addr:    "localhost:8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	});
	server.ListenAndServe()
}