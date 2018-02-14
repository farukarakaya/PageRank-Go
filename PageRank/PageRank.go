package PageRank

import (
	"PageRank-Go/PageRank/MathFunc"
	"encoding/json"
	//"github.com/json-iterator/go"
)

const beta float32 = 0.8

//var json = jsoniter.ConfigCompatibleWithStandardLibrary

//var jsonGraph []byte
type Node struct {
	Id     int
	Adress string
	Value  float32
}
type Edge struct {
	From int
	To   int
}
type Graph struct {
	Nodes []Node
	Edges []Edge
}
type Datarow struct {
	adress      string
	destination []int
}
func initData(data []Datarow, rnew []float32) {
	//rnew = make([]float32, len(data))
	MathFunc.ArrInit((1.0-beta)/float32(len(data)), rnew)
	//fmt.Println((1.0-beta)/float32(len(data)))
}

func calcPageRank(data []Datarow, rnew []float32) {
	isEqual := false
	//fmt.Println(data)
	rold := make([]float32, len(data))
	MathFunc.ArrInit(1/float32(len(data)), rold)
	for !isEqual {
		for i := 0; i < len(data); i++ {
			if len(data[i].destination) > 0 {
				for _, k := range data[i].destination {
					rnew[k] += beta * rold[i] / float32(len(data[i].destination))
				}
			} else {
				for t := 0; t < len(data); t++ {
					rnew[t] += rold[i] / float32(len(data))
				}
			}
		}
		isEqual = MathFunc.CompareWithError(rnew, rold, 0.0000001)
		if isEqual {
			break
		}
		rold = rnew
		rnew = nil
		rnew = make([]float32, len(data))
		MathFunc.ArrInit((1.0-beta)/float32(len(data)), rnew)
	}
}

func serveJsonData(data []Datarow, rnew []float32) string {
	var nodes []Node
	var edges []Edge
	for i := 0; i < len(data); i++ {
		nodes = append(nodes, Node{i, data[i].adress, rnew[i]})
		for _, k := range data[i].destination {
			edges = append(edges, Edge{i, k})
		}
	}
	graph := Graph{nodes, edges}
	jsong, _ := json.Marshal(graph)
	return string(jsong)
}

func GetPageRankJson(adrs []string, dests [][]int) string {
	var data []Datarow
	rnew := make([]float32, len(adrs))
	for i := 0; i < len(adrs); i++ {
		data = append(data, Datarow{adrs[i], dests[i]})
	}
	initData(data,rnew)
	calcPageRank(data,rnew)
	return serveJsonData(data,rnew)
}
