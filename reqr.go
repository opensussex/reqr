package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
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
		os.Exit(1)
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "\t")
	if err != nil {
		fmt.Println("Ooops! somethings gone wrong")
		os.Exit(1)
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

	//get the local IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Ooops! somethings gone wrong")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("Send Requests to: http://" + ipnet.IP.String() + ":8666")
				//os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
	fmt.Println("Ctrl-C exit!")
	http.HandleFunc("/", req)
	err = http.ListenAndServe(":8666", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
