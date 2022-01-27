package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raptor72/glink"
)


func main() {
	resp, err := http.Get("http://example.com/")
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

}