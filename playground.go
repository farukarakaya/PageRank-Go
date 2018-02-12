package main

import (
	"./PageRank"
	"./Wrapper"
	"fmt"
)

func main() {
	url := "http://pileyzmakineleri.blogspot.com.tr"
	destinations, urls := Wrapper.Get2DArray(url)
	fmt.Println(len(destinations))
	fmt.Println(len(urls))
	fmt.Println(destinations)
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
