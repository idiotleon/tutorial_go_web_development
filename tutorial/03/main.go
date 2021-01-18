package main_03

import "net/http"

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type aboutHandler struct{}

func (a *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func main() {
	hh := helloHandler{}
	ah := aboutHandler{}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}

	http.Handle("/hello", &hh)
	http.Handle("/about", &ah)

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home!"))
	})
	http.HandleFunc("/welcome", welcome)
	// http.Handle("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe()
}
