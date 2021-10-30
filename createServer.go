package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"tawesoft.co.uk/go/dialog"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		content, err := os.ReadFile("html/index.html")
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(content))
	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Received a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}

func main() {

	content, err := os.ReadFile("properties/port.txt")
	if err != nil {
		log.Fatal()
	}
	var converted string
	converted = string(content)
	if converted == "" {
		dialog.Alert("Port unspecified.")
		log.Fatal("Port for server is not specified, exiting.")
	}
	http.HandleFunc("/", helloWorld)

	http.ListenAndServe(":"+converted, nil)
}
