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
        XMLName   xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
		Urls      []Url
	}
	v := &Url{}
    var all_urls []Url
	for key := range m {
        v.Loc = key
        all_urls = append(all_urls, *v)
	}
	sitemap := &SiteMap{Urls: all_urls}
	xdata, err := xml.MarshalIndent(sitemap, " ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// xdata = append(xdata, []uint8(xml.Header))
	// os.Stdout.Write([]byte(xml.Header))
 	// os.Stdout.Write(xdata)
    fmt.Println(xml.Header + string(xdata))
	// fmt.Printf("%T\n", xml.Header)
}