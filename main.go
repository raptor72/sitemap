package main

import (
	"fmt"
    "flag"
	"log"
	"net/http"
	"net/url"
	"github.com/raptor72/glink"
	"encoding/xml"
)

type Url struct {
	XMLName   xml.Name `xml:"url"`
	Loc       string   `xml:"loc"`
}
type SiteMap struct {
	XMLName   xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	Urls      []Url
}

func bfsLinkCollector(currentlinksMap map[string]bool, domain string, maxDepth int) map[string]bool {
	resMap := make(map[string]bool)
	for key := range currentlinksMap {
		resMap[key] = true
	}
    if maxDepth <= 1 {
		return resMap
	}
    currentDepth := 1
	for i:= len(currentlinksMap); i > 0; {
        currentDepth += 1
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
        if currentDepth == maxDepth {
			return resMap
		}
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

func map2xml(m map[string]bool, domain string) (string, error) {
	urlBlock := &Url{}
    var all_urls []Url
	for key := range m {
        urlBlock.Loc = domain + key
        all_urls = append(all_urls, *urlBlock)
	}
	sitemap := &SiteMap{Urls: all_urls}
	xdata, err := xml.MarshalIndent(sitemap, " ", "    ")
	if err != nil {
		return "", err
	}
    return xml.Header + string(xdata), nil
}

func main() {
    domain := flag.String("domain", "http://127.0.0.1:8080", "The domain to building map")
    depth := flag.Int("depth", 3, "the depth of searching links from target domain")
	flag.Parse()
    fmt.Printf("Bulding the site map of %s with depth of %v.\n\n", *domain, *depth)
    firstMap, err := getLink(*domain, "/")
    if err != nil {
		log.Fatal(err)
	}
	resMap := bfsLinkCollector(firstMap, *domain, *depth)
    sxml, err := map2xml(resMap, *domain)
    if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sxml)
}