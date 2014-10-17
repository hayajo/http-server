package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		address = flag.String("a", "0.0.0.0", "address")
		port    = flag.Int("p", 8080, "port")
	)
	flag.Parse()
	listen := fmt.Sprintf("%s:%d", *address, *port)
	root := flag.Arg(0)
	log.Printf("Listening at http://%s", listen)
	log.Fatal(http.ListenAndServe(listen, http.FileServer(http.Dir(root))))
}
