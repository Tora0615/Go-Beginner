常量 (Const)
===

# 介紹與宣告
## A. 介紹
* 常量是**恆定不變**的值，如圓周率
* 常量是值得標示符，在程式執行時，**不會被修改**
* 常量只能是以下幾種 : 
    * 布林型態
    * 數字型態 (整數、小數、複數)
    * 字串
* 常量宣告而未使用，**不會在編譯時報錯**
## B. 宣告
### 1. 格式 : 
`const 識別符號(名稱) [類型(選填)] = 值`
* 類型可省略，編譯器可以自動推斷

比較 : 變數的宣告，最前面用的是 var (variable,變數)
### 2. 顯式類型定義範例
```go=
const A string = "Kevin"
```

### 3. 隱式類型定義範例
```go=
const A = "Kevin"
```
### 4. 多個常量同時聲明
```go=
const width, height = 100, "兩百"
```

### 5. 用常量組模擬枚舉
Go 目前沒有枚舉的功能，可以使用常量組模擬枚舉
```go=
const(
    unknown = 0
    female = 1
    male = 2
)
fmt.Printf("%d, %d, %d\n", unknown,female,male)
```

### 6. 常量組特性
常量組之中，若不指定類型與初始值，則會**與上一行非空常量相同**。

若都沒設值，會有錯誤 : Missing value in the const declaration
```go=
const (
    C = 10
    D
    E
)
fmt.Printf("%d, %d, %d\n", C,D,E)
	
// output : 10, 10, 10
```

## C. 特殊常量 IOTA
### 1. 介紹 : 
* iota 是一種特殊的常量 : 
  * 他是 **可被編譯器修改的常量**
  * 他只能用在 **常量的賦值** 中
  * 初始值為零
  * 每遇到一個 **常量**，其值自動加一
    * 可以理解成 : **常量的計數器**
  * 每遇到一個 **常量關鍵字(const)**，其值歸零
### 2. 簡易例子 : 
#### 例子 1 : 
iota 每遇到一個常量，值加一
```go=
const (
    F = iota
    G = iota
    H = iota
)
fmt.Printf("%d, %d, %d\n", F,G,H)
// output : 0, 1, 2
```
#### 例子 2 :
1. 不指定類型與初始值，與則會與上一行非空常量相同 : 
    * 因此 J 與 K 都為 iota
2. 每遇到一個常量，iota 值加一
    * 遇到 J，值從 0 加為 1
3. iota 賦值給常數
    * 值為 1 的 iota 的值設定給 J

```go=
const (
    I = iota
    J
    K
)
fmt.Printf("%d, %d, %d\n", I,J,K)
// output : 0, 1, 2
```
### 3. 複雜例子 :
#### 例子 1 : 
```go=
const (
    a = iota   //iota = 0
    b          //值與上面相同為 iota , iota += 1 (1)
    c          //值與上面相同為 iota , iota += 1 (2)
    d = "ha"   //獨立值，iota += 1 (3)
    e          //值與上面相同為 "ha" , iota += 1 (4)
    f = 100    //iota +=1 (5)
    g          //值與上面相同為 100  ,  iota +=1 (6)
    h = iota   //iota += 1 (7)
    i          //值與上面相同為 iota , iota += 1 (8)
)
fmt.Println(a,b,c,d,e,f,g,h,i)
// output : 0 1 2 ha ha 100 100 7 8
```
#### 例子 2 :
```go=
const (
    j=1<<iota   //1乘2的0次方 = 1
    k=3<<iota   //3乘2的1次方 = 6
    l           //3乘2的2次方 = 12
    m           //3乘2的3次方 = 24
)
fmt.Println("j=",j)
fmt.Println("k=",k)
fmt.Println("l=",l)
fmt.Println("m=",m)

// output :
// j= 1
// k= 6
// l= 12
// m= 24 
```

#### 例子 3 :
```go=
const (
    u         = iota * 42  // u == 0     (untyped integer constant)
    v float64 = iota * 42  // v == 42.0  (float64 constant)
    w         = iota * 42  // w == 84    (untyped integer constant)
)
fmt.Printf("%d, %f, %d", u,v,w)
// output : 0, 42.000000, 84
```