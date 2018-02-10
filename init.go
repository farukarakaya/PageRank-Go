package main

import (
	"./MathFunc"
	"fmt"
)

const beta float32 = 0.8

var rnew []float32

type datarow struct {
	destination []int
}

var data []datarow

func addrow(outgoingNodes []int) {
	s1 := datarow{destination: outgoingNodes}
	data = append(data, s1)
}

func initData() {
	rnew = make([]float32, len(data))
	MathFunc.ArrInit((1.0-beta)/float32(len(data)), rnew)
	//fmt.Println((1.0-beta)/float32(len(data)))
}

func calcPageRank() {
	var rold []float32
	isEqual := false
	//fmt.Println(data)
	rold = make([]float32, len(data))
	MathFunc.ArrInit(1/float32(len(data)), rold)
	for !isEqual {
		for i := 0; i < len(data); i++ {
			for _, k := range data[i].destination {
				rnew[k] += beta * rold[i] / float32(len(data[i].destination))
			}
		}
		isEqual = MathFunc.CompareWithError(rnew, rold, 0.0000001)
		if isEqual {
			break
		}
		rold = rnew
		rnew = make([]float32, len(data))
		MathFunc.ArrInit((1.0-beta)/float32(len(data)), rnew)
	}
}

func main() {
	addrow([]int{0, 1})
	addrow([]int{0, 2})
	addrow([]int{1})

	initData()
	calcPageRank()
	fmt.Println(rnew)
}
