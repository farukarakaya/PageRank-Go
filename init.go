package main
import "fmt"
import "./MathFunc"

func main() {
  var testArr []float32
  testArr = make([]float32,2)
  testArr[0] = 1
  testArr[1] = 100

  testArr = MathFunc.ArrMult(testArr, 2);
	fmt.Println(testArr)
}
