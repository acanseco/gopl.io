// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		dst := os.Stdout
		b, err := io.Copy(dst, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
