package Wrapper

import "../Crawler/Walker"

var Destinations [][]int
var URL_Paths []string

func Get2DArray(url string) ([][]int, []string) {
	allRes := Walker.PageWalker(url)
	for elemIndex := 0; elemIndex < len(allRes); elemIndex++ {
		Destinations = append(Destinations, allRes[elemIndex].DESTINATION)
		URL_Paths = append(URL_Paths, allRes[elemIndex].URL)
	}

	return Destinations, URL_Paths
}
