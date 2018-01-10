package main

import (
	"net/http"
	"fmt"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request from", r.RemoteAddr)
	w.Write([]byte("indexPage"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8888", nil)
}
