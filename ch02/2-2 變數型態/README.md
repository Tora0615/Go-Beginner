變數型態
===

# 概述 
變數型態分為 : 
* 基礎變數型態 (原生資料型別，Primitive DataType)
    * 整數型 (int)
    * 浮點數型 (float)
    * 複數型 (complex)
    * 布林變數型 (bool)
    * 字符型 (string)
    * 字串型 (byte, rune)
* 複合變數型態 (衍生資料型態\別，Derived DataType)
    * 陣列 (array)
    * 切片 (slice)
    * 映射 (map)
    * 函數 (function)
    * 結構體 (struct)
    * 通道 (channel)
    * 介面 (interface)
    * 指標 (pointer)
  
在此單元，只介紹基礎變數型態

# 整數型 (int)

## A. 分類
整數分兩大類 
* 有符號的整數 (signed integer)
  * 指的是有正負號之分的整數
  * int, int8, int16, int32, int64
* 無符號的整數 (unsigned integer)
  * 指的是只有正的整數
  * uint, uint8, uint16, uint32, uint64
  
特殊的 int : uintptr
* 用來存 pointer 

## B. 取值範圍
int / uint 後的 8,16,32,64 指的是容許多少 **bit** 
  * 例如 8 指的就是可容許 2 的 8 次方個數字
    * 2 的 8 次方為 256 
      * 在 int 中的範圍就是 -128 ~ 127
      * 在 uint 中的範圍就是 0 ~ 255
  
若 int 後無指定要多少 bit (如單純的 int / uint 宣告)，取值範圍則是取決於系統
* 若系統為 32 位元，則是 int / uint 32
* 若系統為 64 位元，則是 int / uint 64

## C. 取值範圍總表
* uint
  * uint8       : the set of all unsigned  8-bit integers (0 to 255)
  * uint16      : the set of all unsigned 16-bit integers (0 to 65535)
  * uint32      : the set of all unsigned 32-bit integers (0 to 4294967295)
  * uint64      : the set of all unsigned 64-bit integers (0 to 18446744073709551615)
* int
  * int8        : the set of all signed  8-bit integers (-128 to 127)
  * int16       : the set of all signed 16-bit integers (-32768 to 32767)
  * int32       : the set of all signed 32-bit integers (-2147483648 to 2147483647)
  * int64       : the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
* uintptr  
  * an unsigned integer large enough to store the uninterpreted bits of a pointer value

  
# 浮點數型 (float)
## A. 分類
有兩種類型 : float32、float64 
## B. 取值範圍
* float32 
  * 最大值 : 3.402823 x 10^38
  * 正數最小值 : 1.401298 x 10^-45
* float64 
  * 最大值 : 1.797693 x 10^308
  * 正數最小值 : 4.940656 x 10^-324 
  
# 複數 (complex)
用以表示數學中的複數
## A. 分類
有兩種類型 : complex64、complex128。

初始化讓 Go 自動判斷型態的型態會是 complex128

## B. 取值範圍
complex64 : 實數與虛數範圍個別是一個 float32 的範圍
complex128 : 實數與虛數範圍個別是一個 float64 的範圍

## C. 宣告與印出
```go= 
var complexValue1 complex64 = 1.2 + 12i
fmt.Println("complexValue1 =", complexValue1)
// Output : complexValue1 = (1.2+12i)

complexValue2 := complex(3.2, 12)  //自動判斷預設 complex128
fmt.Println("complexValue2 實數 =", real(complexValue2))
fmt.Println("complexValue2 虛數 =", imag(complexValue2))
// Ountput :
// complexValue2 實數 = 3.2
// complexValue2 虛數 = 12
```

# 布林變數型 (String)
## A. 類型
就只有一種 : bool 

## B. 特性
1. 只可以是常量 true 或 false
   * 不像 C 語言，非 0 為真、 0 為假
2. 無法參與數值運算、也無法進行型態轉換

# 字串類型 (String)
## A. 類型
只有一種 : string

## B. 單行與多行宣告
* 預設用雙引號宣告的一長串字串，是無法為了排版而直接用Enter進行換行宣告
* 若按了 Enter 宣告，其仍是同一句，印出來不會換行
  ```go=
  stringB := "02 一串按了Enter" +
          "但同一行的文字B\n"
  fmt.Printf("%s",stringB)
  
  // Output : 02 一串按了Enter但同一行的文字B
  ```
* 若要印出換行可以用換行轉義符號"\n"
  ```go=
  stringC := "03 一串加了換行\n轉義符號的文字C\n"
  fmt.Printf("%s",stringC)
  
  // Output : 
  03 一串加了換行
  轉義符號的文字C
  ```

* 同樣的可以利用反引號"\`"直接宣告多行字串。多行字串中，未靠左到底的所有 Tab 與空格都會被計入
  ```go=
  stringD := `04 多行字串範例
    未靠左，有 Tab 的示範
  靠左的示範
  `
  fmt.Printf("%s",stringD)

  // Output : 
  04 多行字串範例
      未靠左，有 Tab 的示範
  靠左的示範
  ```

* 在多行字串中，放入程式碼是不會被編譯器識別的
  ```
  stringWithCode := `05 stringWithCode
  a := 10
  print (a)
  `
  fmt.Printf("%s",stringWithCode)

  // Output : 
  05 stringWithCode
  a := 10
  print (a)
  ```