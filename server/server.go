package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"	
)

func main() {
	// command line flags
	port := flag.Int("port", 80, "port to serve on")  // default is 80
	dir := flag.String("directory", "web/", "directory for web files") // default is /web
	flag.Parse()

	// handle requests serving a file of the same name
	fs := http.Dir(*dir)
	handler := http.FileServer(fs)
	http.Handle("/", handler)
	http.HandleFunc("/go", handlerGo)

	log.Printf("running on port %d\n", port)

	addr := fmt.Sprintf("127.0.0.1:%d", *port)

	// from this block, the program runs forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}

func handlerGo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("It's Alive, alive!")
}
