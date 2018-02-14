package main

import (
	"fmt"
	"PageRank-Go/Server"
	"PageRank-Go/PageRank"
)
func main() {
	adrs := []string{"www.1.com","www.2.com","www.3.com"}
	dests := [][]int{{0,1},{0,2},{1}}
	
	//fmt.Println(adrs)
	fmt.Println(PageRank.GetPageRankJson(adrs,dests))
	fmt.Println(	)
	fmt.Println(PageRank.GetPageRankJson(adrs,dests))
	fmt.Println(	)
	Server.ServeJson()
}
