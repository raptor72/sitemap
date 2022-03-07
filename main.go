package main

import (
	"fmt"
    "flag"
	"log"
	"strings"
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

func bfsLinkCollector(currentlinksMap map[string]struct{}, domain string, maxDepth int) map[string]struct{} {
	resMap := make(map[string]struct{})
	for key := range currentlinksMap {
		resMap[key] = struct{}{}
	}
    if maxDepth <= 1 {
		return resMap
	}
    currentDepth := 1
	for i:= len(currentlinksMap); i > 0; {
        currentDepth += 1
		middleMap := make(map[string]struct{})
		for link := range currentlinksMap {
			newMap, _ := getLink(domain, link)
			for j := range newMap {
				if _, ok := resMap[j]; ok {
                    continue
				}
    			resMap[j] = struct{}{}
	    		middleMap[j] = struct{}{}
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

func checkIsTheSameDoamin(domain, link string) (same bool, preffix string) {
	d, _ := url.Parse(domain)
	l, err := url.Parse(link)
    if err != nil {
		return false, ""
	}
    if l.Scheme == "" && l.Host == "" {
		return true, domain
	}
    if l.Scheme == d.Scheme && l.Host == d.Host {
		return true, ""
	}
    return false, ""
}

func getLink(domain string, path string) (map[string]struct{}, error) {
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
	v := make(map[string]struct{})
	requested_url := sub_u.String()
	if sub_u.Scheme == "" && sub_u.Host == "" {
        requested_url = u.ResolveReference(sub_u).String()
	}
	resp, err := http.Get(requested_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	links, err := glink.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	for _, value := range links {
		same, preffix := checkIsTheSameDoamin(domain, value.Href)
		if same {
			v[preffix + value.Href] = struct{}{}
		}	
	}
	return v, nil
}

func map2xml(m map[string]struct{}) (string, error) {
	urlBlock := &Url{}
    var all_urls []Url
	for key := range m {
        urlBlock.Loc = key
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
    cut_domain := strings.TrimSuffix(*domain, "/")
	fmt.Printf("Bulding the site map of %s with depth of %v.\n\n", cut_domain, *depth)
    firstMap, err := getLink(cut_domain, "/")
    if err != nil {
		log.Fatal(err)
	}
	resMap := bfsLinkCollector(firstMap, cut_domain, *depth)
    sxml, err := map2xml(resMap)
    if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sxml)
}