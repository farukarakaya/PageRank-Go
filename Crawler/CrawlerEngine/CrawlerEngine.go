package CrawlerEngine

import "net/http"
import "strings"
import "./vendor/golang.org/x/net/html"

func SingleCrawler(url string, staticURL string) []string {
	var allURL []string
	resp, err := http.Get(url)
	if err != nil {
		return nil
	} else {
		tmpBody := resp.Body
		defer tmpBody.Close()
		z := html.NewTokenizer(tmpBody)
		for {
			tt := z.Next()
			switch {
			case tt == html.ErrorToken:
				return allURL
			case tt == html.StartTagToken:
				t := z.Token()
				if t.Data == "a" {
					for _, tmpAttr := range t.Attr {
						if tmpAttr.Key == "href" {
							if strings.Index(tmpAttr.Val, staticURL) != -1 {
								endpoint := strings.Replace(tmpAttr.Val, staticURL, "", -1)
								if len(endpoint) > 1 {
									endpoint = PurifyString(endpoint)
									allURL = append(allURL, endpoint)
								}
							} else if strings.HasPrefix(tmpAttr.Val, "/") {
								if !strings.HasPrefix(tmpAttr.Val, "//") {
									tmpVal := tmpAttr.Val
									tmpVal = PurifyString(tmpVal)
									allURL = append(allURL, tmpVal)
								}
							}
							break
						}
					}
				}
			}
		}
	}
	return allURL
}

func PurifyString(str string) string {
	var tmpLoc int = strings.Index(str, "#")
	if tmpLoc != -1 {
		str = str[:tmpLoc]
	}
	tmpLoc = strings.Index(str, "?")
	if tmpLoc != -1 {
		str = str[:tmpLoc]
	}
	return str
}
