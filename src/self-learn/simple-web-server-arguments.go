// MIT License
// Copyright (c) 2017 @harkamals

// Simple web server-with-arguments
// Usage: go run simple-web-server-arguments.go -port 5051 -path /mypath

package self_learn

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	var dir string
	port := flag.String("port", "3000", "port to serve HTTP on")
	path := flag.String("path", "", "path to serve")
	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}
