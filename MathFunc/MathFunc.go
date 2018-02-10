package MathFunc


func ArrMult(a []float32, mult int) []float32 {
     for k := 0; k < len(a); k++ {
        a[k] *= float32(mult)
     }
     return a
   }
func ArrInit( num float32, arr []float32) {
	for i := 0; i < len(arr); i++ {
		arr[i] = num
	}
}