package main

import (
	"fmt"
	"./PageRank"
)
func main() {
	adrs := []string{"www.1.com","www.2.com","www.3.com"}
	dests := [][]int{{0,1},{0,2},{1}}
	
	fmt.Println(PageRank.GetPageRankJson(adrs,dests))
	
}
