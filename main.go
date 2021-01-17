package main

import "net/http"

type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	mh := myHandler{}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &mh,
	}

	server.ListenAndServe()
}
