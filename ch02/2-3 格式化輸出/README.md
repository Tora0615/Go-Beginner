格式化輸出(印出)
===
格式化輸出通常都會使用 fmt package，包含以下幾種印出格式。

# 通用印出格式
## A. 印出格式類型
有四種 : 
* %v
    * %v print the value in a default format 
    * 印出 Go 語言自動判斷的默認格式
* %+v
    * when printing structs, the plus flag (%+v) adds field names 
    * 在印出結構的時候，會印出欄位值 (Field,對岸稱字段) 
* %#v
    * a Go-syntax representation of the value
    * 用 GO 語法印出值
* %T
    * 印出值的類型
    
## B. 範例 : 
### 1. 小範例 : 
```go=
testIntA := 32
fmt.Printf("%T,%v\n", testIntA,testIntA)
testString := "測試"
fmt.Printf("%T,%v\n", testString,testString)
var testCharA byte = 'a'
fmt.Printf("%T,%v\n", testCharA,testCharA)  // byte 為 uint8 的別名，因此預設不會印出該字元
var testCharB rune = '龜'
fmt.Printf("%T,%v\n", testCharB,testCharB)  // rune 為 int32 的別名，因此預設不會印出該字元
```

### 2. %v %+v %#v 的不同
```go=
//宣告結構
type student struct {  
    id int
    name string
}
//結構賦值
studentA := student{id: 01, name: "小明"}  

// %v print the value in a default format  (以預設格式印出)
fmt.Printf("%T,%v\n",studentA,studentA)  
// Output 只印出值 : main.student, {1 小明}

//when printing structs, the plus flag (%+v) adds field names (在印出結構的時候，會印出欄位值 (Field,對岸稱字段) )
fmt.Printf("%T,%+v\n",studentA,studentA) 
// Output 印出欄位 & 值 : main.student, {id:1 name:小明}

//a Go-syntax representation of the value (用 GO 語法印出值)
fmt.Printf("%T,%#v\n",studentA,studentA) 
// Output 印出型態 & 欄位 & 值 : main.student, main.student{id:1, name:"小明"}
```


# 布林型態印出格式
## A. 印出格式類型
只有一種 : 
* %t
    * 會印出 true or false
## B. 範例
```go=
boolA := true
fmt.Printf("%T, %t", boolA, boolA)
// output : bool, true
```

# 整數印出格式
## A. 印出格式類型
有 6 種 : 
* 常用 : 
    * %d : 用以印出十進位制整數
        * 進階用法 : 前面加數字
            * %8d : 表示整數長度有8個，不足8個 **補空格**，超出以實際為準
            * %08d : 表示整數長度有8個，不足8個 **補0**，超出以實際為準
* 不那麼常用 : 
    * %b : 用以印出二進位制整數
    * %o : 用以印出八進位制整數
    * %x : 用以印出十六進位制整數(小寫)
    * %X : 用以印出十六進位制整數(大寫)


## B. 例子
```go=
intA := 123
// 常用
fmt.Printf("%d\n",intA)   //印出整數
// Output : 123
fmt.Printf("%1d\n",intA)  //表示整數長度有1個，不足1個補空格，超出以實際為準
// Output : 123
fmt.Printf("%4d\n",intA)  //表示整數長度有4個，不足4個補空格，超出以實際為準
// Output : (空格)123
fmt.Printf("%04d\n",intA) //表示整數長度有4個，不足4個補0，超出以實際為準
// Output : 0123

// 不那麼常用
fmt.Printf("%b\n", intA)  //二進位
// Output : 1111011
fmt.Printf("%o\n", intA)  //八進位
// Output : 173
fmt.Printf("%x\n", intA)  //十六進位(小寫)
// Output : 7b
fmt.Printf("%X\n", intA)  //十六進位(大寫)
// Output : 7B
```

# 字元印出形式
## A. 印出格式類型
有3種
* 常用
    * %c : 用此值對應的 unicode 值 **印出字元**
* 幾乎不會用
    * %U : 用 unicode 格式表示
    * %q : 用以印出**單引號**圍繞的字符字面值(此處為字元)，由 Go 語法安全地轉義
## B. 例子
```go=
charA := 123
charB := '龜'
// 常用
fmt.Printf("%T, %c, %T, %c\n",charA,charA,charB,charB)  //用此值對應的 unicode 值印出字元
// Output : int, {, int32, 龜

// 不常用
fmt.Printf("%T, %q, %T, %q\n", charA,charA,charB,charB)  //用以印出單引號圍繞的字符字面值(字元)，由 Go 語法安全地轉義
// Output : int, '{', int32, '龜'
fmt.Printf("%T, %U, %T, %U\n", charA,charA,charB,charB)  //用 unicode 格式表示
// Output : int, U+007B, int32, U+9F9C
```

# 浮點數印出形式
## A. 印出格式類型
有 7 種 : 
* 常用
    * %f : 印出含小數的數字 (預設值等同於 %.6f)
        * 可在 %f 中間加 .數字，改變要呈現小數點幾位
    * %F : 等同於 %f
    * %e : 將小數用科學表示法表示 (預設值等同於 %.6e)
        * 可在 %e 中間加 .數字，改變要呈現小數點幾位
    * %E : 幾乎等同於 %e，但科學記號輸出裡的 e 會是大寫
* 不那麼常用
    * %g : 自動選擇 %f 或是 %e (e小寫)
    * %G : 自動選擇 %f 或是 %E (e大寫)
* 不常用
    * %b : 用二進位制表示
    
## B. 例子
```go=
floatA := 1234567.1234567
fmt.Printf("%.2f\n", floatA)  //限制到小數點下兩位
// output : 1234567.12
fmt.Printf("%e, %E\n", floatA,floatA)  //科學記號輸出裡的 e 小寫 / 大寫
// output : 1.234567e+06, 1.234567E+06

floatB := 43.21
fmt.Printf("%g, %g\n", floatA,floatB)  //自動選擇 %f 或是 %e (小寫)
// output : 1.2345671234567e+06, 43.21
fmt.Printf("%G, %G\n", floatA,floatB)  //自動選擇 %f 或是 %E (大寫)
// output : 1.2345671234567E+06, 43.21

fmt.Printf("%b\n", floatB)  //用二進位表示
// output : 6081266871833723p-47
```

# 複數印出形式
## A. 印出格式類型
無格式化輸出佔位符號，較為特殊，只能用 fmt.Print / fmt.Println。

## B. 例子
```go=
complexA := 1 + 2i  // 可以這樣設定初值
fmt.Println(complexA)  //印整個複數
// output : (1+2i)

complexB := complex(3,4)  // 也能這樣設定初值
fmt.Println(real(complexB))  //印出實數部分
// output : 3
fmt.Println(imag(complexB))  //印出虛數部分
// output : 4
```

# 字串與字節陣列 (byte array) 印出形式
## A. 印出格式類型
有4種 : 
* 常用
    * %s : 直接輸出字串或是字節陣列 (byte array) (若是 rune 陣列，%s 印不出來)
* 不那麼常用
    * %x : 每用到一個 byte，就改用兩個十六進位字元表示(小寫)
    * %X : 每用到一個 byte，就改用兩個十六進位字元表示(大寫)
* 不常用
    * %q : 用以印出**雙引號**圍繞的字符字面值(此處為字串)，由 Go 語法安全地轉義
## B. 例子
```go=
stringA := "一串文字"
fmt.Printf("%s\n", stringA)  //直接輸出字串或是字節陣列 (byte array) (若是 rune 陣列，%s 印不出來)
// output : 一串文字
fmt.Printf("%x\n", stringA)  //每用到一個 byte，就改用兩個十六進位字元表示(小寫)
// output : e4b880e4b8b2e69687e5ad97
fmt.Printf("%X\n", stringA)  //每用到一個 byte，就改用兩個十六進位字元表示(大寫)
// output : E4B880E4B8B2E69687E5AD97
fmt.Printf("%q\n", stringA)  //用以印出雙引號圍繞的字符字面值(此處為字串)，由 Go 語法安全地轉義
// output : "一串文字"

// 字元陣列
charArray := []byte{'a','b','c','d'}  // 若是 rune 陣列，%S 印不出來
fmt.Printf("%T, %s\n", charArray,charArray)
// output : []uint8, abcd
fmt.Printf("%T, %x\n", charArray,charArray)
// output : []uint8, 61626364
fmt.Printf("%T, %X\n", charArray,charArray)
// output : []uint8, 61626364
fmt.Printf("%T, %q\n", charArray,charArray)
// output : []uint8, "abcd"
```

# 補充 : 字元、字串、Unicode 與 utf-8 編碼
## 問題
為什麼 string "龜龜" 16 進制格式化輸出 (%X) 為 E9BE9CE9BE9C

而不是字元 '龜' 的 16 進制格式化輸出 9F9C 拿來重複兩次呢 ?

### A. 簡答
因為 Go 的字元 16 進制格式化輸出，是直接將 Unicode 的碼點用 16 進位顯示。(Unicode 碼點原本是以 int32 儲存起來的)

而字串的 16 進制格式化輸出，是顯示文字的 unicode 經 utf-8 編碼過的 16 進位結果

### B. 十六進位字串轉回個別字元之細節剖析
1. 已知 string "龜龜" %X 輸出為 : 
   * E9BE 9CE9 BE9C

2. 轉二進位 : 
    * 1110 1001 1011 1110 1001 1100 1110 1001 1011 1110 1001 1100

3. 以 Byte (8bit) 為一組 : 
   * 11101001 10111110 10011100 11101001 10111110 10011100
4. 由首個 byte 1110 有三個 1，可知這個字含自己有三個 byte 
   * 分成兩個字 : 
      * 11101001 10111110 10011100
      * 11101001 10111110 10011100

5. 去除第一個 byte 開頭符號 1110 與後續每個 byte 的 10 開頭
   * ~~1110~~1001 ~~10~~111110 ~~10~~011100 
   * ~~1110~~1001 ~~10~~111110 ~~10~~011100

6. 剩餘的合併分組
   * 1001 1111 1001 1100 
   * 1001 1111 1001 1100
7. 轉為十六進位，對應就是 Unicode 碼點 : 
   * 9 F 9 C 
   * 9 F 9 C
8. 分別拿去 Unicode 字元解碼，可得都是 '龜'

### C. 參考資料
[UTF-8 編碼格式之簡單講解](https://davidhu0903ex3.pixnet.net/blog/post/468848723-utf-8-%E7%B7%A8%E7%A2%BC%E6%A0%BC%E5%BC%8F%E4%B9%8B%E7%B0%A1%E5%96%AE%E8%AC%9B%E8%A7%A3)