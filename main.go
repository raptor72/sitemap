package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/raptor72/glink"
)

// func bfsLinkCollector(linksArray []string, depth int, domain string) ([]string, int) {
// func bfsLinkCollector(linksMap map[string]string, depth int, domain string) ([]string, int) {
// 	u, err := url.Parse(domain)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	resp, err := http.Get(domain)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var startedDepth int
// 	currentLevelLinks := make(map[string]string)
// 	// var currentLevelLinks []string

// 	links, err := glink.Parse(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//     startedDepth += 1
// 	for _, link := range links {
// 		// currentLevelLinks = append(currentLevelLinks, link.Href)
//         currentLevelLinks[link.Href] = link.Href
// 	}

//     // здесь надо смерджить linksArray и currentLevelLinks
// 	// скорее всего потом это будет отдельной функцией
//     for key, value := range currentLevelLinks {
//     	value, ok := linksMap[key]
//     	if ok {
// 		    // клюбч уже есть в списке
// 		    delete(currentLevelLinks, key)
// 		} else {
// 		     // этого линка нет, добавляем его в результирующую карту
//     		linksMap[key] = value
// 	    }
// 	}
// }


func getLink(domain string, path string) (io.ReadCloser, error) {
	u, err := url.Parse(domain)
	if err != nil {
		log.Fatal(err)
        return nil, err
	}
	sub_u, err := url.Parse(path)
        if err != nil {
	    log.Fatal(err)
        return nil, err
	}
    // referer like /some_link
    if sub_u.Scheme == "" && sub_u.Host == "" {
	    fmt.Println("just path")
	    resp, err := http.Get(u.ResolveReference(sub_u).String())
	    if err != nil {
		    log.Fatal(err)
    	}
	    return resp.Body, nil
    } 
	// full referer like "http://domain.com/some_link"
	if sub_u.Scheme == u.Scheme && sub_u.Host == u.Host {
	    fmt.Println("full url")
	    fmt.Println(sub_u.String())
	    resp, err := http.Get(sub_u.String())
	    if err != nil {
        	log.Fatal(err)
	    }
	    return resp.Body, nil
    }
    return nil, nil
}


func main() {
    // domain := "http://example.com/"
	domain := "http://127.0.0.1:8080"

	// u, err := url.Parse(domain)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(u.Scheme)
	// fmt.Println(u.Host)
	// fmt.Println(u.Path)

	resp, err := http.Get(domain)
	if err != nil {
		log.Fatal(err)
	}
    // fmt.Println(resp)
    // fmt.Printf("%T\n", resp)
    // fmt.Printf("%T\n", resp.Body)

	links, err := glink.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
    fmt.Println(links)

	for _, lnk := range links {
        body, _ := getLink(domain, lnk.Href)
		links, err := glink.Parse(body)
		if err != nil {
			panic(err)
		}
		fmt.Println(links)
	}
}