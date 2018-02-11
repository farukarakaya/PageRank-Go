package MathFunc

import "math"

func ArrMult(a []float32, mult int) []float32 {
	for k := 0; k < len(a); k++ {
		a[k] *= float32(mult)
	}
	return a
}

func ArrInit(num float32, arr []float32) {
	for i := 0; i < len(arr); i++ {
		arr[i] = num
	}
}

func CompareWithError(arrOne []float32, arrTwo []float32, err float32) bool {
	for index, elemVal := range arrOne {
		if math.Abs(float64(elemVal-arrTwo[index])) > float64(err) {
			return false
		}
	}
	return true
}
