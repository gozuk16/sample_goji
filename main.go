package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	result := fmt.Sprintf("Hello, %s!", c.URLParams["name"])
	encoder.Encode(result)
}

func main() {
	goji.Get("/hello/:name", hello)
	goji.Get("/assets/*", http.FileServer(http.Dir(".")))
	goji.Serve()
}
