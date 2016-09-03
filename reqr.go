package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func req(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Ooops! somethings gone wrong")
		panic(err)
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "\t")
	if err != nil {
		fmt.Println("Ooops! somethings gone wrong")
		panic(err)
	}
	fmt.Println("Body: ")
	fmt.Println(string(prettyJSON.Bytes()))
	fmt.Println("Params: ")
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}

func main() {
	fmt.Println("Send Requests to: http://localhost:8666")
	fmt.Println("Ctrl-C exit!")
	http.HandleFunc("/", req)
	err := http.ListenAndServe(":8666", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
