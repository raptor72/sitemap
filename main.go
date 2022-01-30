package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"github.com/raptor72/glink"
)


func main() {
    // domain := "http://example.com/"
	domain := "http://127.0.0.1:8080"

	u, err := url.Parse(domain)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)

	resp, err := http.Get(domain)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println(resp)
    fmt.Printf("%T\n", resp)
    fmt.Printf("%T\n", resp.Body)

	links, err := glink.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
    fmt.Println(links)

	for _, lnk := range links {


		sub_u, err := url.Parse(lnk.Href)
		if err != nil {
			log.Fatal(err)
		}
        // print(sub_u.Path)
        // print(sub_u.Scheme)
		if sub_u.Scheme == sub_u.Host == nil {
			resp, err := http.Get(u.Scheme + u.Host + sub_u.Path)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(resp)
		}

		if sub_u.Host == u.Host && sub_u.Scheme == u.Scheme  {
			resp, err := http.Get(sub_u.String())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(resp)
		}
	}
}