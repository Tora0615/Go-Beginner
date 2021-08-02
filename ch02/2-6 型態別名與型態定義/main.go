package main

import "fmt"


type NewInt int  // 型態定義。設定為新的類型，叫做 NewInt，新類型本身具有 int 的特性
type IntAlias = int  // 型態別名。將 IntAlias 定義為 int 的別名 (綽號)，此綽號只存在於此，編譯時會以原本的型態覆蓋掉

func main()  {

	var A NewInt
	fmt.Printf("A type: %T\n", A)
	// A type: main.NewInt

	var B IntAlias
	fmt.Printf("B type: %T\n", B)
	// B type: int
}
