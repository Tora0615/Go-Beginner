package main

func main() {

	/* A. 變數宣告部分 */
	//基礎宣告模式之一 : 有指定變數型態，非同句賦值
	var a int  //var 變數名稱 變數型態
	a = 10    //變數型態 = 數值
	println(a)

	//基礎宣告模式之二 : 有指定變數型態
	var b int = 20  //var 變數名稱 變數型態 = 數值
	println(b)

	//基礎宣告模式之三 : 編譯器自動判斷變數型態
	var c = 30
	println(c)

	// 基礎宣告模式之四 : 批量宣告
	var (
		d int  //預設為 0
		f float32 //預設為 0
	)
	println(d)
	println(f)

	// 簡短聲明模式之一 : 宣告並賦值單個、自動判斷型態
	g := 40
	println(g)

	// 簡短聲明模式之二 : 宣告並賦值多個、自動判斷型態
	h,i := 50,60
	println(h)
	println(i)

	/*------------------------*/
	/* B. 變數多重賦值部分 */
	var j int = 80
	var k int = 70

	//左值右值按照左到右順序賦值
	j,k = k,j
	println(j,k)

	/*------------------------*/
	/* C. 匿名變量部分 */

	//GetData()函數要宣告在main函數外

	// 捨棄第二個回傳值
	l,_ := GetData()
	// 捨棄第一個回傳值
	_,m := GetData()

	println(l)
	println(m)
}

func GetData() (int,int){
	return 90,100
}
