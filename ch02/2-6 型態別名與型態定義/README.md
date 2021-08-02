# 型態別名與型態定義

## 型態別名 (暱稱)
* 概述 : 
    * 從既有的資料型態，定義一個型態別名 (暱稱)
    * 此綽號只存在於程式碼，編譯時會以原本的型態覆蓋掉
* 宣告格式 : 
    * `type 型態別名 = 資料型態`
## 型態定義
* 概述 :
    * 定義一種新的資料型態
    * 本身會依據後方所寫的資料型態，而擁有其特性
    * struct 應該與此有關
* 宣告格式 :
    * `type 新的型態名稱 資料型態`

## 範例 :
```go=
package main
import "fmt"

// 型態定義
// 設定為新的類型，叫做 NewInt，新類型本身具有 int 的特性
type NewInt int  

// 型態別名
// 將 IntAlias 定義為 int 的別名 (綽號)，此綽號只存在於此，編譯時會以原本的型態覆蓋掉
type IntAlias = int  

func main()  {

	var A NewInt
	fmt.Printf("A type: %T\n", A)
	// A type: main.NewInt

	var B IntAlias
	fmt.Printf("B type: %T\n", B)
	// B type: int
}
```