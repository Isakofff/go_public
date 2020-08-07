package main

import (
    "fmt"
    "net/http"
)

func contentType(url string) (string, error) {
    resp, err := http.Get(url)
    // if it's not nill that means that some error has happened
    if err != nil {
        return "", err
    }

    // make sure we close the body
    defer resp.Body.Close()

    response := resp.Header.Get("Content-Type")
    // return error if Content-Type header not found
    if response == "" {
        return "", fmt.Errorf("can't find Content-Type header")
    }
    
    return response, nil
}

func main() {
	header, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(header)
	}
}
