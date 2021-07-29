package main

import "fmt"

func main()  {







	/* 在十六顯示下，字串為何不是字元重複兩次呢? */
	wordTurtle := '龜'
	fmt.Printf("%T,%X\n",wordTurtle ,wordTurtle)
	stringTurtle := "龜龜"
	fmt.Printf("%T,%X\n", stringTurtle, stringTurtle)
	//因為 Go 的字元 16 進制格式化輸出，是直接將 Unicode 的碼點用 16 進位顯示。
	//(Unicode 碼點原本是以 int32 儲存起來的。)
	//而字串的 16 進制格式化輸出，是顯示文字的 unicode 經 utf-8 編碼過的 16 進位結果

}
