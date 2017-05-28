package main

import "fmt"
import "net/url"

func main() {
	s := "mongodb://user:pass@localhost:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)

}
