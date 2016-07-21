package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akfork/app"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method, "url: ", r.URL)
	fmt.Fprintln(w, "hello world")
}

func main() {
	port := flag.Int("port", 8888, "Listen on port")
	flag.Parse()

	a := app.New()

	a.Get("/", handleRoot)

	a.Listen(":" + strconv.Itoa(*port))
}
