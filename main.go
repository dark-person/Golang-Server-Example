package main

import (
	"AnimeImageAnalyser-v2/server"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world!")

	r := server.New()

	http.ListenAndServe(":3050", r)
}
