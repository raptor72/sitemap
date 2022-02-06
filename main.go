package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/raptor72/glink"
)


// func bfsLinkCollector(linksMap map[string]bool, depth int, domain string) (map[string]bool, int) {

// 	currentLevelLinks := make(map[string]bool)
// }
	// 	// var currentLevelLinks []string

// 	// links, err := glink.Parse(resp.Body)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
//     // startedDepth += 1
// 	// for _, link := range linksMap {
// 		// currentLevelLinks = append(currentLevelLinks, link.Href)
//         // currentLevelLinks[link] = true
// 	for key := range linksMap {
// 		body, _ := getLink(domain, key)
// 		links, err := glink.Parse(body)
// 		if err != nil {
// 			panic(err)
// 		}
// 		for key2 := range links {
// 			exists := currentLevelLinks[key2]
// 			if exists {
// 				// клюбч уже есть в списке
// 				delete(currentLevelLinks, key2)
// 			} else {
// 				 // этого линка нет, добавляем его в результирующую карту
// 				linksMap[key] = true
// 			}
// 		}

// 	}



func getLink(domain string, path string) (map[string]bool, error) {
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
	v := make(map[string]bool)
	// referer like /some_link
	if sub_u.Scheme == "" && sub_u.Host == "" {
	    // fmt.Println("just path")
	    resp, err := http.Get(u.ResolveReference(sub_u).String())
	    if err != nil {
		    return nil, err
    	}
		links, err := glink.Parse(resp.Body)
		if err != nil {
			return nil, err
		}
		for _, value := range links {
			v[value.Href] = true
		}
	} 
	// full referer like "http://domain.com/some_link"
	if sub_u.Scheme == u.Scheme && sub_u.Host == u.Host {
	    // fmt.Println("full url")
	    resp, err := http.Get(sub_u.String())
	    if err != nil {
        	return nil, err
	    }
		links, err := glink.Parse(resp.Body)
		if err != nil {
			return nil, err
		}
		for _, value := range links {
			v[value.Href] = true
		}
	}
    return v, nil
}


func main() {
    // domain := "http://example.com/"
	domain := "http://127.0.0.1:8080"
    firstMap, _ := getLink(domain, "/")
    fmt.Println(firstMap)
}