package main

import (
	"fmt"
	"net/http"
	"strings"
)

func req(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}

func main() {
	fmt.Println("Send Requests to: http://localhost:8666")
	http.HandleFunc("/", req)
	err := http.ListenAndServe(":8666", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
