package main

import (
	"api"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/duration", api.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("html/index.html")
	defer file.Close()
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), 500)
		return
	}
	page, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), 500)
		return
	}
	fmt.Fprint(w, string(page))

}
