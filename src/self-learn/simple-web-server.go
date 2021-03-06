// MIT License
// Copyright (c) 2017 @harkamals

// Simple web server

package self_learn

import (
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
}
