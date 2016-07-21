package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jeffail/gabs"
	"github.com/akfork/app"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method, "url: ", r.URL)

	jsonObj := gabs.New()

	jsonObj.Set(0, "code")
	jsonObj.Set("hello world", "data")

	fmt.Fprintf(w, jsonObj.String())
}

func main() {
	port := flag.Int("port", 8888, "Listen on port")
	flag.Parse()

	a := app.New()

	a.Get("/", handleRoot)

	a.Listen(":" + strconv.Itoa(*port))
}
