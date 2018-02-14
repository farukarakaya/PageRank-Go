package Walker

import "PageRank-Go/Crawler/CrawlerEngine"

type TableRow struct {
	DESTINATION []int
	URL         string
}

type WalkerData struct {
	ROWS []TableRow
}



func PageWalker(url string) []TableRow {
	var ResultData WalkerData
	var firstRow TableRow
	firstRow.URL = url
	ResultData.ROWS = append(ResultData.ROWS, firstRow)
	firstRow.DESTINATION = DestinationSeparator(CrawlerEngine.SingleCrawler(url, url), firstRow, ResultData)

	for index, rowElem := range ResultData.ROWS {
		if index != 0 {
			rowElem.DESTINATION = DestinationSeparator(CrawlerEngine.SingleCrawler(url+rowElem.URL, url), rowElem, ResultData)
			ResultData.ROWS[index] = rowElem
		}
	}

	return ResultData.ROWS
}

func DestinationSeparator(urls []string, row TableRow, ResultData WalkerData) []int {
	for _, strElem := range urls {
		ind := Matcher(strElem, ResultData)
		if ind != -1 {
			row.DESTINATION = append(row.DESTINATION, ind)
		} else {
			row.DESTINATION = append(row.DESTINATION, len(ResultData.ROWS))
			var tempRow TableRow
			tempRow.URL = strElem
			ResultData.ROWS = append(ResultData.ROWS, tempRow)
		}
	}

	return row.DESTINATION
}

func Matcher(url string, ResultData WalkerData) int {
	for index, rowData := range ResultData.ROWS {
		if url == rowData.URL {
			return index
		}
	}

	return -1
}
