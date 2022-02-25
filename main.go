package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/raptor72/glink"
)


// func bfsLinkCollector(currentlinksMap map[string]bool, alllinksMap map[string]bool, depth int, domain string) (map[string]bool, int) {
// // func bfsLinkCollector(currentlinksMap map[string]bool, alllinksMap map[string]bool, depth int, domain string) {

//     nextLevelmap := make(map[string]bool)

// 	for key := range currentlinksMap {   // map[/denver:true /new-york:true]
// 		if alllinksMap[key] {
// 			fmt.Printf("%s lik already in result map\n", key)
// 		} else {
// 			newMap, err := getLink(domain, key)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			// fmt.Println(newMap) // map[/debate:true /home:true] // map[/debate:true /home:true]
// 			alllinksMap[key] = true
// 			for subKey := range newMap {
// 				nextLevelmap[subKey] = true
// 				if alllinksMap[subKey] {
// 					continue
// 				} else {
// 					alllinksMap[subKey] = true
// 				}
					
// 			}
            
// 		}

// 	}
//     fmt.Println("nextLevelmap", nextLevelmap)
// 	fmt.Println("allinksMap", alllinksMap)
// 	depth  += 1
//     for len(nextLevelmap) > 0 {
// 		bfsLinkCollector(currentlinksMap, alllinksMap, depth, domain)
// 	} 

//     return nextLevelmap, depth
// }


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
    // fmt.Println(v)
	return v, nil
}


func main() {
    // domain := "http://example.com/"
	domain := "http://127.0.0.1:8080"
    firstMap, _ := getLink(domain, "/") // map[/denver:true /new-york:true]
    // fmt.Println(firstMap)

	resMap := make(map[string]bool)
    for key := range firstMap {
		resMap[key] = true
	}


    for i:= len(firstMap); i > 0; {
		middleMap := make(map[string]bool)
		for link := range firstMap {
            // fmt.Println(link)
			newMap, _ := getLink(domain, link)
            // fmt.Println(newMap)
			for j := range newMap {
				if !resMap[j] {
                    resMap[j] = true
					middleMap[j] = true
				}
			}
			firstMap = middleMap
		}
        i = len(firstMap)
	}
    fmt.Println(resMap)
}