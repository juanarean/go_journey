package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)

	// n, readErr := resp.Body.Read(bs)

	// fmt.Println(string(bs))
	// fmt.Println(n)
	// fmt.Println(readErr)
	io.Copy(os.Stdout, resp.Body)
}
