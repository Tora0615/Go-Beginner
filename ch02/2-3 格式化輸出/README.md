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



# 整數印出格式

# 字元、字串、Unicode 與 utf-8 編碼
## 問題
為什麼 string "龜龜" 16 進制格式化輸出 (%X) 為 E9BE9CE9BE9C

而不是字元 '龜' 的 16 進制格式化輸出 9F9C 拿來重複兩次呢 ?

## 簡答
因為 Go 的字元 16 進制格式化輸出，是直接將 Unicode 的碼點用 16 進位顯示。(Unicode 碼點原本是以 int32 儲存起來的)

而字串的 16 進制格式化輸出，是顯示文字的 unicode 經 utf-8 編碼過的 16 進位結果

## 十六進位字串轉回個別字元之細節剖析
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

## 參考資料
[UTF-8 編碼格式之簡單講解](https://davidhu0903ex3.pixnet.net/blog/post/468848723-utf-8-%E7%B7%A8%E7%A2%BC%E6%A0%BC%E5%BC%8F%E4%B9%8B%E7%B0%A1%E5%96%AE%E8%AC%9B%E8%A7%A3)