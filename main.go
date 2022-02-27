package main

import (
    // "os"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"github.com/raptor72/glink"
	// "encoding/xml"
)


// type SiteMap struct {
// 	XMLName   xml.Name `xml:"person"`
//     Url       string   `xml:"url>loc"`
// }


func bfsLinkCollector(currentlinksMap map[string]bool, domain string) map[string]bool {
	resMap := make(map[string]bool)
    for key := range currentlinksMap {
		resMap[key] = true
	}
    for i:= len(currentlinksMap); i > 0; {
		middleMap := make(map[string]bool)
		for link := range currentlinksMap {
			newMap, _ := getLink(domain, link)
			for j := range newMap {
				if !resMap[j] {
                    resMap[j] = true
					middleMap[j] = true
				}
			}
			currentlinksMap = middleMap
		}
        i = len(currentlinksMap)
	}
    return resMap
}

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
	resMap := bfsLinkCollector(firstMap, domain)
    fmt.Println(resMap)
}