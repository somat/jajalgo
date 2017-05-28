package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://golang.org")
	fmt.Println(resp)
}
