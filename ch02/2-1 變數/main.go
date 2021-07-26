package main

import "fmt"

func main() {
	var a int = 100 //基礎宣告模式 (var 變數名稱 變數型態 = 數值)
	var b = "abc" //自動辦別資料型態
	c := 123  //var自動判別之簡寫
	fmt.Printf("%d\n%s\n%d\n",a,b,c) //格式化印出

	d,e := 'e','d'  //多重宣告
	d,e = e,d  //互換
	fmt.Printf("%c,%c\n",d,e)

	f,_ := 'f','g'  //_為匿名變量，把值吃掉用
	fmt.Printf("%c\n",f)

	temp :=
		`Line 01
Line 02`
	fmt.Println(temp)

	g :=  "abc"
	fmt.Printf("%T , %v , %+v , %#v\n", g,g,g,g)
}
