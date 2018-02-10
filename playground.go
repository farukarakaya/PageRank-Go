package main

import "fmt"
import "net/http"
import "strings"
import "golang.org/x/net/html"

func singleCrawler(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Can't reach " + url + ". Check your code")
		return
	} else {
		var linkCount int
		linkCount = 0
		tmpBody := resp.Body
		defer tmpBody.Close()
		z := html.NewTokenizer(tmpBody)
		for {
			tt := z.Next()
			switch {
			case tt == html.ErrorToken:
				fmt.Println("Total link count:", linkCount)
				return
			case tt == html.StartTagToken:
				t := z.Token()
				if t.Data == "a" {
					for _, tmpAttr := range t.Attr {
						if tmpAttr.Key == "href" {
							if strings.Index(tmpAttr.Val, url) != -1 {
								endpoint := strings.Replace(tmpAttr.Val, url, "", -1)
								if len(endpoint) > 1 {
									fmt.Println(endpoint)
									linkCount = linkCount + 1
								}
							} else if strings.HasPrefix(tmpAttr.Val, "/") {
								fmt.Println(tmpAttr.Val)
								linkCount = linkCount + 1
							}
							break
						}
					}
				}
			}
		}
	}
}

func main() {
	url := "http://pileyzmakineleri.blogspot.com.tr/"
	singleCrawler(url)
}
