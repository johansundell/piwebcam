package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	/*cmd := exec.Command("mjpg_streamer", "-o 'output_http.so -w /home/johan/src/mjpg-streamer/mjpg-streamer-experimental/www'", "-i 'input_raspicam.so  -ex -hf -vf night'")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}*/

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme:   "http",
		Host:     "127.0.0.1:8080",
		Path:     "/",
		RawQuery: "action=stream",
	})
	log.Fatal(http.ListenAndServe(":80", proxy))
}
