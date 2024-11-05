// Package handler contains handler implementation for HTTP server.
package handler

import "net/http"

// HelloHandler is a handler function for "GET /hello" endpoint.
var HelloHandler = func(w http.ResponseWriter, _ *http.Request, _ map[string]string) {
	_, _ = w.Write([]byte("Hello, World!\n"))
}
