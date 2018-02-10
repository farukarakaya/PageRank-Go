package main
import (
"fmt"
"./MathFunc"
)

const beta float32 = 0.85
var rnew []float32
type datarow struct {
	destination []int
}

var data []datarow
func addrow(outgoingNodes []int) {
	s1 := datarow{destination : outgoingNodes}
	data = append(data,s1)
}
func initData() {
	rnew = make([]int ,len(data)) 
	MathFunc.ArrInit((1-beta)/len(data),rnew)
}
func calcPageRank() {
	rold
	for i := 0; i < len(data); i++ {
		
	}
}


func main() {
  	var testArr []int
  	testArr = make([]int,2)
  	testArr[0] = 1
  	testArr[1] = 100

  	addrow(testArr)
	fmt.Println("Hello, Bitches")
	fmt.Println(data)
}
