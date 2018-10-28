package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm version 2!"))
	})
	http.ListenAndServe(":8001", nil)
}
