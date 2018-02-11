package main

import (
	"./PageRank"
	"./Wrapper"
	"fmt"
)

func main() {
	url := "http://ciceklab.cs.bilkent.edu.tr/ercumentcicek/"
	destinations, urls := Wrapper.Get2DArray(url)
	//fmt.Println("DEST SIZE ", len(destinations), " URL SIZE", len(urls))
	fmt.Println(PageRank.GetPageRankJson(urls, destinations))
}
