package main

import (
	"net/http"
	"fmt"
	"log"
	"io"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Println("request from", r.RemoteAddr)
		
		var user string
		cookie, err := r.Cookie("user")
		if err == nil {
			user = cookie.Value 
		} else {
			user = ""
		}

		if user != "" {
			io.WriteString(w, "Hello " + user)
		} else {
			w.Write([]byte("indexPage"))
		}
		
		return
	}

	NotFoundHandler(w, r)
}

func main() {
	http.Handle("/assets/", http.FileServer(http.Dir("assets")))
	
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
