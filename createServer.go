package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {

	case "GET":
		content, err := os.ReadFile("html/index.html")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		w.Write([]byte(content))

	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Recieved a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}
func main() {
	content2, err := os.ReadFile("properties/port.txt")
	if err != nil {
		log.Fatal(err)
	}
	 http.HandleFunc("/", helloWorld)
	 http.ListenAndServe(string(content2), nil)
}
