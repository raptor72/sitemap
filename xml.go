package main

import (
	// "os"
	"fmt"
	"encoding/xml"
)

func main() {
    m := make(map[string]bool)
    m = map[string]bool{"/debate":true, "/denver":true, "/home":true, "/mark-bates":true, "/sean-kelly":true}


	type Url struct {
		XMLName   xml.Name `xml:"url"`
		Loc       string   `xml:"loc"`
	}

    type SiteMap struct {
		XMLName   xml.Name `xml:"urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9""`
		Url       []Url
	}

	fmt.Println(m)

	v := &Url{}

	var all_urls []*Url
	for key := range m {
        v.Loc = key
		// xdata, err := xml.MarshalIndent(v, " ", "    ")
		// if err != nil {
			// fmt.Printf("error: %v\n", err)
		// }
		// os.Stdout.Write(xdata)
        // fmt.Printf("%T\n", xdata)
        all_urls = append(all_urls, v)
	}
    fmt.Println(all_urls)
	sitemap := &SiteMap{Url: *all_urls}
	// sitemap.Url = []*all_urls



}