package main

import (
	"fmt"
	"net/http"
	"os"
)

func reportError(err error) {
	fmt.Println("Błąd")
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	source := os.Args[1]
	target := os.Args[2]

	resp, err := http.Get(source)

	if err != nil {
		reportError(err)
	}

	contentType := resp.Header["Content-Type"][0]

	_, err = http.Post(target, contentType, resp.Body)

	resp.Body.Close()

	if err != nil {
		reportError(err)
	}

	fmt.Println("Zakończono pomyślnie")
	os.Exit(0)
}
