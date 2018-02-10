package main
import (
"fmt"
"./MathFunc"
"reflect"
)

const beta float32 = 0.8
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
	rnew = make([]float32 ,len(data)) 
	MathFunc.ArrInit((1.0-beta)/float32(len(data)),rnew)
	//fmt.Println((1.0-beta)/float32(len(data)))
}
func calcPageRank() {
	var rold []float32
	isEqual :=reflect.DeepEqual(rnew, rold)
	//fmt.Println(data)
	rold = make([]float32 ,len(data))
	MathFunc.ArrInit(1/float32(len(data)),rold)
	for !isEqual {	
		for i := 0; i < len(data); i++ {
			for _,k := range data[i].destination{
				rnew[k] += beta * rold[i] / float32(len(data[i].destination))
				//fmt.Println(beta * rold[k] / float32(len(data[i].destination)))
			}
		}
		isEqual = reflect.DeepEqual(rnew, rold)
		fmt.Printf("%.32f     %.32f\n",rold, rnew)
		rold = rnew
		rnew = make([]float32 ,len(data))
		MathFunc.ArrInit((1.0-beta)/float32(len(data)),rnew)
	}	
}


func main() {
  	addrow([]int{1,2})  
	addrow([]int{0,2})
  	addrow([]int{1})
  	//addrow([]int{0,2})
  	initData()
  	calcPageRank()
	fmt.Println("Hello, Bitches")
	fmt.Println(data)
	fmt.Println(rnew)
}
