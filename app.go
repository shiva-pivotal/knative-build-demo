package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlePost(rw http.ResponseWriter, req *http.Request) {

	// body, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	panic(err)
	// }
	//log.Println(string(body))
	fmt.Fprintf(rw, "%s", "Hello World")
}

func main() {
	log.Print("Starting server...")
	http.HandleFunc("/", handlePost)
	log.Fatal(http.ListenAndServe(":80", nil))
}
