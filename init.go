package main

import (
	"fmt"
	"./PageRank"
	"./Server"
)
func main() {
	adrs := []string{"www.1.com","www.2.com","www.3.com"}
	dests := [][]int{{0,1},{0,2},{1}}
	
	//fmt.Println(adrs)
	//fmt.Println(dests)
	str := PageRank.GetPageRankJson(adrs,dests)
	fmt.Println(str)
	Server.ServeJson(str)
}
