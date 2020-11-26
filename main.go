package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "URL---PATH = %q \n", req.URL.Path)
	})
	r.Run(":9999")
}
