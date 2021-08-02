package main

import "fmt"

func main()  {

	/* float to int */
	// 會有精準度的損失
	floatA := 1234.56
	floatToIntA := int(floatA)
	fmt.Printf("float : %f to int : %d\n",floatA, floatToIntA)

	/* function return float and convert to int*/
	// 會有精準度的損失
	returnFloatToInt := int(ReturnFloat())
	fmt.Printf("function's return flow : %f to int : %d\n",ReturnFloat(), returnFloatToInt)

	/* int to float*/
	intA := 68
	intToFloatA := float64(intA)
	fmt.Printf("int %d to float %f\n",intA, intToFloatA)

	/* int to string : 會顯示字面值 */
	// 相當於 byte or rune 轉 string
	// 此處轉 string 就是將對應編號的字元查找出來
	intToStringA := string(intA)
	fmt.Printf("int %d to string %s\n",intA, intToStringA)

	intB := 40860
	intToStringB := string(intB)
	fmt.Printf("int %d to string %s\n",intB, intToStringB)


	/* 轉換會有錯的部分 */

	/* float to string */
	// 不能直接轉換 :
	// string is the set of all strings of 8-bit bytes,
	// conventionally but not necessarily representing UTF-8-encoded text.
	// A string may be empty, but not nil.
	// Values of string type are immutable.

	// floatToStringA := string(floatA)
	// fmt.Printf("float %f to string %s",floatA, floatToStringA)

	/* string to int*/
	// 不允許字串轉int，會有以下錯誤 :
	// cannot convert str (type string) to type int


}

func ReturnFloat() float64 {
	return 12.34
}
