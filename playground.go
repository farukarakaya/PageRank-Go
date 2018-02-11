package main

import (
	"./PageRank"
	"./Wrapper"
	"fmt"
)

func main() {
<<<<<<< Updated upstream
	url := "https://www.google.com.tr/search?q=sucuk&oq=sucuk&aqs=chrome.0.69i59j0l5.795j0j8&sourceid=chrome&ie=UTF-8"
	destinations, urls := Wrapper.Get2DArray(url)
	//fmt.Println("DEST SIZE ", len(destinations), " URL SIZE", len(urls))
	//PrintLines(destinations)
=======
	url := "http://pileyzmakineleri.blogspot.com.tr/"
	destinations, urls := Wrapper.Get2DArray(url)
	fmt.Println(len(destinations))
	fmt.Println(len(urls))
	fmt.Println(destinations)
>>>>>>> Stashed changes
	fmt.Println(PageRank.GetPageRankJson(urls, destinations))
}

func PrintLines(a [][]int) {
	for _, elem := range a {
		for _, elem2 := range elem {
			fmt.Print(elem2, ",")
		}
		fmt.Println(" ")
	}
}
