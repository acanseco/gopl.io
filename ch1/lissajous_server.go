// Lissajous server that prints a lissajous figure to the browser.
package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles, ok := r.URL.Query()["cycles"]
		if !ok || len(cycles[0]) < 1 {
			log.Println("URL Param 'cycles' is missing")
		}
		c, _ := strconv.Atoi(cycles[0])
		lissajous(w, c)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
