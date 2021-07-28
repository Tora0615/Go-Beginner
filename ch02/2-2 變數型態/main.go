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

	/*字串相關*/

	// 預設用雙引號宣告的一長串字串，是無法為了排版而直接用Enter進行換行宣告
	stringA := "01 一串文字A\n"
	fmt.Printf("%s",stringA)

	// 若按了 Enter 宣告，其仍是同一句，印出來不會換行
	stringB := "02 一串按了Enter" +
		"但同一行的文字B\n"
	fmt.Printf("%s",stringB)

	// 若要印出換行可以用換行轉義符號"\n"
	stringC := "03 一串加了換行\n轉義符號的文字C\n"
	fmt.Printf("%s",stringC)

	// 同樣的可以利用反引號"`"直接宣告多行字串
	// 多行字串中，未靠左到底的所有 Tab 與空格都會被計入
	stringD := `04 多行字串範例
	未靠左，有 Tab 的示範
靠左的示範
`
	fmt.Printf("%s",stringD)

	// 在多行字串中，放入程式碼是不會被編譯器識別的
	stringWithCode := `05 stringWithCode
a := 10
print (a)
`
	fmt.Printf("%s",stringWithCode)
	fmt.Printf("-------------------------------\n")

	/*字元*/
	var byteA byte = 'a'
	fmt.Printf("byte 是 uint8 的別名類型，表示 UTF-8 字串單個字元的值\n")
	fmt.Printf("byte 預設型態印出 : %v\n", byteA)
	fmt.Printf("byte 指定型態印出 : %c\n", byteA)

	var runeA rune = '龜'
	fmt.Printf("rune 是 int32 的別名類型，表示單個 unicode 字符\n")
	fmt.Printf("rune 預設型態印出 : %v\n", runeA)
	fmt.Printf("rune 指定型態印出 : %c\n", runeA)
}
