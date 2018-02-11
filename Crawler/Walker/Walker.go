package Walker

import "../CrawlerEngine"

type TableRow struct {
	DESTINATION []int
	URL         string
}

type WalkerData struct {
	ROWS []TableRow
}

var ResultData WalkerData

func PageWalker(url string) []TableRow {
	var firstRow TableRow
	firstRow.URL = url
	ResultData.ROWS = append(ResultData.ROWS, firstRow)
	firstRow.DESTINATION = DestinationSeparator(CrawlerEngine.SingleCrawler(url, url), firstRow)

	for index, rowElem := range ResultData.ROWS {
		if index != 0 {
			rowElem.DESTINATION = DestinationSeparator(CrawlerEngine.SingleCrawler(url+rowElem.URL, url), rowElem)
			ResultData.ROWS[index] = rowElem
		}
	}

	return ResultData.ROWS
}

func DestinationSeparator(urls []string, row TableRow) []int {
	for index, strElem := range urls {
		if Matcher(strElem) != -1 {
			row.DESTINATION = append(row.DESTINATION, index)
		} else {
			row.DESTINATION = append(row.DESTINATION, len(ResultData.ROWS))
			var tempRow TableRow
			tempRow.URL = strElem
			ResultData.ROWS = append(ResultData.ROWS, tempRow)
		}
	}

	return row.DESTINATION
}

func Matcher(url string) int {
	for index, rowData := range ResultData.ROWS {
		if url == rowData.URL {
			return index
		}
	}

	return -1
}
