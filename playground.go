package main

import "./Crawler/Walker"
import "fmt"

func main() {
	url := "http://pileyzmakineleri.blogspot.com.tr/"
	printWalker(Walker.PageWalker(url), url)
}

func printWalker(allRes []Walker.TableRow, url string) {
	for elemIndex := 0; elemIndex < len(allRes); elemIndex++ {
		fmt.Println("Path: ", allRes[elemIndex].URL, " Going Destinations: ", allRes[elemIndex].DESTINATION)
	}

	fmt.Println("Total count: ", len(allRes))
}
