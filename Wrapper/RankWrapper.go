package Wrapper

import "PageRank-Go/Crawler/Walker"



func Get2DArray(url string) ([][]int, []string) {
	var Destinations [][]int
	var URL_Paths []string	
	allRes := Walker.PageWalker(url)
	for elemIndex := 0; elemIndex < len(allRes); elemIndex++ {
		Destinations = append(Destinations, allRes[elemIndex].DESTINATION)
		URL_Paths = append(URL_Paths, allRes[elemIndex].URL)
	}

	return Destinations, URL_Paths
}
