package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	 
	// handle requests serving a file of the same name
	fs := http.Dir("html/")
	handler := http.FileServer(fs)
	http.Handle("/", handler)
	http.HandleFunc("/go", handlerGo)

	log.Print("server running on port 80")

	addr := fmt.Sprintf("127.0.0.1:80")

	// from this block, the program runs foreve	r
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}

func handlerGo(w http.ResponseWriter, r *http.Request) {
	log.Print("It's Alive!")
}
