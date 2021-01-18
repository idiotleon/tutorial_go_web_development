package main_04

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot")))
}
