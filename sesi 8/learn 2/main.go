package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "https://hikiotect.com/articles?name=Fajri&age=24"

	var u, e = url.Parse(urlString)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]

	fmt.Printf("name: %s, age: %s\n", name, age)
}
