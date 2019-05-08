package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// how to make HTTP request
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			// exit the process
			os.Exit(1)
		}
		// resp is a struct, the body field is a readable stream
		b, err := ioutil.ReadAll(resp.Body)
		// close body to avoid leaking resources
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// write to standard output
		fmt.Printf("%s", b)
	}
}
