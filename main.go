package main

import (
	"fmt"
	"net/http"
	"servers/server"
)

func main() {
	fmt.Println("Hello world!")

	r := server.New()

	http.ListenAndServe(":3050", r)
}
