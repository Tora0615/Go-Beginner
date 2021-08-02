package main

import "fmt"

func main()  {
	const A string = "Kevin"
	const B = "Kevin"

	const width, height = 100, "兩百"
	fmt.Printf("%d,%s\n", width,height)

	//模擬枚舉
	const(
		unknown = 0
		female = 1
		male = 2
	)
	fmt.Printf("%d, %d, %d\n", unknown,female,male)

	const (
		C = 10
		D
		E
	)
	fmt.Printf("%d, %d, %d\n", C,D,E)

	//iota 簡單練習1
	const (
		F = iota
		G = iota
		H = iota
	)
	fmt.Printf("%d, %d, %d\n", F,G,H)

	//iota 簡單練習2
	const (
		I = iota
		J
		K
	)
	fmt.Printf("%d, %d, %d\n", I,J,K)

	//複雜iota例子1
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

	//複雜iota例子2
	const (
		j=1<<iota //1乘2的0次方 = 1
		k=3<<iota //3乘2的1次方 = 6
		l		  //3乘2的2次方 = 12
		m		  //3乘2的3次方 = 24
	)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
	fmt.Println("m=",m)

	//複雜iota例子3
	const (
		u         = iota * 42  // u == 0     (untyped integer constant)
		v float64 = iota * 42  // v == 42.0  (float64 constant)
		w         = iota * 42  // w == 84    (untyped integer constant)
	)
	fmt.Printf("%d, %f, %d", u,v,w)


}
