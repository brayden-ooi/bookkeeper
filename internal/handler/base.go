package handler

import "net/http"

func Base(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong!"))
}
