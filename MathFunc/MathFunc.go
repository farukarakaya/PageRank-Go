package MathFunc


func ArrMult(a []float32, mult int) []float32 {
     for k := 0; k < len(a); k++ {
        a[k] *= float32(mult)
     }
     return a
   }
