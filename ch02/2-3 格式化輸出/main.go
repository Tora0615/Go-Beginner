package main

import "fmt"

func main()  {

	/*通用印出格式*/

	// 幾個小範例
	testIntA := 32
	fmt.Printf("%T,%v\n", testIntA,testIntA)
	testString := "測試"
	fmt.Printf("%T,%v\n", testString,testString)
	var testCharA byte = 'a'
	fmt.Printf("%T,%v\n", testCharA,testCharA)  // byte 為 uint8 的別名，因此預設不會印出該字元
	var testCharB rune = '龜'
	fmt.Printf("%T,%v\n", testCharB,testCharB)  // rune 為 int32 的別名，因此預設不會印出該字元


	// %v %+v %#v 的不同
	type student struct {  //宣告結構
		id int
		name string
	}
	studentA := student{id: 01, name: "小明"}  //結構賦值

	// %v print the value in a default format  (以預設格式印出)
	fmt.Printf("%T,%v\n",studentA,studentA)  // 只印出值 : main.student, {1 小明}

	//when printing structs, the plus flag (%+v) adds field names (在印出結構的時候，會印出欄位值 (Field,對岸稱字段) )
	fmt.Printf("%T,%+v\n",studentA,studentA) // 印出欄位 & 值 : main.student, {id:1 name:小明}

	//a Go-syntax representation of the value (用 GO 語法印出值)
	fmt.Printf("%T,%#v\n",studentA,studentA) //印出型態 & 欄位 & 值 : main.student, main.student{id:1, name:"小明"}
	fmt.Printf("---------------------------------\n")

	/*---------------------------------------------------------------*/

	/*布林型態格式格式化印出*/
	boolA := true
	fmt.Printf("%T, %t\n", boolA, boolA)
	fmt.Printf("---------------------------------\n")

	/*---------------------------------------------------------------*/
	/*整數型態格式化印出*/

	intA := 123
	// 常用
	fmt.Printf("%d\n",intA)   //印出整數
	fmt.Printf("%1d\n",intA)  //表示整數長度有1個，不足1個補空格，超出以實際為準
	fmt.Printf("%4d\n",intA)  //表示整數長度有4個，不足4個補空格，超出以實際為準
	fmt.Printf("%04d\n",intA) //表示整數長度有4個，不足4個補0，超出以實際為準

	// 不那麼常用
	fmt.Printf("%b\n", intA)  //二進位
	fmt.Printf("%o\n", intA)  //八進位
	fmt.Printf("%x\n", intA)  //十六進位(小寫)
	fmt.Printf("%X\n", intA)  //十六進位(大寫)
	fmt.Printf("---------------------------------\n")

	/*---------------------------------------------------------------*/
	/*字元型態格式化印出*/
	charA := 123
	charB := '龜'
	// 常用
	fmt.Printf("%T, %c, %T, %c\n",charA,charA,charB,charB)  //用此值對應的 unicode 值印出字元

	// 不常用
	fmt.Printf("%T, %q, %T, %q\n", charA,charA,charB,charB)  //用以印出單引號圍繞的字符字面值，由 Go 語法安全地轉義
	fmt.Printf("%T, %U, %T, %U\n", charA,charA,charB,charB)  //用 unicode 格式表示
	fmt.Printf("---------------------------------\n")

	/*---------------------------------------------------------------*/

	/*浮點數型態*/
	floatA := 1234567.1234567
	fmt.Printf("%.2f\n", floatA)  //限制到小數點下兩位
	fmt.Printf("%e,%E\n", floatA,floatA)  //科學記號輸出裡的 e 小寫 / 大寫

	floatB := 43.21
	fmt.Printf("%g,%g\n", floatA,floatB)  //自動選擇 %f 或是 %e (小寫)
	fmt.Printf("%G,%G\n", floatA,floatB)  //自動選擇 %f 或是 %E (大寫)

	fmt.Printf("%b\n", floatB)  //用二進位表示
	fmt.Printf("---------------------------------\n")

	/*---------------------------------------------------------------*/
	/*複數*/
	complexA := 1 + 2i  // 可以這樣設定初值
	fmt.Println(complexA)  //印整個複數

	complexB := complex(3,4)  // 也能這樣設定初值
	fmt.Println(real(complexB))  //印出實數部分
	fmt.Println(imag(complexB))  //印出虛數部分
	fmt.Printf("---------------------------------\n")

	/*---------------------------------------------------------------*/
	// 字串
	stringA := "一串文字"
	fmt.Printf("%s\n", stringA)  //直接輸出字串或是字節陣列 (byte array)
	fmt.Printf("%x\n", stringA)  //每用到一個 byte，就改用兩個十六進位字元表示(小寫)
	fmt.Printf("%X\n", stringA)  //每用到一個 byte，就改用兩個十六進位字元表示(大寫)
	fmt.Printf("%q\n", stringA)  //用以印出雙引號圍繞的字符字面值(此處為字串)，由 Go 語法安全地轉義

	// 字元陣列
	charArray := []byte{'a','b','c','d'}  // 若是 rune 陣列，%S 印不出來
	fmt.Printf("%T, %s\n", charArray,charArray)
	fmt.Printf("%T, %x\n", charArray,charArray)
	fmt.Printf("%T, %X\n", charArray,charArray)
	fmt.Printf("%T, %q\n", charArray,charArray)
	fmt.Printf("---------------------------------\n")

	/*---------------------------------------------------------------*/

	/* 補充 : 在十六顯示下，字串為何不是字元重複兩次呢? */
	wordTurtle := '龜'
	fmt.Printf("%T,%X\n",wordTurtle ,wordTurtle)
	stringTurtle := "龜龜"
	fmt.Printf("%T,%X\n", stringTurtle, stringTurtle)
	//因為 Go 的字元 16 進制格式化輸出，是直接將 Unicode 的碼點用 16 進位顯示。
	//(Unicode 碼點原本是以 int32 儲存起來的。)
	//而字串的 16 進制格式化輸出，是顯示文字的 unicode 經 utf-8 編碼過的 16 進位結果

}
