package main

import (
	"fmt"
	"math"
)

func main()  {

	/* Float 範圍印出 */
	var a float32 = math.MaxFloat32
	var b float32 = math.SmallestNonzeroFloat32
	fmt.Printf("float32 最大值 : %E \n", a)
	fmt.Printf("float32 正數最小值 : %E \n", b)

	var c float64 = math.MaxFloat64
	var d float64 = math.SmallestNonzeroFloat64
	fmt.Printf("float64 最大值 : %E \n", c)
	fmt.Printf("float64 正數最小值 : %E \n", d)

	fmt.Printf("-------------------------------\n")

	/*複數的宣告與印出*/
	var complexValue1 complex64 = 1.2 + 12i
	fmt.Println("complexValue1 =", complexValue1)

	complexValue2 := complex(3.2, 12)  //自動判斷預設 complex128
	fmt.Println("complexValue2 實數 =", real(complexValue2))
	fmt.Println("complexValue2 虛數 =", imag(complexValue2))
	fmt.Printf("-------------------------------\n")


}
