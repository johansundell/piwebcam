package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme:   "http",
		Host:     "127.0.0.1:8080",
		Path:     "/",
		RawQuery: "action=stream",
	})
	http.ListenAndServe(":80", proxy)
}
